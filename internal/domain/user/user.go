package user

import (
	valueobject "Musica-Backend/internal/domain/value_object"

	"Musica-Backend/internal/domain/post"
)

type User struct {
	Id       valueobject.UserId
	Username string
	UserIconUrl *string
	UserInfo string
	Posts    []post.Post
	
}

