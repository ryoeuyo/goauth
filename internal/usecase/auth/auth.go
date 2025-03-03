package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ryoeuyo/goauth/internal/storage"
	"github.com/ryoeuyo/goauth/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	st     storage.UserStorage
	ttl    time.Duration
	secret string
}

func New(
	st storage.UserStorage,
	ttl time.Duration,
	secret string,
) *UseCase {
	return &UseCase{
		st:     st,
		ttl:    ttl,
		secret: secret,
	}
}

func (us *UseCase) Login(ctx context.Context, email, password string) (string, error) {
	const fn = "auth.Login"

	user, err := us.st.GetUser(email)
	if err != nil {
		if errors.Is(err, storage.ErrEmailNotFound) {
			return "", fmt.Errorf("%s:%w", fn, usecase.ErrUserNotFound)
		}

		return "", fmt.Errorf("%s:%w", fn, err)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(password),
	); err != nil {
		return "", fmt.Errorf("%s:%w", fn, usecase.ErrInvalidCredentials)
	}

	token, err := NewToken(user, us.ttl, us.secret)
	if err != nil {
		return "", fmt.Errorf("%s:%w", fn, err)
	}

	return token, nil
}

func (us *UseCase) Register(ctx context.Context, email, password string) (string, error) {
	const fn = "auth.Register"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%s:%w", fn, err)
	}

	id, err := us.st.SaveUser(storage.User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: string(hashedPass),
		CreatedAt:      time.Now(),
	})

	if err != nil {
		if errors.Is(err, storage.ErrEmailIsExists) {
			return "", fmt.Errorf("%s:%w", fn, usecase.ErrUserIsExists)
		}
	}

	return id, err
}
