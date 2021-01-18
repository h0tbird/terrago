#------------------------------------------------------------------------------
# Terrago is using the Terraform DAG implementation. The initial idea was to
# consume the upstream 'dag' package as a library. But importing and using
# 'terraform' into this project alongside 'terraform-plugin-sdk' causes panics
# when the types that are duplicated between those two repos are attempted to
# be initialized for a second time. The solution below copies the 'dag'
# package from upstream and it also copies (and renames to 'tfd') the 'tfdiags'
# package which is the one containing the duplicated types.
#
# Related to:
#  https://github.com/hashicorp/terraform-plugin-sdk/issues/277
#  https://github.com/hashicorp/terraform/issues/23725
#------------------------------------------------------------------------------

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
