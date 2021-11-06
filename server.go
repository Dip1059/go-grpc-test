package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-grpc-test/chat"
	"go-grpc-test/protos"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	//proto test
	simpleProtoTest()

	//grpc test
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err.Error())
	}

	s := &chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}

func simpleProtoTest() {
	person := &protos.Person{
		Name: "Dipankar",
		Age: 29,
		Address: "Khulna BD",
	}
	data, err := proto.Marshal(person)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(data)

	person2 := &protos.Person{}
	err = proto.Unmarshal(data, person2)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(person2.GetName())
}
