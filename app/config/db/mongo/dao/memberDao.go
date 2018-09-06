package dao

import (
	"go_rest_api/app/model"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// MemberModel is the database representation of User object
type MemberModel struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Surname     string    `bson:"surname"`
	Email       string    `bson:"email"`
	Nationality string    `bson:"nationality"`
	Age         int       `bson:"age"`
	Gender      string    `bson:"gender"`
	Status      string    `bson:"status"`
	StartDate   time.Time `bson:"start_date"`
}

// MemberModels is an array of MemberModel
type MemberModels []MemberModel

// MemberModelIndex is used as an index.
func MemberModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewMemberModel is for converting the model data in which the data comes from the api, to mongodb model.
func NewMemberModel(m *model.Member) *MemberModel {
	return &MemberModel{
		ID:          m.ID,
		Name:        m.Name,
		Surname:     m.Surname,
		Email:       m.Email,
		Nationality: m.Nationality,
		Age:         m.Age,
		Gender:      m.Gender,
		Status:      m.Status,
		StartDate:   time.Now()}
}

// ToMember is a method which is used for getting data from the database and pushing it to the api used to show data to the client.
func (m *MemberModel) ToMember() *model.Member {
	return &model.Member{
		ID:          m.ID,
		Name:        m.Name,
		Surname:     m.Surname,
		Email:       m.Email,
		Nationality: m.Nationality,
		Age:         m.Age,
		Gender:      m.Gender,
		Status:      m.Status,
		StartDate:   time.Now()}
}
