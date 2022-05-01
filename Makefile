GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=api-test
TEST?="./"

default: build

build: validate
	@sh -c "./scripts/build"

ci: validate
	@sh -c "./scripts/ci"

package:
	@sh -c "./scripts/package"

docker-build: 
	@sh -c "./scripts/docker-build build"

docker-ci: 
	@sh -c "./scripts/docker-build ci"

validate: fmtcheck lint vet mod-verify

vet:
	@echo "==> Checking that code complies with go vet requirements..."
	@go vet $$(go list ./... | grep -v vendor/); if [ $$? -gt 0 ]; then \
		echo ""; \
		echo "go vet reported suspicious constructs. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review."; \
	fi

lint:
	@echo "==> Checking that code complies with golint requirements..."
	@GO111MODULE=off go get -u golang.org/x/lint/golint
	@if [ -n "$$(golint $$(go list ./...) | grep -v 'should have comment.*or be unexported' | tee /dev/stderr)" ]; then \
		echo ""; \
		echo "golint found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

mod-verify:
	@echo "==> Verifying go modules..."
	@go mod verify

bin:
	go build -o $(PKG_NAME)

fmt:
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "./scripts/go-fmt-check"

.PHONY: build vet lint fmt fmtcheck bin package mod-verify docker-ci docker-build
