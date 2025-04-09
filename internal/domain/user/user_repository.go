package user

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"context"
)

type IUserRepository interface {
	FindById(ctx context.Context, id valueobject.UserId) (User, error)
}