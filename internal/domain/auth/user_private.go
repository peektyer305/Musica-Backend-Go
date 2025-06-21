package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

// 　ユーザープライベートテーブルにユーザーテーブルを入れる
type UserPrivate struct {
	Id           valueobject.UserPrivateId
	UserId       valueobject.UserId
	Email        string
	UserName     string
	UserInfo     string
	UserIconUrl  string
	UserClientId string
	CreatedAt    time.Time
}
