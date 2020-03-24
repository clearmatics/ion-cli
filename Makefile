.DEFAULT_GOAL := build
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PACKAGES = $(shell find ./ -type d -not -path "./keystore" -not -path "./vendor" -not -path "./vendor/*" -not -path "./.github" -not -path "./.github/*" -not -path "./.git" -not -path "./.git/*" -not -path "./config/files")
SHELL=/bin/bash
TARGET_BINARY=ion-cli

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
HASH             := $(shell git rev-parse HEAD)
VERSION_FLAGS    := -ldflags='-X "main.BuildVersion=$(VERSION)" -X "main.BuildDateTime=$(DATE)" -X "main.CommitHash=$(HASH)"'


documentation: build
	@$(TARGET_BINARY) --docgen

clean:
	@rm -f ion-cli

test:
	@go test ./... -v -short

integration-test:
	@go test ./... -v

format:
	@gofmt -s -w ${PACKAGES}

modules: 
	@go mod download
	@go get 

build: modules
	# Build project
	@go build $(VERSION_FLAGS) -o $(TARGET_BINARY)

check:
	@if [ -n "$(shell gofmt -l ${PACKAGES})" ]; then \
		echo 1>&2 'The following files need to be formatted:'; \
		gofmt -l .; \
		exit 1; \
		fi

vet:
	@go vet ${PACKAGES}

lint:
	@golangci-lint run -E gosec -E gofmt --deadline 4m0s $(PACKAGES)

coverage: modules 
	echo "mode: count" > coverage-all.out 
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile coverage.out -covermode count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html coverage-all.out
