# Basic Go gRPC Server and Client

This is a basic gRPC server and client written in Go. It is based on the [gRPC Quickstart](https://grpc.io/docs/quickstart/go.html) and [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) tutorials.

We have implemented a simple gRPC server and client with the following functionality:
- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

# Setting up a gRPC-Go project
1. Create a new directory for your project and cd into it

```bash
mkdir basic-go-grpc
cd basic-go-grpc
mkdir client server proto
```

2. Installing the gRPC Go plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Initialize a Go module

```bash
go mod init github.com/your_username/basic-go-grpc

go mod tidy
```

4. Create the proto file with the required services and messages in the proto directory

5. Generate .pb.go files from the proto file

depending on what path you mention in your greet.proto file, you will either run this - 

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```
OR this -

```bash
protoc --go_out=. --go_opt=module=github.com/akhil/basic-go-grpc --go-grpc_out=. --go-grpc_opt=module=githu
b.com/akhil/basic-go-grpc proto/greet.proto
```

6. Create the server and client directories and create the main.go files with necessary controllers and services


# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```