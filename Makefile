dag-code: NAME := terraform
dag-code: VERSION := v0.14.4
dag-code: TMPDIR := $(shell mktemp -d)
dag-code:
	@git clone --depth 1 --branch ${VERSION} https://github.com/hashicorp/${NAME}.git ${TMPDIR}
	@rsync -a --delete --exclude='.*' ${TMPDIR}/dag/ internal/dag
	@rsync -a --delete --exclude='.*' ${TMPDIR}/tfdiags/ internal/tfd
	@rm -rf ${TMPDIR}
	@gsed -i 's/tfdiags/tfd/g' internal/dag/* internal/tfd/*
	@gsed -i 's_github.com/hashicorp/terraform/tfd_github.com/h0tbird/terrago/internal/tfd_g' internal/dag/*
	@gsed -i '/github.com\/hashicorp\/terraform\/internal\/logging/d' internal/dag/dag_test.go
	@gofmt -w internal
