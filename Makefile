dag-code:
	cp -R ~/git/hashicorp/terraform/dag/. internal/dag
	cp -R ~/git/hashicorp/terraform/tfdiags/. internal/tfd
	gsed -i 's/tfdiags/tfd/g' internal/dag/* internal/tfd/*
	gsed -i 's_github.com/hashicorp/terraform/tfd_github.com/h0tbird/terrago/internal/tfd_g' internal/dag/*
	gsed -i '/github.com\/hashicorp\/terraform\/internal\/logging/d' internal/dag/dag_test.go
	gofmt -w internal
