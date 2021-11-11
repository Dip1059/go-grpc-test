package chat

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"log"
)

func DBconnect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=itech dbname=chat_db password=123456 sslmode=disable")
	if err != nil {
		log.Println("log", err.Error())
		return nil
	}
	return db
}

type Server struct {
}

func (s *Server) Signup(ctx context.Context, user *User) (*User, error) {
	db := DBconnect()
	if err := db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Message from Client: %s", message.Body)
	return &Message{Body: "Hello Client"}, nil
}
