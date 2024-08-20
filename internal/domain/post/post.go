package post

import (
	"net/url"

	valueobject "Musica-Backend/internal/domain/value_object"
)

type Post struct {
	Id valueobject.PostId
	UserId valueobject.UserId
	Title string
	Content *string
	MusicUrl url.URL
	ImageUrl *url.URL
	UserIconUrl *url.URL
	UserName string
}
