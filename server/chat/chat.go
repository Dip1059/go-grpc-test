package chat

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"log"
	"os"
)

func DBconnect() *gorm.DB {
	var db *gorm.DB
	var err error
	if os.Getenv("DB_DIALECT") == "mysql" {
		db, err = gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+
			os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?parseTime=true")
	} else if os.Getenv("DB_DIALECT") == "postgres" {
		db, err = gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USERNAME")+
			" dbname="+os.Getenv("DB_NAME")+" password="+os.Getenv("DB_PASSWORD")+" sslmode=disable")
	}
	if err != nil {
		log.Println("log", err.Error())
		return db
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
