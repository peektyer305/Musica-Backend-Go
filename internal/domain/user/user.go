package user

import (
	valueobject "Musica-Backend/internal/domain/value_object"
)

type User struct {
	Id       valueobject.UserId
	Username string
	Email    string
	Password string
}

