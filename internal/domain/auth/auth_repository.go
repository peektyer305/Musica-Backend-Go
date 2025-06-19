package auth

type IAuthRepository interface {
	FindByEmail(email string) (*User, error)
}
