package post

import (
	domain "Musica-Backend/internal/domain/post"
	"context"
)

type PostUseCase struct {
	PostRepository domain.IPostRepository
}

func (p *PostUseCase) FindAll(ctx context.Context) ([]domain.Post, error) {
	return p.PostRepository.FindAll(ctx)
}
