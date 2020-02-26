.PHONY: proto build

GCTPKG = $(shell go list -e)
LINTPKG = github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21.0
LINTBIN = $(GOPATH)/bin/golangci-lint
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")

# 设置 golang 全局代理
export GOPROXY=https://mirrors.aliyun.com/goproxy/,https://goproxy.cn,https://proxy.golang.org/

get:
	GO111MODULE=on go get $(GCTPKG)

linter:
	GO111MODULE=on go get $(LINTPKG)
	golangci-lint run --verbose --skip-dirs vendor | tee /dev/stderr

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic  ./...

update_deps:
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy
	rm -rf vendor
	GO111MODULE=on go mod vendor

fmt:
	gofmt -l -w -s $(shell find . -path './vendor' -prune -o -type f -name '*.go' -print)
