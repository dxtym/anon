package models

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty" validate:"required,alphanum,min=5,max=15"`
	Password  string    `json:"password,omitempty" validate:"required,alphanum,min=8"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
