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

	s.HandleFunc("/members", mc.GetMembers).Methods("GET")
	s.HandleFunc("/member/{id}", mc.GetMember).Methods("GET")
	s.HandleFunc("/updateMember", mc.UpdateMember).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(address, nil))
}
