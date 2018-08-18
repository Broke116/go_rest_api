package controller

import (
	"encoding/json"
	"fmt"
	"go_rest_api/app/model"
	"go_rest_api/app/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var result interface{}
var resultMessage string
var httpStatusCode int

var members = []model.Member{
	model.Member{ID: "1", Name: "Ekin", Surname: "Yucel", Nationality: "TUR", Age: 23, Gender: "M"},
	model.Member{ID: "2", Name: "Susan", Surname: "Doe", Nationality: "TUR", Age: 24, Gender: "F"},
	model.Member{ID: "3", Name: "Alberto", Surname: "Yucel", Nationality: "BRA", Age: 30, Gender: "M"},
	model.Member{ID: "4", Name: "Jane", Surname: "Doe", Nationality: "TUR", Age: 26, Gender: "F"},
	model.Member{ID: "5", Name: "John", Surname: "Doe", Nationality: "ENG", Age: 28, Gender: "M"},
	model.Member{ID: "6", Name: "Test", Surname: "Test", Nationality: "JPN", Age: 23, Gender: "F"},
	model.Member{ID: "7", Name: "Mirel", Surname: "Souza", Nationality: "TUR", Age: 25, Gender: "F"},
	model.Member{ID: "8", Name: "Donald", Surname: "Duck", Nationality: "ITA", Age: 50, Gender: "F"},
	model.Member{ID: "9", Name: "Tekin", Surname: "Acar", Nationality: "TUR", Age: 18, Gender: "M"},
	model.Member{ID: "10", Name: "Roberto", Surname: "Soldado", Nationality: "SPA", Age: 24, Gender: "F"},
}

// MemberController is a definition for member controller
type MemberController struct {
	Controller
}

// GetMembers method is used to return all members
func (m *MemberController) GetMembers(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)
	m.SendJSON(w, r, members, http.StatusOK)
}

// GetMember method is used to return a specific member
func (m *MemberController) GetMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)

	result = nil

	for _, member := range members {
		if member.ID == string(id) {
			result = member
		}
	}

	m.SendJSON(w, r, result, http.StatusOK)
}

// InsertMember endpoint is used to insert new members to the db
func (m *MemberController) InsertMember(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	utils.Error(w, err, 500)

	var member model.Member
	err = json.Unmarshal(b, &member)
	member.ID = strconv.Itoa(len(members) + 1)

	utils.Error(w, err, 500)

	members = append(members, member)
	resultMessage = "success"

	m.SendJSON(w, r, resultMessage, http.StatusOK)
}

// UpdateMember method is used to update a specific member
func (m *MemberController) UpdateMember(w http.ResponseWriter, r *http.Request) {
	utils.Log(r.Method, r.URL)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	utils.Error(w, err, 500)

	var member model.Member
	err = json.Unmarshal(b, &member)

	utils.Error(w, err, 500)

	for i, m := range members {
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
	}

	m.SendJSON(w, r, resultMessage, httpStatusCode)
}

// DeleteMember endpoint is used for deleting members
func (m *MemberController) DeleteMember(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	utils.Log(r.Method, r.URL)
	fmt.Println("id ", id)

	for i, m := range members {
		if m.ID == string(id) {
			members = append(members[:i], members[i+1:]...)
			resultMessage = "success"
			httpStatusCode = 200
			break
		} else {
			resultMessage = "member not found"
			httpStatusCode = 404
		}
	}

	m.SendJSON(w, r, resultMessage, httpStatusCode)
}
