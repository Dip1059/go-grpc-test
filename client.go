package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"go-grpc-test/chat"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()

	chatClient := chat.NewChatServiceClient(conn)

	message, err := chatClient.SayHello(context.Background(), &chat.Message{Body: "Hello server"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Message from server: %s", message.Body)
}
