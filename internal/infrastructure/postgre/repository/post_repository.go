package repository

import (
	"Musica-Backend/internal/domain"

	"gorm.io/gorm"
)

type PostRepository struct {
	Db *gorm.DB
}

func (p *PostRepository) findAll() ([]domain.Post, error) {
	var posts []domain.Post
	//とりあえず50件まで取得
	err := p.Db.Limit(50).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		posts = append(posts, post.ToDomain())
	}

	return posts, nil
}