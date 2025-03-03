package storage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	LastSignIn     time.Time
}

type UserStorage interface {
	GetUser(email string) (*User, error)
	SaveUser(user User) (string, error)
}
