PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.PHONY: .install-oapi-codegen
.install-oapi-codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: .oapi-codegen-candidate
.oapi-codegen-candidate:
	oapi-codegen -generate types,gin -package candidates -o api/swagger/transfers/swagger_transfers.go api/swagger/transfers/swagger_transfers.yaml

# open api golang
# API FIRST
# CODE FIRST