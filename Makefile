# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Setup name variables for the package/tool
NAME := cucli
PKG := github.com/fantasticrabbit/$(NAME)

# Set any default go build tags
BUILDTAGS :=

# Set the build dir, where built cross-compiled binaries will be output
BUILDDIR := ${PREFIX}/cross

# Populate version variables
# Add to compile time flags
VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)

ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

CTIMEVAR=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

# Set our default go compiler
GO := go

# List the GOOS and GOARCH to build
GOOSARCHES = $(shell cat .goosarch)

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME): $(wildcard *.go) $(wildcard */*.go) VERSION.txt
	@echo "+ $@"
	$(GO) build -tags "$(BUILDTAGS)" ${GO_LDFLAGS} -o $(NAME) .

.PHONY: static
static: ## Builds a static executable
	@echo "+ $@"
	CGO_ENABLED=0 $(GO) build \
				-tags "$(BUILDTAGS) static_build" \
				${GO_LDFLAGS_STATIC} -o $(NAME) .

all: clean build install
## all: clean build fmt lint test staticcheck vet install ## Runs a clean, build, fmt, lint, test, staticcheck, vet and install
## 
## .PHONY: fmt
## fmt: ## Verifies all files have been `gofmt`ed
## 	@echo "+ $@"
## 	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr
## 
## .PHONY: lint
## lint: ## Verifies `golint` passes
## 	@echo "+ $@"
## 	@golint ./... | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr
## 
## .PHONY: test
## test: ## Runs the go tests
## 	@echo "+ $@"
## 	@$(GO) test -v -tags "$(BUILDTAGS) cgo" $(shell $(GO) list ./... | grep -v vendor) |  sed ''/PASS/s//$(shell printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(shell printf "\033[31mFAIL\033[0m")/''
## 
## .PHONY: vet
## vet: ## Verifies `go vet` passes
## 	@echo "+ $@"
## 	@$(GO) vet $(shell $(GO) list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr
## 
## .PHONY: staticcheck
## staticcheck: ## Verifies `staticcheck` passes
## 	@echo "+ $@"
## 	@staticcheck $(shell $(GO) list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr
## 
## .PHONY: cover
## cover: ## Runs go test with coverage
## 	@echo "" > coverage.txt
## 	@for d in $(shell $(GO) list ./... | grep -v vendor); do \
## 		$(GO) test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
## 		if [ -f profile.out ]; then \
## 			cat profile.out >> coverage.txt; \
## 			rm profile.out; \
## 		fi; \
## 	done;

.PHONY: dep
dep: ## Pulls in package dependencies
	@echo "+ $@"
	$(GO) mod download
	$(GO) mod tidy
	
.PHONY: install
install: ## Installs the executable or package
	@echo "+ $@"
	$(GO) install -a -tags "$(BUILDTAGS)" ${GO_LDFLAGS} .

define buildpretty
mkdir -p $(BUILDDIR)/$(1)/$(2);
GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 $(GO) build \
	 -o $(BUILDDIR)/$(1)/$(2)/$(NAME) \
	 -a -tags "$(BUILDTAGS) static_build netgo" \
	 -installsuffix netgo ${GO_LDFLAGS_STATIC} .;
md5sum $(BUILDDIR)/$(1)/$(2)/$(NAME) > $(BUILDDIR)/$(1)/$(2)/$(NAME).md5;
sha256sum $(BUILDDIR)/$(1)/$(2)/$(NAME) > $(BUILDDIR)/$(1)/$(2)/$(NAME).sha256;
endef

.PHONY: cross
cross: *.go VERSION.txt ## Builds the cross-compiled binaries, creating a clean directory structure (eg. GOOS/GOARCH/binary)
	@echo "+ $@"
	$(foreach GOOSARCH,$(GOOSARCHES), $(call buildpretty,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))

define buildrelease
GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 $(GO) build \
	 -o $(BUILDDIR)/$(NAME)-$(1)-$(2) \
	 -a -tags "$(BUILDTAGS) static_build netgo" \
	 -installsuffix netgo ${GO_LDFLAGS_STATIC} .;
md5sum $(BUILDDIR)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(NAME)-$(1)-$(2).md5;
sha256sum $(BUILDDIR)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(NAME)-$(1)-$(2).sha256;
endef

.PHONY: release
release: *.go VERSION.txt ## Builds the cross-compiled binaries, naming them in such a way for release (eg. binary-GOOS-GOARCH)
	@echo "+ $@"
	$(foreach GOOSARCH,$(GOOSARCHES), $(call buildrelease,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))

## .PHONY: bump-version
## BUMP := patch
## bump-version: ## Bump the version in the version file. Set BUMP to [ patch | major | minor ]
## 	@$(GO) get -u github.com/jessfraz/junk/sembump # update sembump tool
## 	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))
## 	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
## 	echo $(NEW_VERSION) > VERSION.txt
## 	@echo "Updating links to download binaries in README.md"
## 	sed -i s/$(VERSION)/$(NEW_VERSION)/g README.md
## 	git add VERSION.txt README.md
## 	git commit -vam "Bump version to $(NEW_VERSION)"
## 	@echo "Run make tag to create and push the tag for new version $(NEW_VERSION)"

.PHONY: tag
tag: ## Create a new git tag to prepare to build a release
	git tag -a $(VERSION) -m "$(VERSION)"
	@echo "Run git push origin $(VERSION) to push your new tag to GitHub and trigger a build."

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(RM) $(NAME)
	$(RM) -r $(BUILDDIR)

# if this session isn't interactive, then we don't want to allocate a
# TTY, which would fail, but if it is interactive, we do want to attach
# so that the user can send e.g. ^C through.
INTERACTIVE := $(shell [ -t 0 ] && echo 1 || echo 0)
ifeq ($(INTERACTIVE), 1)
    DOCKER_FLAGS += -t
endif

GRPC_API_DIR=api/grpc
.PHONY: protoc
protoc: $(CURDIR)/$(GRPC_API_DIR)/api.pb.go ## Generate the protobuf files

$(CURDIR)/$(GRPC_API_DIR)/api.pb.go: image $(CURDIR)/$(GRPC_API_DIR)/api.proto
			docker run $(DOCKER_FLAGS) \
						$(DOCKER_IMAGE) \
						protoc -I ./$(GRPC_API_DIR) \
						./$(GRPC_API_DIR)/api.proto \
						--go_out=plugins=grpc:$(GRPC_API_DIR)

.PHONY: help
help: ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'