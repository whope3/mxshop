GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: generate
generate:
ifeq ($(type), api)
	cd app/apis/cmd/$(name)/desc &&\
    goctl api go -api *.api -dir ../  --style=goZero
else ifeq ($(type), rpc)
	cd app/$(name)/rpc/pb &&\
    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero -m
endif


.PHONY: generate_all
generate_all:
	cd app/apis/cmd/app-api/desc &&\
    goctl api go -api *.api -dir ../  --style=goZero &&\
    cd - &&\
	cd app/apis/cmd/external-api/desc &&\
    goctl api go -api *.api -dir ../  --style=goZero &&\
    cd - &&\
	cd app/video/rpc/pb &&\
    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero -m &&\
    cd - &&\
	cd app/usercenter/rpc/pb &&\
    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero -m &&\
    cd - &&\
    cd app/marketing/rpc/pb &&\
    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero -m &&\
    cd - &&\
    cd app/reward/rpc/pb &&\
    goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero -m &&\
    echo "finish"

.PHONY: run
run:
ifeq ($(type), api)
	cd app/apis/cmd/$(name) && go run .
else ifeq ($(type), rpc)
	cd app/$(name)/rpc && go run .
endif

.PHONY: air
air:
ifeq ($(type), api)
	cd app/apis/cmd/$(name) && air
else ifeq ($(type), rpc)
	cd app/$(name)/rpc && air
endif

.PHONY: build
build:
ifeq ($(type), api)
	mkdir -p bin/ && go build -o bin/$(name)  app/apis/cmd/$(name)/app.go
else ifeq ($(type), rpc)
	mkdir -p bin/ && go build -o bin/$(name)  app/$(name)/rpc/$(name).go
endif
