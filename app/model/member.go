package model

import (
	"time"
)

// Member struct is defined for holding the member data
type Member struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	Nationality string    `json:"nationality"`
	Age         int       `json:"age"`
	Gender      string    `json:"gender"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
}

// Members holds multiple member values
type Members []Member
