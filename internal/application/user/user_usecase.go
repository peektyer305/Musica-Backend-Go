package user

import (
	domain "Musica-Backend/internal/domain/user"
	valueobject "Musica-Backend/internal/domain/value_object"
	"context"
)

type UserUseCase struct {
	UserRepository domain.IUserRepository
}

func (u *UserUseCase) FindUserById(ctx context.Context, id valueobject.UserId) (domain.User, error) {
	return u.UserRepository.FindById(ctx, id)
}
