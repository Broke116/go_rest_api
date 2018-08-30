package main

import (
	"go_rest_api/app/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	address = ":8080"
)

func main() {
	mc := &controller.MemberController{}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/", mc.GetStatus).Methods("GET")
	s.HandleFunc("/members", mc.GetMembers).Methods("GET")
	s.HandleFunc("/member/{id}", mc.GetMember).Methods("GET")
	s.HandleFunc("/insertMember", mc.InsertMember).Methods("POST")
	s.HandleFunc("/updateMember", mc.UpdateMember).Methods("POST")
	s.HandleFunc("/deleteMember/{id}", mc.DeleteMember).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(address, nil))
}
