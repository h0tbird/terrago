package main

//-----------------------------------------------------------------------------
// This is work-in-progress
//-----------------------------------------------------------------------------

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
)

//	m.Resources["nodesPolicy"] = &tg.Resource{
//		ResourceLogicalID: "NodesPolicy",
//		ResourceType:      "aws_iam_policy",
//		ResourceConfig: map[string]interface{}{
//			"name":        "nodes.cluster-api-provider-aws.sigs.k8s.io",
//			"description": "For the Kubernetes Cloud Provider AWS nodes",
//			"policy":      nodesPolicy,
//		},
//	}

//  7: *ast.AssignStmt {
//  .  Lhs: []ast.Expr (len = 1) {
//  .  .  0: *ast.IndexExpr {
//  .  .  .  X: *ast.SelectorExpr {
//  .  .  .  .  X: *ast.Ident {
//  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:2
//  .  .  .  .  .  Name: "m"
//  .  .  .  .  .  Obj: *(obj @ 220)
//  .  .  .  .  }
//  .  .  .  .  Sel: *ast.Ident {
//  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:4
//  .  .  .  .  .  Name: "Resources"
//  .  .  .  .  }
//  .  .  .  }
//  .  .  .  Lbrack: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:13
//  .  .  .  Index: *ast.BasicLit {
//  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:14
//  .  .  .  .  Kind: STRING
//  .  .  .  .  Value: "\"nodesPolicy\""
//  .  .  .  }
//  .  .  .  Rbrack: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:27
//  .  .  }
//  .  }
//  .  TokPos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:29
//  .  Tok: =
//  .  Rhs: []ast.Expr (len = 1) {
//  .  .  0: *ast.UnaryExpr {
//  .  .  .  OpPos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:31
//  .  .  .  Op: &
//  .  .  .  X: *ast.CompositeLit {
//  .  .  .  .  Type: *ast.SelectorExpr {
//  .  .  .  .  .  X: *ast.Ident {
//  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:32
//  .  .  .  .  .  .  Name: "tg"
//  .  .  .  .  .  }
//  .  .  .  .  .  Sel: *ast.Ident {
//  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:35
//  .  .  .  .  .  .  Name: "Resource"
//  .  .  .  .  .  }
//  .  .  .  .  }
//  .  .  .  .  Lbrace: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:70:43
//  .  .  .  .  Elts: []ast.Expr (len = 3) {
//  .  .  .  .  .  0: *ast.KeyValueExpr {
//  .  .  .  .  .  .  Key: *ast.Ident {
//  .  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:71:3
//  .  .  .  .  .  .  .  Name: "ResourceLogicalID"
//  .  .  .  .  .  .  }
//  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:71:20
//  .  .  .  .  .  .  Value: *ast.BasicLit {
//  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:71:22
//  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  Value: "\"NodesPolicy\""
//  .  .  .  .  .  .  }
//  .  .  .  .  .  }
//  .  .  .  .  .  1: *ast.KeyValueExpr {
//  .  .  .  .  .  .  Key: *ast.Ident {
//  .  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:72:3
//  .  .  .  .  .  .  .  Name: "ResourceType"
//  .  .  .  .  .  .  }
//  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:72:15
//  .  .  .  .  .  .  Value: *ast.BasicLit {
//  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:72:22
//  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  Value: "\"aws_iam_policy\""
//  .  .  .  .  .  .  }
//  .  .  .  .  .  }
//  .  .  .  .  .  2: *ast.KeyValueExpr {
//  .  .  .  .  .  .  Key: *ast.Ident {
//  .  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:3
//  .  .  .  .  .  .  .  Name: "ResourceConfig"
//  .  .  .  .  .  .  }
//  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:17
//  .  .  .  .  .  .  Value: *ast.CompositeLit {
//  .  .  .  .  .  .  .  Type: *ast.MapType {
//  .  .  .  .  .  .  .  .  Map: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:19
//  .  .  .  .  .  .  .  .  Key: *ast.Ident {
//  .  .  .  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:23
//  .  .  .  .  .  .  .  .  .  Name: "string"
//  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  Value: *ast.InterfaceType {
//  .  .  .  .  .  .  .  .  .  Interface: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:30
//  .  .  .  .  .  .  .  .  .  Methods: *ast.FieldList {
//  .  .  .  .  .  .  .  .  .  .  Opening: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:39
//  .  .  .  .  .  .  .  .  .  .  Closing: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:40
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  .  Incomplete: false
//  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  Lbrace: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:73:41
//  .  .  .  .  .  .  .  Elts: []ast.Expr (len = 3) {
//  .  .  .  .  .  .  .  .  0: *ast.KeyValueExpr {
//  .  .  .  .  .  .  .  .  .  Key: *ast.BasicLit {
//  .  .  .  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:74:4
//  .  .  .  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  .  .  .  Value: "\"name\""
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:74:10
//  .  .  .  .  .  .  .  .  .  Value: *ast.BasicLit {
//  .  .  .  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:74:19
//  .  .  .  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  .  .  .  Value: "\"nodes.cluster-api-provider-aws.sigs.k8s.io\""
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  1: *ast.KeyValueExpr {
//  .  .  .  .  .  .  .  .  .  Key: *ast.BasicLit {
//  .  .  .  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:75:4
//  .  .  .  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  .  .  .  Value: "\"description\""
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:75:17
//  .  .  .  .  .  .  .  .  .  Value: *ast.BasicLit {
//  .  .  .  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:75:19
//  .  .  .  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  .  .  .  Value: "\"For the Kubernetes Cloud Provider AWS nodes\""
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  2: *ast.KeyValueExpr {
//  .  .  .  .  .  .  .  .  .  Key: *ast.BasicLit {
//  .  .  .  .  .  .  .  .  .  .  ValuePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:76:4
//  .  .  .  .  .  .  .  .  .  .  Kind: STRING
//  .  .  .  .  .  .  .  .  .  .  Value: "\"policy\""
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  .  Colon: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:76:12
//  .  .  .  .  .  .  .  .  .  Value: *ast.Ident {
//  .  .  .  .  .  .  .  .  .  .  NamePos: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:76:19
//  .  .  .  .  .  .  .  .  .  .  Name: "nodesPolicy"
//  .  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  }
//  .  .  .  .  .  .  .  Rbrace: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:77:3
//  .  .  .  .  .  .  .  Incomplete: false
//  .  .  .  .  .  .  }
//  .  .  .  .  .  }
//  .  .  .  .  }
//  .  .  .  .  Rbrace: /Users/mvillacorta/git/h0tbird/clusterawsadm/main.go:78:2
//  .  .  .  .  Incomplete: false
//  .  .  .  }
//  .  .  }
//  .  }
//  }

func main() {

	// Parse the code
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "/Users/mvillacorta/git/h0tbird/clusterawsadm/main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// ast.Print(fset, f)

	ast.Inspect(f, func(n ast.Node) bool {

		var buf bytes.Buffer

		// Find KeyValueExpr statements
		if kve, ok := n.(*ast.KeyValueExpr); ok {

			// Get the key
			err := printer.Fprint(&buf, fset, kve.Key)
			if err != nil {
				log.Fatal(err)
			}

			// If key is 'ResourceType'
			if buf.String() == "ResourceType" {
				buf.Reset()
				// Get the value
				err := printer.Fprint(&buf, fset, kve.Value)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(buf.String())
			}
		}

		return true
	})
}
