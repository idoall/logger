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

proto:
	for d in {api,srv}; do \
		for f in $$d/**/**/proto/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done


	for d in {grpc,1}; do \
		for f in $$d/**/**/proto/*.proto; do \
		    protoc -I/usr/local/include -I. \
              -I${GOPATH}/src \
              -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
              --go_out=plugins=grpc:. $$f; \
              echo generates gRPC code compiled: $$f; \
			protoc -I/usr/local/include -I. \
              -I${GOPATH}/src \
              -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
              --grpc-gateway_out=logtostderr=true:. $$f; \
              echo generates gRPC reverse-proxy code compiled: $$f; \
		done \
	done

build:
	./bin/build.sh
	./bin/ci.build.image.sh registry.cn-beijing.aliyuncs.com/mshk/go-exchange

push:
	./bin/ci.push.image.sh registry.cn-beijing.aliyuncs.com/mshk/go-exchange

dbproxy:
	./bin/replace.dbproxy.host.sh

run:
# 启动 Mysql 代理
	docker-compose up

runnewweb:
# 启动 Mysql 代理
	docker-compose -f docker-compose.newweb.yml up

rund:
	docker-compose up -d

down:
	docker-compose down


clear:
	./bin/clear.sh