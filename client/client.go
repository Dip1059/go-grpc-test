package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-client/chat"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Success template.HTML
	Fail    template.HTML
}

var (
	Msg Message
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		return
	}
	engine := gin.Default()
	engine.Static("/assets", "./")
	engine.LoadHTMLGlob("views/**/*.gohtml")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.gohtml", map[string]interface{}{"Msg": Msg})
		Msg.Success = ""
		Msg.Fail = ""
	})
	engine.POST("/signup", func(c *gin.Context) {
		grpcConn, chatClient := getChatClient()
		defer grpcConn.Close()
		user, err := chatClient.Signup(context.Background(), &chat.User{Name: c.PostForm("name"), Email: c.PostForm("email")})
		if err != nil {
			log.Println(err.Error())
			Msg.Fail = "Sign up failed. Something went wrong."
			c.Redirect(http.StatusFound, c.Request.Referer())
			return
		}
		log.Println(user)

		message, err := chatClient.SayHello(context.Background(), &chat.Message{UserId: user.Id, Body: "Hello server, I am " + user.Name})
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println(message)
		}

		Msg.Success = "Successfully signed up."
		c.Redirect(http.StatusFound, c.Request.Referer())
		return
	})
	err = engine.Run(":9020")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func getChatClient() (*grpc.ClientConn, chat.ChatServiceClient) {
	creds, err := credentials.NewClientTLSFromFile("server-cert.pem", "")
	if err != nil {
		log.Println(err.Error())
	}
	conn, err := grpc.Dial(os.Getenv("CHAT_SERVER_HOST")+":"+os.Getenv("CHAT_SERVER_PORT"), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println(err.Error())
	}
	return conn, chat.NewChatServiceClient(conn)
}
