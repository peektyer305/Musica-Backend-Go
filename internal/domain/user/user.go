package user

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"net/url"

	domain "Musica-Backend/internal/domain/post"
)

type User struct {
	Id       valueobject.UserId
	Username string
	UserIconUrl *url.URL
	UserInfo string
	Posts    []domain.Post
	
}

