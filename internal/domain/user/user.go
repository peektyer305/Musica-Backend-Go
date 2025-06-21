package user

import (
	valueobject "Musica-Backend/internal/domain/value_object"

	domain "Musica-Backend/internal/domain/post"
)

type User struct {
	Id            valueobject.UserId
	UserPrivateId valueobject.UserPrivateId // ユーザープライベートID
	Email         string
	Username      string
	UserIconUrl   *string
	UserInfo      string
	UserClientId  string // ユーザークライアントID
	Posts         []domain.Post
}
