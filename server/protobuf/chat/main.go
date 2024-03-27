package chat

import (
	"context"
	"log"
)

type Server struct{}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer.
func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	panic("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received: %v", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
