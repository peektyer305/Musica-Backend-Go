package rest

import (
	usecase "Musica-Backend/internal/application/post"
	domain "Musica-Backend/internal/domain/post"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostUseCase *usecase.PostUseCase
}

func (p *PostHandler) FindAll(ctx echo.Context) ([]domain.Post, error) {
	posts, err := p.PostUseCase.FindAll(ctx.Request().Context())
	if err != nil {
		return nil, err
	}
	return posts, nil
}