package main

import (
	"context"
	"github.com/michaelwp/go-gRpc/client/chat"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Error close connection: %s", err)
		}
	}()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{Body: "Hello from the client"}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling say hello: %s", err)
	}

	log.Printf("Response from server: %s", response)
}
