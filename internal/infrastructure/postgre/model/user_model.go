package model

import (
	postdomain "Musica-Backend/internal/domain/post"
	userdomain "Musica-Backend/internal/domain/user"
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

type User struct {
    Id          valueobject.UserId `gorm:"type:uuid;primaryKey;schema:app"`
    Username    string             `gorm:"type:varchar(50);unique;not null"`
    UserIconUrl *string          `gorm:"type:varchar(255);"`
    UserInfo    string             `gorm:"type:text;"`
    PrivateInfo UserPrivate        `gorm:"foreignKey:UserId;references:Id"` // UserPrivate の UserId が User の Id を参照
    Posts       []Post             `gorm:"foreignKey:UserId;references:Id;schema:app"` // Specify schema here if needed
    CreatedAt   time.Time          `gorm:"type:timestamp;not null"`
    UpdatedAt   time.Time          `gorm:"type:timestamp;not null"`
}

func (user *User) ToDomain() userdomain.User {
    var posts []postdomain.Post
    for _, post := range user.Posts {
        posts = append(posts, post.ToDomain())
    }
    return userdomain.User{
        Id:          user.Id,
        Username:    user.Username,
        UserIconUrl: user.UserIconUrl,
        UserInfo:    user.UserInfo,
        Posts:       posts,
    }
}
