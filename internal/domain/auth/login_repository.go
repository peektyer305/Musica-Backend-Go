package auth

type ILoginRepository interface {
	FindByEmail(email string) (*User, error)
}
