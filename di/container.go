//go:build wireinject
// +build wireinject

package di

import (
	postusecase "Musica-Backend/internal/application/post"
	postdomain "Musica-Backend/internal/domain/post"
	"Musica-Backend/internal/infrastructure/postgre"
	"Musica-Backend/internal/infrastructure/postgre/repository"

	handler "Musica-Backend/internal/presentation/rest"

	userusecase "Musica-Backend/internal/application/user"
	userdomain "Musica-Backend/internal/domain/user"

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

func NewPostHandler(postUseCase *postusecase.PostUseCase) *handler.PostHandler {
	return &handler.PostHandler{PostUseCase: postUseCase}
}

func InitializePostHandler() *handler.PostHandler {
	wire.Build(providerSet)
	return &handler.PostHandler{}
}

func NewUserRepository(db *gorm.DB) userdomain.IUserRepository {
	return &repository.UserRepository{Db: db}
}
func NewUserUseCase(userRepository userdomain.IUserRepository) *userusecase.UserUseCase {
	return &userusecase.UserUseCase{UserRepository: userRepository}
}
func NewUserHandler(userUseCase *userusecase.UserUseCase) *handler.UserHandler {
	return &handler.UserHandler{UserUseCase: userUseCase}
}
func InitializeUserHandler() *handler.UserHandler {
	wire.Build(providerSet)
	return &handler.UserHandler{}
}