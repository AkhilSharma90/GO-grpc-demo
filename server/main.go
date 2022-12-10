package main

import (
	"log"
	"net"

	pb "github.com/akhil/grpc-demo/proto"
	"google.golang.org/grpc"
)

//define the port
const (
	port = ":8080"
)

//this is the struct to be created, pb is imported upstairs
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
