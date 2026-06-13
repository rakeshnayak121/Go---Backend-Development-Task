package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	DOB       string    `json:"dob" validate:"required"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
