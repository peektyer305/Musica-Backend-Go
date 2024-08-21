package model

import (
	"Musica-Backend/internal/domain/post"
	valueobject "Musica-Backend/internal/domain/value_object"
	url "net/url"
	"time"
)

type Post struct {
    Id      valueobject.PostId `gorm:"type:uuid;primaryKey"`
    UserId  valueobject.UserId `gorm:"type:uuid;not null"`
    Title   string             `gorm:"type:varchar(50);not null"`
    Content *string             `gorm:"type:text;"`
    User    User               `gorm:"foreignKey:UserId"`
    MusicUrl string            `gorm:"type:varchar(255); not null"`
    ImageUrl *string            `gorm:"type:varchar(255);"`
    CreatedAt time.Time           `gorm:"type:timestamp;not null"`
    UpdatedAt time.Time           `gorm:"type:timestamp;not null"`
}

func (p *Post) ToDomain() post.Post {
    var musicUrl *url.URL
    musicUrl, err := url.Parse(p.MusicUrl)
    if err != nil {
        panic(err)
    }
    var imageUrl *url.URL
    if p.ImageUrl != nil {
        imageUrl, err = url.Parse(*p.ImageUrl)
        if err != nil {
            panic(err)
        }

    }

    var userIconUrl *url.URL
    if p.User.UserIconUrl != nil {
        userIconUrl, err = url.Parse((*p.User.UserIconUrl))
        if err != nil {
            panic(err)
        }
    }

    return post.Post{
        Id:      p.Id,
        UserId:  p.UserId,
        Title:   p.Title,
        Content: p.Content,
        MusicUrl: *musicUrl,
        ImageUrl: imageUrl,
        UserIconUrl: userIconUrl,
        UserName: p.User.Username,
        CreatedAt: p.CreatedAt,
        UpdatedAt: p.UpdatedAt,
    }
}