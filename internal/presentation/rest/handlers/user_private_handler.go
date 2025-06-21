package handlers

import (
	"Musica-Backend/internal/application/auth"
	"context"
)

type UserPrivateHandler struct {
	UserPrivateUsecase auth.UserPrivateUsecase
}

func (u *UserPrivateHandler) FindMe(ctx context.Context, email string) (auth.User, error) {
	return u.UserPrivateUsecase.FindMe(ctx, email)
}
