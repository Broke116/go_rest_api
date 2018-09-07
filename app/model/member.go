package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Member struct is defined for holding the member data
type Member struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Surname     string        `json:"surname"`
	Email       string        `json:"email"`
	Nationality string        `json:"nationality"`
	Age         int           `json:"age"`
	Gender      string        `json:"gender"`
	Status      string        `json:"status"`
	StartDate   time.Time     `json:"start_date"`
}

// Members holds multiple member values
type Members []Member
