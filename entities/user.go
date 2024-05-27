package entities

import (
	"context"
)

// Use pointers so that some structs are optional
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUseCase interface {
	Login(ctx context.Context, user User) (string, error)
}
