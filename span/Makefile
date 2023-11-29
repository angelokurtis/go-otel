SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: build
build: format ## Compiles the source code.
	go build ./...

.PHONY: format
format: goimports-reviser gofumpt wsl ## Cleans up the code for easier reading and collaboration.
	$(GOIMPORTS_REVISER) -use-cache -rm-unused -format ./...
	$(GOFUMPT) -w -extra .
	$(WSL) -fix ./...

.PHONY: lint
lint: golangci-lint ## Analyze and report style, formatting, and syntax issues in the source code.
	$(GOLANGCI_LINT) run  ./...

##@ Tool Binaries

GOIMPORTS_REVISER = $(shell pwd)/bin/goimports-reviser
.PHONY: goimports-reviser
goimports-reviser: ## Checks for goimports-reviser installation and downloads it if not found.
	$(call go-get-tool,$(GOIMPORTS_REVISER),github.com/incu6us/goimports-reviser/v3@v3.5.6)

GOFUMPT = $(shell pwd)/bin/gofumpt
.PHONY: gofumpt
gofumpt: ## Checks for gofumpt installation and downloads it if not found.
	$(call go-get-tool,$(GOFUMPT),mvdan.cc/gofumpt@v0.5.0)

WSL = $(shell pwd)/bin/wsl
.PHONY: wsl
wsl: ## Checks for wsl installation and downloads it if not found.
	$(call go-get-tool,$(WSL),github.com/bombsimon/wsl/v4/cmd...@v4.1.0)

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
.PHONY: golangci-lint
golangci-lint: ## Checks for golangci-lint installation and downloads it if not found.
	$(call go-get-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
