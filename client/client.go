package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-client/chat"
	"log"
)

func main() {
	conn, err := grpc.Dial("grpc-server:9000", grpc.WithInsecure())
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
