package chat

import (
	"golang.org/x/net/context"
	_ "google.golang.org/protobuf/runtime/protoimpl"
	"log"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "hello from the server!"}, nil
}
