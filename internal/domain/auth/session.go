package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

type Session struct {
	ID         valueobject.SessionId `json:"id"`
	UserId     valueobject.UserId    `json:"user_id"`
	Expiration time.Time             `json:"expiration"`
}
