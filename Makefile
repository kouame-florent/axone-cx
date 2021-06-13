.DEFAULT_GOAL := gen

build_gen:
	go build -o bin/ builder/*
.PHONY:build_gen

gen: build_gen
	go generate generator/*
.PHONY:gen

compile:
	protoc -I=./api/ --go_out=./api --go-grpc_out=./api ./api/axone.proto