package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var port = ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	fmt.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
