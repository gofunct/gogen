PATH := ${PWD}/bin:${PATH}
export PATH

.DEFAULT_GOAL := all

REVISION ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date +'%Y-%m-%dT%H:%M:%SZ')

GO_BUILD_FLAGS := -v
GO_TEST_FLAGS := -v -timeout 2m
GO_COVER_FLAGS := -coverprofile coverage.txt -covermode atomic
SRC_FILES := $(shell go list -f '{{range .GoFiles}}{{printf "%s/%s\n" $$.Dir .}}{{end}}' ./...)

XC_ARCH := 386 amd64
XC_OS := darwin linux windows


#  App
#----------------------------------------------------------------
BIN_DIR := ./bin
OUT_DIR := ./dist
GENERATED_BINS :=
PACKAGES :=

define cmd-tmpl

$(eval NAME := $(notdir $(1)))
$(eval OUT := $(addprefix $(BIN_DIR)/,$(NAME)))
$(eval LDFLAGS := -ldflags "-X main.revision=$(REVISION) -X main.buildDate=$(BUILD_DATE)")

$(OUT): $(SRC_FILES)
	go build $(GO_BUILD_FLAGS) $(LDFLAGS) -o $(OUT) $(1)

.PHONY: $(NAME)
$(NAME): $(OUT)

.PHONY: $(NAME)-package
$(NAME)-package: $(NAME) $(BIN_DIR)/gox
	gox \
		$(LDFLAGS) \
		-os="$(XC_OS)" \
		-arch="$(XC_ARCH)" \
		-output="$(OUT_DIR)/$(NAME)_{{.OS}}_{{.Arch}}" \
		$(1)

$(eval GENERATED_BINS += $(OUT))
$(eval PACKAGES += $(NAME)-package)

endef

$(foreach src,$(wildcard ./cmd/*),$(eval $(call cmd-tmpl,$(src))))


#  Commands
#----------------------------------------------------------------
.PHONY: setup
setup: ## setup
ifdef CI
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep ensure -v -vendor-only
	gex --build --verbose

.PHONY: clean
clean: ## clean bin
	rm -rf $(BIN_DIR)/*

.PHONY: gen
gen:## go generate
	go generate ./...

.PHONY: lint
lint: ## lint
ifdef CI
	gex reviewdog -reporter=github-pr-review
else
	gex reviewdog -diff="git diff master"
endif

.PHONY: test
test: ## test all
	go test $(GO_TEST_FLAGS) ./...

.PHONY: cover
cover: ## test coverage
	go test $(GO_TEST_FLAGS) $(GO_COVER_FLAGS) ./...

.PHONY: test-e2e
test-e2e: ## end to end test
	@./_tests/e2e/run_test.sh

.PHONY: all
all: $(GENERATED_BINS) ## generate all bins

.PHONY: packages
packages: $(PACKAGES) ## packages

install: ## go install all programs
	go install ./...

fmt: ## go install all programs
	go fmt ./...

demogen: ## delete and regenerate ../demogen
	rm -rf ../demogen
	cd ..; gogen init demogen

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
