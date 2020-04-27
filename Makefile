
GO ?= go
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
GOFILES := find . -name "*.go" -type f -not -path "./vendor/*"

COVERFILE := coverage.txt

.PHONY: deps
deps:
	$(GO) mod download

.PHONY: clean
clean:
	$(GO) clean -modcache -cache -i
	rm -rf ./build

.PHONY: fmt
fmt:
	$(GOFILES) | xargs $(GOFMT) -w

.PHONY: fmt-check
fmt-check:
	@files=$$($(GOFILES) | xargs $(GOFMT) -l); if [ -n "$$files" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${files}"; \
		exit 1; \
		fi;

.PHONY: test
test: fmt-check
	$(GO) test -v -coverprofile=$(COVERFILE) ./...

coverage: test
	$(GO) tool cover -html=$(COVERFILE)

e2e:
	env $$(cat .env) env ENV=test env E2E=1 $(GO) test -run ^TestE2EApp

.PHONY: swagger
swagger:
	swag init -g ./cmd/api/main.go

.PHONY: build
build: 
	CGO_ENABLED=0 $(GO) build -installsuffix 'static' -ldflags "-X main.Version=`date -u +1.%Y%m%d.%H%M%S`" -o ./build/api ./cmd/api

.PHONY: run
run:
	make build
	env $$(cat .env) ./build/api
