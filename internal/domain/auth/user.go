package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
)

type User struct {
	ID    valueobject.UserId `json:"id"`
	Email string             `json:"email"`
	Name  string             `json:"name"`
}
