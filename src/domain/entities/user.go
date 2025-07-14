package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}

// NewUser is a factory function that validates input and creates a new User.
func NewUser(name string, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, errors.New("user name is required")
	}
	if email == "" {
		return nil, errors.New("user email is required")
	}

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}, nil
}
