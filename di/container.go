//go:build wireinject
// +build wireinject

package di

import (
	usecase "Musica-Backend/internal/application/post"
	domain "Musica-Backend/internal/domain/post"
	"Musica-Backend/internal/infrastructure/postgre"
	"Musica-Backend/internal/infrastructure/postgre/repository"

	handler "github.com/Musica-Backend/internal/presentation/rest"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(

	postgre.NewGormPostgres,

	NewPostRepository,
	
	NewPostUseCase,

	
)

func NewPostRepository(db *gorm.DB) domain.IPostRepository {
	return &repository.PostRepository{Db: db}
}

func NewPostUseCase(postRepository domain.IPostRepository) *usecase.PostUseCase {
	return &usecase.PostUseCase{PostRepository: postRepository}
}

func NewPostHandler(postUseCase *usecase.PostUseCase) *handler.PostHandler {
	return &handler.PostHandler{PostUseCase: postUseCase}
}

