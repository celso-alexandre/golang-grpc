package main

import (
	"fmt"
	"net"

	"github.com/celso-alexandre/golang-grpc/server/protobuf/chat"
	"google.golang.org/grpc"
)

func main() {
	var port = ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	chatServer := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	fmt.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
