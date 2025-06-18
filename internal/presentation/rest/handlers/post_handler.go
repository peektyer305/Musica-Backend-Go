package handlers

import (
	usecase "Musica-Backend/internal/application/post"
	"Musica-Backend/internal/presentation/rest/response"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostUseCase *usecase.PostUseCase
}

func (p *PostHandler) FindAll(ctx echo.Context) ([]response.PostResponse, error) {
	posts, err := p.PostUseCase.FindAll(ctx.Request().Context())
	if err != nil {
		return nil, err
	}
	var responses []response.PostResponse
	for _, post := range posts {
		responses = append(responses, response.DomainToResponse(post))
	}

	return responses, nil
}
