package models

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty" validate:"required,alphanum,min=5,max=15"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,alphanum,min=8"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
