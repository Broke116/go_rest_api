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

	/*mc := &controller.MemberController{}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/", mc.GetStatus).Methods("GET")
	s.HandleFunc("/members", mc.GetMembers).Methods("GET")
	s.HandleFunc("/member/{id}", mc.GetMember).Methods("GET")
	s.HandleFunc("/insertMember", mc.InsertMember).Methods("POST")
	s.HandleFunc("/updateMember", mc.UpdateMember).Methods("POST")
	s.HandleFunc("/deleteMember/{id}", mc.DeleteMember).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(address, nil))*/
}
