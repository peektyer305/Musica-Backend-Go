package auth

import (
	"time"

	"Musica-Backend/internal/domain/auth"

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

func DomainToResponse(user auth.UserPrivate) UserPrivateResponse {
	var id = user.Id.GetUUID()
	var userId = user.UserId.GetUUID()

	return UserPrivateResponse{
		Id:           id,
		UserId:       userId,
		Email:        user.Email,
		UserName:     user.UserName,
		UserInfo:     user.UserInfo,
		UserIconUrl:  user.UserIconUrl,
		UserClientId: user.UserClientId,
		CreatedAt:    user.CreatedAt,
	}
}
