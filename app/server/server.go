package server

import (
	"go_rest_api/app/controller"
	"go_rest_api/app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server holds the values for server instantination
type Server struct {
	router *mux.Router
}

// NewServer is used to create a server
func NewServer(ms *service.MemberService) *Server {
	s := Server{router: mux.NewRouter()}
	controller.MemberControllerConstructor(ms, s.newSubRouter("/api/v1"))
	return &s
}

// Start is used to start a server
func (s *Server) Start() {
	//flag.Parse()
	log.Println("Listening on port 3030")

	if err := http.ListenAndServe(":3030", s.router); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}

}

func (s *Server) newSubRouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
