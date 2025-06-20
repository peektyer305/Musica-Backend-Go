package auth

import valueobject "Musica-Backend/internal/domain/value_object"

type LoginResponse struct {
	Token  string                    `json:"token"`
	Email  string                    `json:"email"`
	Name   string                    `json:"name"`
	ID     valueobject.UserPrivateId `json:"id"`
	UserId valueobject.UserId        `json:"user_id"`
}
