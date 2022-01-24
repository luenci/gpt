package content

// Makefile is the Makefile template used for new projects.
var Makefile = `GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u github.com/gin-gonic/gin

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o -tags=jsoniter -gcflags='-l=4' -ldflags='-s -w' {{ProjectName}} main.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t {{ProjectName}}:latest .
`
