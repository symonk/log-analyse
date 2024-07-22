MAIN_PACKAGE_PATH := ./cmd/
BINARY_NAME := log-analyse

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: no-dirty
no-dirty:
	git diff --exit-code

# ==================================================================================== #
# Quality 
# ==================================================================================== #

## tidy: format the code and tidy the mod file
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go mod vet ./...
	go test -race -buildvcs -vet=off ./...

# ==================================================================================== #
# Development 
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## cover: run all tests and display the coverage
.PHONY: cover
cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the binary
.PHONY: build
build:
	go build -o=dist/${BINARY_NAME} .

## run: run the built binary (no args)
.PHONY: run
run: build
	dist/${BINARY_NAME}


# ==================================================================================== #
# Operations 
# ==================================================================================== #

## upgrade: upgrade go dependencies
.PHONY: upgrade
upgrade:
	go get -u ./...
