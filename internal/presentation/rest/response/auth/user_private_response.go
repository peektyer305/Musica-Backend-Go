package auth

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserPrivateResponse struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	Email        string
	UserName     string
	UserInfo     string
	UserIconUrl  string
	UserClientId string
	CreatedAt    time.Time
}

func DomainToResponse(user UserPrivate) UserPrivateResponse
