package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

type User struct {
	Id           valueobject.UserId
	UserName     string
	UserInfo     string
	UserIconUrl  string
	UserClientId string // ユーザークライアントID
	Email        string
	CreatedAt    time.Time
}
