package handlers

import (
	"Musica-Backend/internal/application/auth"
	private "Musica-Backend/internal/presentation/rest/response/private"
	"context"
)

type UserPrivateHandler struct {
	UserPrivateUsecase *auth.UserPrivateUsecase
}

func (u *UserPrivateHandler) FindMe(ctx context.Context, email string) (private.UserPrivateResponse, error) {
	user, err := u.UserPrivateUsecase.FindMe(ctx, email)
	if err != nil {
		return private.UserPrivateResponse{}, err
	}
	return private.DomainToResponse(user), nil
}
