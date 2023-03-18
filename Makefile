# gloabal
.DEFAULT_GOAL = help

comma := ,
empty :=
space := $(empty) $(empty)

# go
GOLANG_VERSION = 1.19

# tools
TOOLS := $(shell cd tools; go list -f '{{ join .Imports " " }}' -tags=tools)
TOOLS_DIR := ${CURDIR}/tools
TOOLS_BIN := ${TOOLS_DIR}/bin

.PHONY: tools
tools: tools/bin/'' tools/wrench

tools/bin/%:
	@cd tools; \
		for t in ${TOOLS}; do \
			if [ -z '$*' ] || [ $$(basename $$t) = '$*' ]; then \
				echo "Install $$t ..." >&2; \
				GOBIN=${TOOLS_BIN} go install -v -mod=readonly "$${t}"; \
			fi \
		done

.PHONY: tools/wrench
tools/wrench:
	wget -q --output-document ${TOOLS_BIN}/wrench.tar.gz https://github.com/cloudspannerecosystem/wrench/releases/download/v1.5.0/wrench-1.5.0-darwin-amd64.tar.gz
	tar xf ${TOOLS_BIN}/wrench.tar.gz --directory=${TOOLS_BIN} wrench
	rm ${TOOLS_BIN}/wrench.tar.gz


.PHONY: gen/yo
gen/yo:
	yo $(SPANNER_PROJECT_ID) $(SPANNER_INSTANCE_ID) $(SPANNER_DATABASE_ID) -o infrastructure/spanner/models
