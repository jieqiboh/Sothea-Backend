package usecases

import (
	"context"
	"github.com/jieqiboh/sothea_backend/controllers/middleware"
	"github.com/jieqiboh/sothea_backend/entities"
	"time"
)

type loginUsecase struct {
	contextTimeout time.Duration
	secretKey      []byte
}

// NewLoginUseCase
func NewLoginUseCase(timeout time.Duration) entities.LoginUseCase {
	return &loginUsecase{
		contextTimeout: timeout,
	}
}

func (l *loginUsecase) Login(ctx context.Context, user entities.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	if user.Username == "admin" && user.Password == "admin" { // Todo: replace this with a database query
		token, err := middleware.CreateToken(user.Username)
		return token, err
	} else {
		return "", entities.ErrLoginFailed
	}
}
