//go:build wireinject
// +build wireinject

package di

import (
	postusecase "Musica-Backend/internal/application/post"
	postdomain "Musica-Backend/internal/domain/post"
	"Musica-Backend/internal/infrastructure/postgre"
	"Musica-Backend/internal/infrastructure/postgre/repository"

	userusecase "Musica-Backend/internal/application/user"
	userdomain "Musica-Backend/internal/domain/user"
	handlers "Musica-Backend/internal/presentation/rest/handlers"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(

	postgre.NewGormPostgres,

	NewPostRepository,

	NewPostUseCase,

	NewPostHandler,

	NewUserRepository,

	NewUserUseCase,

	NewUserHandler,
)

func NewPostRepository(db *gorm.DB) postdomain.IPostRepository {
	return &repository.PostRepository{Db: db}
}

func NewPostUseCase(postRepository postdomain.IPostRepository) *postusecase.PostUseCase {
	return &postusecase.PostUseCase{PostRepository: postRepository}
}

func NewPostHandler(postUseCase *postusecase.PostUseCase) *handlers.PostHandler {
	return &handlers.PostHandler{PostUseCase: postUseCase}
}

func InitializePostHandler() *handlers.PostHandler {
	wire.Build(providerSet)
	return &handlers.PostHandler{}
}

func NewUserRepository(db *gorm.DB) userdomain.IUserRepository {
	return &repository.UserRepository{Db: db}
}
func NewUserUseCase(userRepository userdomain.IUserRepository) *userusecase.UserUseCase {
	return &userusecase.UserUseCase{UserRepository: userRepository}
}
func NewUserHandler(userUseCase *userusecase.UserUseCase) *handlers.UserHandler {
	return &handlers.UserHandler{UserUseCase: userUseCase}
}
func InitializeUserHandler() *handlers.UserHandler {
	wire.Build(providerSet)
	return &handlers.UserHandler{}
}
