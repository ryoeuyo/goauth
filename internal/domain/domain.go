package domain

import "time"

type User struct {
	Email      string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	LastSignIn time.Time `json:"last_sign_in"`
}
