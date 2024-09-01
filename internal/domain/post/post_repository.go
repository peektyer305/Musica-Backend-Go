package post

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"context"
)

type IPostRepository interface {
	FindById(ctx context.Context, id valueobject.PostId) (Post, error)
	FindAll(ctx context.Context) ([]Post, error)
	Create(ctx context.Context,post Post) error
	DeleteById(ctx context.Context,id valueobject.PostId) error
}