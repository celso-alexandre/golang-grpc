package main

import (
	"context"
	"log"

	"github.com/celso-alexandre/golang-grpc/protobuf/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From the Client!"})
	if err != nil {
		panic(err)
	}
	log.Printf("Response from Server: %s", response.Body)
}
