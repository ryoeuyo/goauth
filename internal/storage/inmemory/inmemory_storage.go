package inmemory

import (
	"time"

	"github.com/ryoeuyo/goauth/internal/storage"
)

type Storage struct {
	cache []storage.User
}

func NewStorage() *Storage {
	return &Storage{
		cache: make([]storage.User, 0, 32),
	}
}

func (s *Storage) GetUser(email string) (*storage.User, error) {
	for i, u := range s.cache {
		if u.Email == email {
			s.cache[i].LastSignIn = time.Now()

			return &s.cache[i], nil
		}
	}

	return nil, storage.ErrEmailNotFound
}

func (s *Storage) SaveUser(user storage.User) (string, error) {
	for _, u := range s.cache {
		if u.Email == user.Email {
			return "", storage.ErrEmailIsExists
		}
	}
	s.cache = append(s.cache, user)

	return user.ID.String(), nil
}
