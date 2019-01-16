package main

import (
	"go_rest_api/app/config/db/mongo"
	"go_rest_api/app/server"
	"go_rest_api/app/service"
	"log"
)

const (
	address = ":3030"
)

func main() {
	ms, err := mongo.NewSession("mongodb://db:27017") // do not hard code the server address

	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	defer ms.Close()

	memberService := service.MemberServiceConstructor(ms.Copy(), "airline", "members")
	s := server.NewServer(memberService)
	s.Start()
}
