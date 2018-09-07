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
)

var result interface{}
var resultMessage string
var httpStatusCode int

/*var members = []model.Member{
	model.Member{ID: "1", Name: "Ekin", Surname: "Yucel", Email: "sad@asd.com", Nationality: "TUR", Age: 23, Gender: "M", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "2", Name: "Susan", Surname: "Doe", Email: "sad@asd.com", Nationality: "TUR", Age: 24, Gender: "F", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "3", Name: "Alberto", Surname: "Yucel", Email: "sad@asd.com", Nationality: "BRA", Age: 30, Gender: "M", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "4", Name: "Jane", Surname: "Doe", Email: "sad@asd.com", Nationality: "TUR", Age: 26, Gender: "F", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "5", Name: "John", Surname: "Doe", Email: "sad@asd.com", Nationality: "ENG", Age: 28, Gender: "M", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "6", Name: "Test", Surname: "Test", Email: "sad@asd.com", Nationality: "JPN", Age: 23, Gender: "F", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "7", Name: "Mirel", Surname: "Souza", Email: "sad@asd.com", Nationality: "TUR", Age: 25, Gender: "F", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "8", Name: "Donald", Surname: "Duck", Email: "sad@asd.com", Nationality: "ITA", Age: 50, Gender: "F", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "9", Name: "Tekin", Surname: "Acar", Email: "sad@asd.com", Nationality: "TUR", Age: 18, Gender: "M", Status: "AC", StartDate: time.Now()},
	model.Member{ID: "10", Name: "Roberto", Surname: "Soldado", Email: "sad@asd.com", Nationality: "SPA", Age: 24, Gender: "F", Status: "AC", StartDate: time.Now()},
}*/

// MemberController is a definition for member controller
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
	utils.Log(r.Method, r.URL)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	utils.Error(w, err, 500)

	var member model.Member
	err = json.Unmarshal(b, &member)

	utils.Error(w, err, 500)

	/*for i, m := range members {
		if m.ID == member.ID {
			members[i].Name = member.Name
			members[i].Surname = member.Surname
			members[i].Gender = member.Gender
			members[i].Age = member.Age
			members[i].Nationality = member.Nationality
			resultMessage = "success"
			httpStatusCode = 200
			break
		} else {
			resultMessage = "not success"
			httpStatusCode = 404
		}
	}*/

	resultMessage = "not success"
	httpStatusCode = 404

	m.SendJSON(w, r, resultMessage, httpStatusCode)
}

// DeleteMember endpoint is used for deleting members
func (m *memberController) DeleteMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)
	fmt.Println("id ", id)

	/*for i, m := range members {
		if m.ID == string(id) {
			members = append(members[:i], members[i+1:]...)
			resultMessage = "success"
			httpStatusCode = 200
			break
		} else {
			resultMessage = "member not found"
			httpStatusCode = 404
		}
	}*/

	resultMessage = "member not found"
	httpStatusCode = 404

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
