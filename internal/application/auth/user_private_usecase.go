package auth

import (
	domain "Musica-Backend/internal/domain/auth"
	"context"
)

type UserPrivateUsecase struct {
	UserPrivateRepository domain.IUserPrivateRepository
}

func (u *UserPrivateUsecase) FindMe(ctx context.Context, email string) (domain.UserPrivate, error) {
	return u.UserPrivateRepository.FindMe(ctx, email)
}
