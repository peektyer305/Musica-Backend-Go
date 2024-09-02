package response

import (
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"

	domain "Musica-Backend/internal/domain/post"

	util "Musica-Backend/utils"
)

type PostResponse struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"userId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	MusicUrl  map[string]string `json:"music"`
	ImageUrl url.URL `json:"imageUrl"`
	UserIconUrl url.URL `json:"userIconUrl"`
	UserName string `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func DomainToResponse(post domain.Post) PostResponse {
	 var id = post.Id.GetUUID()
	 var userId = post.UserId.GetUUID()
	 music,err := util.FetchMetadata(post.MusicUrl.String())
	 if err != nil {
		 panic(err)
	 }
	return PostResponse{
		Id:       id,
		UserId:   userId,
		Title:    post.Title,
		Content:  *post.Content,
		MusicUrl: music,
		ImageUrl: *post.ImageUrl,
		UserIconUrl: *post.UserIconUrl,
		UserName: post.UserName,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}