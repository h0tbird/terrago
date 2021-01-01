package tg

//-----------------------------------------------------------------------------
// Imports
//-----------------------------------------------------------------------------

import (

	// stdlib
	"context"
	"fmt"
	"sync"

	// community
	"github.com/sirupsen/logrus"

	// terraform
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// terrago
	"github.com/h0tbird/terrago/internal/dag"
	"github.com/h0tbird/terrago/internal/tfd"
)

//-----------------------------------------------------------------------------
// Types
//-----------------------------------------------------------------------------

// Manifest ...
type Manifest struct {
	Resources map[string]*Resource
	Dag       dag.AcyclicGraph
}

//-----------------------------------------------------------------------------
// Methods
//-----------------------------------------------------------------------------

// New ...
func New() *Manifest {
	return &Manifest{
		Resources: map[string]*Resource{},
		Dag:       dag.AcyclicGraph{},
	}
}

// Apply ...
func (h *Manifest) Apply(ctx context.Context, p *schema.Provider, s State) error {

	// Setup the DAG
	for resKey, resVal := range h.Resources {

		// All vertices
		h.Dag.Add(resVal)
		match := false

		// Dependent edges
		for _, fieldVal := range resVal.ResourceConfig {
			submatch := reg.FindStringSubmatch(fieldVal.(string))
			if submatch != nil {
				h.Dag.Connect(dag.BasicEdge(h.Resources[submatch[1]], h.Resources[resKey]))
				match = true
			}
		}

		// Non-dependent edges
		if !match {
			h.Dag.Connect(dag.BasicEdge(0, h.Resources[resKey]))
		}
	}

	// Walk the DAG
	w := &dag.Walker{Callback: walk(ctx, p, s, h.Resources)}
	w.Update(&h.Dag)

	// Wait for the completion of the walk
	diags := w.Wait()
	if diags != nil && diags.HasErrors() {
		for _, d := range diags {
			if d.Severity() == tfd.Error {
				return fmt.Errorf("%v", d.Description())
			}
		}
	}

	// Success
	return nil
}

//-----------------------------------------------------------------------------
// walk
//-----------------------------------------------------------------------------

func walk(ctx context.Context, p *schema.Provider, s State, r map[string]*Resource) dag.WalkFunc {
	var l sync.Mutex
	return func(v dag.Vertex) tfd.Diagnostics {
		l.Lock()
		defer l.Unlock()

		rh := v.(*Resource)
		if err := rh.Reconcile(ctx, p, s, r); err != nil {
			// TODO: Return diagnostics
			logrus.Fatal(err)
		}

		return nil
	}
}
