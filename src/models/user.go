package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `json:"id"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"password"`
}
