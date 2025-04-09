package repository

import (
	domain "Musica-Backend/internal/domain/post"
	valueobject "Musica-Backend/internal/domain/value_object"
	"Musica-Backend/internal/infrastructure/postgre/model"
	"context"

	"gorm.io/gorm"
)

type PostRepository struct {
	Db *gorm.DB
}

func (p *PostRepository) FindAll(ctx context.Context) ([]domain.Post, error) {
	conn:= p.Db.WithContext(ctx)
	var posts []model.Post
	//とりあえず50件まで取得
	err := conn.Limit(50).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	var domainPosts []domain.Post
	for _, post := range posts {
		domainPosts = append(domainPosts, post.ToDomain())
	}

	return domainPosts, nil
}

func (p *PostRepository) FindById(ctx context.Context, id valueobject.PostId) (domain.Post, error) {
	conn:= p.Db.WithContext(ctx)
	var post model.Post
	err := conn.Where("id = ?", id).First(&post).Error
	if err != nil {
		return domain.Post{}, err
	}
	return post.ToDomain(), nil
}
//この下二つは仮実装
func (p *PostRepository) Create(ctx context.Context,post domain.Post) error {
	conn:= p.Db.WithContext(ctx)
	postModel := model.Post{
		Id:       post.Id,
		UserId:   post.UserId,
		Title:    post.Title,
		Content:  post.Content,
		MusicUrl: post.MusicUrl,
	}
	err := conn.Create(&postModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepository) DeleteById(ctx context.Context,id valueobject.PostId) error {
	conn:= p.Db.WithContext(ctx)
	err := conn.Where("id = ?", id).Delete(&model.Post{}).Error
	if err != nil {
		return err
	}
	return nil
}
