package auth

import (
	domain "Musica-Backend/internal/domain/auth"
	"context"
)

type UserPrivateUsecase struct {
	AuthRepository domain.IAuthRepository
}

func (u *UserPrivateUsecase) FindMe(ctx context.Context, email string) (domain.User, error) {
	return u.AuthRepository.FindMe(ctx, email)
}
