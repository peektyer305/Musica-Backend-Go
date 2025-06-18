package auth

type IUserRepository interface {
	FindByEmail(email string) (*User, error)
}
