package response

import (
	domain "Musica-Backend/internal/domain/user"
	util "Musica-Backend/utils"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResponse struct {
	Id           uuid.UUID          `json:"id"`
	Username     string             `json:"username"`
	UserIconUrl  *string            `json:"userIconUrl"`
	UserInfo     string             `json:"userInfo"`
	UserClientId string             `json:"userClientId"` // ユーザークライアントID
	Posts        []UserPostResponse `json:"posts"`
}

type UserPostResponse struct {
	Id        uuid.UUID         `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	MusicUrl  map[string]string `json:"music"`
	ImageUrl  *string           `json:"imageUrl"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func DomainToResponseUser(user domain.User) UserResponse {
	Id := user.Id.GetUUID()
	UserIconUrl := user.UserIconUrl
	posts := make([]UserPostResponse, len(user.Posts))
	for i, post := range user.Posts {
		music, err := util.FetchMetadata(post.MusicUrl)
		if err != nil {
			panic(err)
		}
		posts[i] = UserPostResponse{
			Id:        post.Id.GetUUID(),
			Title:     post.Title,
			Content:   *post.Content,
			MusicUrl:  music,
			ImageUrl:  post.ImageUrl,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
	}
	return UserResponse{
		Id:           Id,
		Username:     user.Username,
		UserIconUrl:  UserIconUrl,
		UserInfo:     user.UserInfo,
		UserClientId: user.UserClientId, // ユーザークライアントID
		Posts:        posts,
	}
}
