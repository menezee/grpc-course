# gRPC Golang

[gRPC [Golang] Master Class: Build Modern API and Microservices](https://www.oreilly.com/library/view/grpc-golang-master/9781838555467/)

### Requirements
* [Go](https://golang.org/doc/install)

### Setup
```shell script
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc
```

### How to run
First, generate the protobuffers:
```
sh generate.sh
```
Then, run the server/client:
```
go run greet/greet_server/server.go & go run greet/greet_client/client.go
```