package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-client/chat"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server-cert.pem", "")
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := grpc.Dial("grpc-server:9000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()

	chatClient := chat.NewChatServiceClient(conn)

	message, err := chatClient.SayHello(context.Background(), &chat.Message{Body: "Hello server"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Message from Server: %s", message.Body)
}
