# Todo 2: Implement gRPC in our app 

> Goal: Change reliance on REST api to gRPC

```console
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Install protoc
```console
apt install -y protobuf-compiler
protoc --version

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```

1. Define protobuf
2. Compile
3. Write server code
4. Write client code


