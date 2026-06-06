LOCAL_BIN := $(CURDIR)/bin
EASYP_BIN := $(LOCAL_BIN)/easyp
GOIMPORTS_BIN := $(LOCAL_BIN)/goimports
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
GO_TEST=$(LOCAL_BIN)/gotest
GO_TEST_ARGS=-race -v -tags=integration_test ./...
PROTOC_DOWNLOAD_LINK="https://github.com/protocolbuffers/protobuf/releases"
PROTOC_VERSION=29.2
UNAME_S := $(shell uname -s)
UNAME_P := $(shell uname -p)

ARCH :=

ifeq ($(UNAME_S),Linux)
    ifeq ($(shell id -u),0)
        INSTALL_CMD = apt-get update && apt-get install -y protobuf-compiler
    else
        INSTALL_CMD = sudo apt-get install -y protobuf-compiler
    endif
    ARCH = linux-x86_64
endif

ifeq ($(UNAME_S),Darwin)
    ifeq ($(UNAME_P),arm)
        INSTALL_CMD = brew install protobuf
        ARCH = osx-universal_binary
    else
        INSTALL_CMD = sudo apt install -y protobuf-compiler
        ARCH = linux-x86_64
    endif
endif

all: generate lint test

.PHONY: lint
lint: .bin-deps
	echo 'Running linter on files...'
	$(GOLANGCI_BIN) run \
	--config=.golangci.yaml \
	--sort-results \
	--max-issues-per-linter=0 \
	--max-same-issues=0

.PHONY: test
test: .bin-deps
	echo 'Running tests...'
	${GO_TEST} ${GO_TEST_ARGS}

.PHONY: update
update:
	@if [ -n "$(shell git status --untracked-files=no --porcelain)" ]; then \
		echo 'You have some changes. Please commit, checkout or stash them.'; \
		exit 1; \
	fi
	@current_branch=$$(git branch --show-current); \
	echo "Current branch: $$current_branch"; \
	git checkout main; \
	git pull; \
	for branch in $$(git branch | sed 's/^[* ]*//'); do \
		git checkout $$branch; \
		if ! git rev-parse --symbolic-full-name @{u} >/dev/null 2>&1; then \
			branch_exists=$$(git ls-remote --heads origin $$branch); \
			if [ -n "$$branch_exists" ]; then \
				echo "Upstream exists for $$branch. Setting upstream to origin/$$branch."; \
				git branch --set-upstream-to=origin/$$branch; \
			else \
				echo "Upstream not found for $$branch. Pushing and setting upstream to origin/$$branch."; \
				git push --set-upstream origin $$branch; \
				git branch --set-upstream-to=origin/$$branch; \
			fi; \
		fi; \
		git pull --rebase; \
		git push -f; \
		git rebase main; \
		git push -f; \
	done; \
	git checkout $$current_branch; \
	echo 'Successfully updated'

.install-protoc:
	@command -v protoc >/dev/null 2>&1 || $(INSTALL_CMD)

bin-deps: .bin-deps

.generate-bin-deps: export GOBIN := $(LOCAL_BIN)
.generate-bin-deps: .create-bin .install-protoc
	GOBIN=$(LOCAL_BIN) go install github.com/easyp-tech/easyp/cmd/easyp@v0.7.11 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.18.1 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.19.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.2.1

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps: .generate-bin-deps
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5 && \
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/gotest@v0.0.6

.create-bin:
	rm -rf ./bin
	mkdir -p ./bin

generate: .generate-bin-deps .generate build
fast-generate: .generate

.generate:
	$(info Generating code...)

	rm -rf ./generated
	mkdir ./generated

	rm -rf ./docs/spec
	mkdir -p ./docs/spec

	rm -rf ~/.easyp/

	(PATH="$(PATH):$(LOCAL_BIN)" && $(EASYP_BIN) mod download && $(EASYP_BIN) generate)

	$(GOIMPORTS_BIN) -w .

build:
	go mod tidy
	go build -o ./bin/library ./cmd/library/
	cp ./bin/library ./integration-test/library

.PHONY: test-outbox
test-outbox: build
	docker exec postgres psql -U library -d library -c "TRUNCATE outbox, author, book, author_book RESTART IDENTITY CASCADE;"
	go test -v -race -timeout 9m -tags=outbox_hw -run '^TestOutbox$$' ./integration-test/outbox_hw/...

K6 ?= k6
K6_BASE_URL ?= http://localhost:8080
K6_VUS ?= 10
K6_DURATION ?= 30s

.PHONY: load-authors load-books load-read load-mixed load-all
load-authors:
	K6_BASE_URL=$(K6_BASE_URL) K6_VUS=$(K6_VUS) K6_DURATION=$(K6_DURATION) $(K6) run scripts/k6/create_authors.js

load-books:
	K6_BASE_URL=$(K6_BASE_URL) K6_VUS=$(K6_VUS) K6_DURATION=$(K6_DURATION) $(K6) run scripts/k6/create_books.js

load-read:
	K6_BASE_URL=$(K6_BASE_URL) K6_VUS=$(K6_VUS) K6_DURATION=$(K6_DURATION) $(K6) run scripts/k6/read_load.js

load-mixed:
	K6_BASE_URL=$(K6_BASE_URL) K6_VUS=$(K6_VUS) K6_DURATION=$(K6_DURATION) $(K6) run scripts/k6/mixed_load.js

load-all: load-authors load-books load-read load-mixed
