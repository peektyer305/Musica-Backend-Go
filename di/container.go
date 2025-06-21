//go:build wireinject
// +build wireinject

package di

import (
	postusecase "Musica-Backend/internal/application/post"
	postdomain "Musica-Backend/internal/domain/post"
	"Musica-Backend/internal/infrastructure/postgre"
	"Musica-Backend/internal/infrastructure/postgre/repository"

	authusecase "Musica-Backend/internal/application/auth"
	userusecase "Musica-Backend/internal/application/user"
	authdomain "Musica-Backend/internal/domain/auth"
	userdomain "Musica-Backend/internal/domain/user"
	authrepository "Musica-Backend/internal/infrastructure/postgre/repository/auth"
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

	NewUserPrivateRepository,

	NewUserPrivateUseCase,

	NewUserPrivateHandler,
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

func NewUserPrivateRepository(db *gorm.DB) authdomain.IUserPrivateRepository {
	return &authrepository.UserPrivateRepository{Db: db}
}

func NewUserPrivateUseCase(userPrivateRepository authdomain.IUserPrivateRepository) *authusecase.UserPrivateUsecase {
	return &authusecase.UserPrivateUsecase{UserPrivateRepository: userPrivateRepository}
}

func NewUserPrivateHandler(userPrivateUseCase *authusecase.UserPrivateUsecase) *handlers.UserPrivateHandler {
	return &handlers.UserPrivateHandler{UserPrivateUsecase: userPrivateUseCase}
}
func InitializeUserPrivateHandler() *handlers.UserPrivateHandler {
	wire.Build(providerSet)
	return &handlers.UserPrivateHandler{}
}
