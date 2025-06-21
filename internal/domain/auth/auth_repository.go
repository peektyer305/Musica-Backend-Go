package auth

import "context"

type IAuthRepository interface {
	FindMe(ctx context.Context, email string) (User, error)
}
