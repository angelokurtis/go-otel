SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: build
build: format ## Compiles the source code.
	go build ./...

.PHONY: generate
generate: wire ## Compile-time Dependency Injection using code generation.
	$(WIRE) ./...

.PHONY: format
format: goimports-reviser gofumpt wsl ## Cleans up the code for easier reading and collaboration.
	$(GOIMPORTS_REVISER) -use-cache -rm-unused -format ./...
	$(GOFUMPT) -w -extra .
	$(WSL) -fix ./...

.PHONY: lint
lint: golangci-lint ## Analyze and report style, formatting, and syntax issues in the source code.
	$(GOLANGCI_LINT) run  ./...

##@ Release

.PHONY: current-version
current-version: svu ## Show the current version of the project.
	@current_version=$$($(SVU) current); \
	echo $$current_version

.PHONY: next-version
next-version: svu ## Calculate the next version, following semantic versioning.
	@current_version=$$($(SVU) current); \
	next_version=$$($(SVU) next); \
	if [ "$$current_version" = "$$next_version" ]; then \
		echo "Error: Current version is equal to next version."; \
		exit 1; \
	fi; \
	echo $$next_version

.PHONY: auto-release
auto-release: build svu ## Automate release tasks based on git log.
	@current_version=$$($(SVU) current); \
	next_version=$$($(SVU) next); \
	if [ "$$current_version" = "$$next_version" ]; then \
		echo "Error: Current version is equal to next version."; \
		exit 1; \
	fi; \
	current_branch=$$(git rev-parse --abbrev-ref HEAD); \
	remote_branch=$$(git for-each-ref --format='%(upstream:short)' refs/heads/"$$current_branch"); \
	if [ "$$remote_branch" != "origin/main" ]; then \
		echo "Error: You are not in the main branch."; \
		exit 1; \
	fi; \
	next_version=$$($(SVU) next); \
	git tag -a $$next_version -m "release $$next_version"; \
	git push origin $$next_version; \
	echo "$$next_version has been released successfully"

.PHONY: major-release
major-release: build svu ## Force a major release with significant changes.
	@current_branch=$$(git rev-parse --abbrev-ref HEAD); \
	remote_branch=$$(git for-each-ref --format='%(upstream:short)' refs/heads/"$$current_branch"); \
	if [ "$$remote_branch" != "origin/main" ]; then \
		echo "Error: You are not in the main branch."; \
		exit 1; \
	fi; \
	next_version=$$($(SVU) major); \
	git tag -a $$next_version -m "release $$next_version"; \
	git push origin $$next_version; \
	echo "$$next_version has been released successfully"

.PHONY: minor-release
minor-release: build svu ## Force a minor release with new features.
	@current_branch=$$(git rev-parse --abbrev-ref HEAD); \
	remote_branch=$$(git for-each-ref --format='%(upstream:short)' refs/heads/"$$current_branch"); \
	if [ "$$remote_branch" != "origin/main" ]; then \
		echo "Error: You are not in the main branch."; \
		exit 1; \
	fi; \
	next_version=$$($(SVU) minor); \
	git tag -a $$next_version -m "release $$next_version"; \
	git push origin $$next_version; \
	echo "$$next_version has been released successfully"

.PHONY: patch-release
patch-release: build svu ## Force a patch release with bug fixes.
	@current_branch=$$(git rev-parse --abbrev-ref HEAD); \
	remote_branch=$$(git for-each-ref --format='%(upstream:short)' refs/heads/"$$current_branch"); \
	if [ "$$remote_branch" != "origin/main" ]; then \
		echo "Error: You are not in the main branch."; \
		exit 1; \
	fi; \
	next_version=$$($(SVU) patch); \
	git tag -a $$next_version -m "release $$next_version"; \
	git push origin $$next_version; \
	echo "$$next_version has been released successfully"

##@ Tool Binaries

SVU = $(shell pwd)/bin/svu
.PHONY: svu
svu: ## Checks for svu installation and downloads it if not found.
	$(call go-get-tool,$(SVU),github.com/caarlos0/svu@v1.11.0)

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

WIRE = $(shell pwd)/bin/wire
.PHONY: wire
wire: ## Checks for wire installation and downloads it if not found.
	$(call go-get-tool,$(WIRE),github.com/google/wire/cmd/wire@v0.5.1-0.20220620021424-0675cdc9191c)

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
.PHONY: golangci-lint
golangci-lint: ## Checks for golangci-lint installation and downloads it if not found.
	$(call go-get-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.1)

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
