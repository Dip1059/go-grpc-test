package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-client/chat"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "./")
	engine.LoadHTMLGlob("views/**/*.gohtml")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.gohtml", nil)
	})
	engine.POST("/signup", func(c *gin.Context) {
		chatClient := getChatClient()
		user, err := chatClient.Signup(context.Background(), &chat.User{Name: c.PostForm("name"), Email: c.PostForm("email")})
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(user)

		message, err := chatClient.SayHello(context.Background(), &chat.Message{UserId: user.Id, Body: "Hello server, I am " + user.Name})
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(message)
	})
	err := engine.Run(":9020")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func getChatClient() chat.ChatServiceClient {
	creds, err := credentials.NewClientTLSFromFile("server-cert.pem", "")
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := grpc.Dial("grpc-server:9000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println(err.Error())
	}
	return chat.NewChatServiceClient(conn)
}
