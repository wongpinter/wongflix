package user

import (
	"time"

	"github.com/google/uuid"
)

type UUID = uuid.UUID

type User struct {
	ID        UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Public struct {
	ID    UUID
	Name  string
	Email string
}

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	ID    UUID
	Email string
}
