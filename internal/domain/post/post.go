package post

import (
	"net/url"
)

type Post struct {
	Id PostId
	UserId UserId
	Title string
	Content *string
	MusicUrl url.URL
	ImageUrl *url.URL
}
