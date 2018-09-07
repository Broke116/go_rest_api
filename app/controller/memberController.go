package controller

import (
	"encoding/json"
	"fmt"
	"go_rest_api/app/model"
	"go_rest_api/app/service"
	"go_rest_api/app/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var result interface{}
var resultMessage string
var httpStatusCode int

type memberController struct {
	memberService *service.MemberService
}

// MemberControllerConstructor is used to initialize member api and define its routes.
func MemberControllerConstructor(ms *service.MemberService, r *mux.Router) *mux.Router {
	memberController := memberController{ms}

	r.HandleFunc("/memberStatus", memberController.GetStatus).Methods("GET")
	r.HandleFunc("/members", memberController.GetMembers).Methods("GET")
	r.HandleFunc("/insertMember", memberController.InsertMember).Methods("POST")
	r.HandleFunc("/updateMember/{id}", memberController.UpdateMember).Methods("PUT")
	r.HandleFunc("/member/{id}", memberController.GetMember).Methods("GET")
	r.HandleFunc("/member/{id}", memberController.DeleteMember).Methods("DELETE")

	return r
}

// GetStatus returns the status of the service
func (m *memberController) GetStatus(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)
	m.SendJSON(w, r, "{status:'ok'}", http.StatusOK)
}

// GetMembers method is used to return all members
func (m *memberController) GetMembers(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)
	members, err := m.memberService.GetMembers()

	if err != nil {
		model.CheckError(w, "Error when getting the members", http.StatusNotFound) // 404 status code
	}
	m.SendJSON(w, r, members, http.StatusOK)
}

// GetMember method is used to return a specific member
func (m *memberController) GetMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)

	member, err := m.memberService.GetMemberByID(id)
	if err != nil {
		model.CheckError(w, "Error when getting the member", http.StatusNotFound) // 404 status code
	}

	m.SendJSON(w, r, member, http.StatusOK)
}

// InsertMember endpoint is used to insert new members to the db
func (m *memberController) InsertMember(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)
	defer r.Body.Close()

	resultMessage = "success"

	var member model.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil { // decode body
		fmt.Println(err)
		model.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	member.StartDate = time.Now()

	if err := m.memberService.InsertMember(&member); err != nil {
		fmt.Println(err)
		model.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	m.SendJSON(w, r, resultMessage, http.StatusOK)
}

// UpdateMember method is used to update a specific member
func (m *memberController) UpdateMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	utils.Error(w, err, 500)

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		model.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	var member model.Member
	err = json.Unmarshal(b, &member)

	fmt.Println("member ", member)

	utils.Error(w, err, 500)

	if err := m.memberService.UpdateMember(&member, id); err != nil {
		model.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	resultMessage = "success"
	httpStatusCode = 200

	m.SendJSON(w, r, resultMessage, httpStatusCode)
}

// DeleteMember endpoint is used for deleting members
func (m *memberController) DeleteMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)

	if !bson.IsObjectIdHex(id) {
		model.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	if err := m.memberService.DeleteMember(id); err != nil {
		model.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	resultMessage = "success"
	httpStatusCode = 200

	m.SendJSON(w, r, resultMessage, httpStatusCode)
}

// SendJSON is used for returning a Json output
func (m *memberController) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Printf("Error while encoding JSON %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal Server Error"}`)
	}

	if v == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"result": "data not found"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}
