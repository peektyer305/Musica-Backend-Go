package user

import (
	domain "Musica-Backend/internal/domain/user"
	"context"
)

type UserUseCase struct {
	UserRepository domain.IUserRepository
}

func (u *UserUseCase) FindUserById(ctx context.Context) (domain.User, error) {
	return u.UserRepository.FindById(ctx)
}
