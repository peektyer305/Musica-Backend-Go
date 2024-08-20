package post

import (
	valueobject "Musica-Backend/internal/domain/value_object"
)

type IPostRepository interface {
	FindById(id valueobject.PostId) (Post, error)
	FindAll() ([]Post, error)
	Create(post Post) error
	DeleteById(id valueobject.PostId) error
}