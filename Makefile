MAIN_PACKAGE_PATH := ./cmd/
BINARY_NAME := log-analyse

# ==========
# [Helpers]
# ==========

.PHONY: help
help:
	@echo 'Usage:''
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


.PHONY: no-dirty
no-dirty:
	git diff --exit-code

# ==========
# [Quality]
# ==========

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit:
	go mod verify
	go mod vet ./...
	go test -race -buildvcs -vet=off ./...

# ==============
# [Development]
# ==============

.PHONY: test
test:
	go test -v -race -buildvcs ./...

## cover: run all tests and display the coverage
.PHONY: cover
cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: build
build:
	go build -o=dist/${BINARY_NAME} .

.PHONY: run
run: build
	dist/${BINARY_NAME}


