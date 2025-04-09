package post

import (
	"time"

	valueobject "Musica-Backend/internal/domain/value_object"
)

type Post struct {
	Id valueobject.PostId
	UserId valueobject.UserId
	Title string
	Content *string
	MusicUrl string
	ImageUrl *string
	UserIconUrl *string
	UserName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
