#------------------------------------------------------------------------------
# All lines of the recipe will be given to a single invocation of the shell
# rather than each line being invoked separately.
#------------------------------------------------------------------------------

.ONESHELL:
SHELL = /bin/bash

#------------------------------------------------------------------------------
# The variables below are used to run the make targets inside docker.
# The version of make in OSX is 3.81 but the one inside the container is 4.3
#------------------------------------------------------------------------------

MAKE_DOCKER_TAG := latest
MAKE_DOCKER_IMG := makefile-builder
MAKE_DOCKER_CMD := ./bin/docker ${MAKE_DOCKER_IMG}:${MAKE_DOCKER_TAG}

#------------------------------------------------------------------------------
# This target opens a SHELL inside the MAKE_DOCKER_IMG
#------------------------------------------------------------------------------

.PHONY: bash
ifeq ($(SKIP_DOCKER),)
bash: ; @ ${MAKE_DOCKER_CMD} make bash
else
bash: ; @ bash
endif

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

.PHONY: dag-code
ifeq ($(SKIP_DOCKER),)
dag-code: ; @ ${MAKE_DOCKER_CMD} make dag-code
else
dag-code: NAME := terraform
dag-code: VERSION := v0.14.8
dag-code: TMPDIR := $(shell mktemp -d)
dag-code:
	@ # Suppress echoing
	git -c "advice.detachedHead=false" clone --depth 1 \
	--branch ${VERSION} https://github.com/hashicorp/${NAME}.git ${TMPDIR}
	rsync -a --delete --exclude='.*' ${TMPDIR}/dag/ internal/dag
	rsync -a --delete --exclude='.*' ${TMPDIR}/tfdiags/ internal/tfd
	sed -i 's_github.com/hashicorp/terraform/tfd_github.com/h0tbird/terrago/internal/tfd_g' internal/dag/*
	sed -i '/github.com\/hashicorp\/terraform\/internal\/logging/d' internal/dag/dag_test.go
	sed -i 's/tfdiags/tfd/g' internal/dag/* internal/tfd/*
	gofmt -w internal && rm -rf ${TMPDIR}
endif
