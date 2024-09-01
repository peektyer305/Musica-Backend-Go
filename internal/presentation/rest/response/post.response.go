package response

import (
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"

	domain "Musica-Backend/internal/domain/post"
)

type PostResponse struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	MusicUrl  url.URL `json:"music_url"`
	ImageUrl url.URL `json:"image_url"`
	UserIconUrl url.URL `json:"user_icon_url"`
	UserName string `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DomainToResponse(post domain.Post) PostResponse {
	 var id = post.Id.GetUUID()
	 var userId = post.UserId.GetUUID()
	return PostResponse{
		Id:       id,
		UserId:   userId,
		Title:    post.Title,
		Content:  *post.Content,
		MusicUrl: post.MusicUrl,
		ImageUrl: *post.ImageUrl,
		UserIconUrl: *post.UserIconUrl,
		UserName: post.UserName,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}