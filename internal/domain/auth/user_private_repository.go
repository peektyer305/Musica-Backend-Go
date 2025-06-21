package auth

import "context"

type IUserPrivateRepository interface {
	FindMe(ctx context.Context, email string) (UserPrivate, error)
}
