package user

import (
	"context"
)

type IUserRepository interface {
	FindById(ctx context.Context) (User, error)
}