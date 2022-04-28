GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=rancher2
TEST?="./"

default: build

build: validate
	@sh -c "./scripts/build"

docker-build: 
	@sh -c "'./scripts/docker-build'"

validate: fmtcheck lint vet

test: fmtcheck
	@echo "==> Running testing..."
	go test ./ || exit 1
	echo ./ | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
vet:
	@echo "==> Checking that code complies with go vet requirements..."
	@go vet $$(go list ./... | grep -v vendor/); if [ $$? -gt 0 ]; then \
		echo ""; \
		echo "WARNING!! Expected vet reported construct:"; \
		echo "rancher2/schema_secret_v2.go:20:2: struct field Type repeats json tag \"type\" also at ../../../../github.com/rancher/norman@v0.0.0-20210225010917-c7fd1e24145b/types/types.go:66"; \
		echo "";\
		echo "If vet reported more suspicious constructs, please check and"; \
		echo "fix them if necessary, before submitting the code for review."; \
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

bin:
	go build -o $(PROVIDER_NAME)

fmt:
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/go-fmt-check'"

.PHONY: build vet fmt fmtcheck bin
