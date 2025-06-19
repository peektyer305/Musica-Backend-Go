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

	authusecase "Musica-Backend/internal/application/auth"
	authdomain "Musica-Backend/internal/domain/auth"

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

	NewAuthRepository,

	NewAuthUseCase,

	NewAuthHandler,
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

// Auth関連のDI設定
func NewAuthRepository(db *gorm.DB) authdomain.IAuthRepository {
	return &repository.AuthRepository{Db: db}
}

func NewAuthUseCase(
	jwtService authdomain.JWTService,
	authRepository authdomain.IAuthRepository,
) *authusecase.AuthUseCase {
	return &authusecase.AuthUseCase{
		JWTService:     jwtService,
		AuthRepository: authRepository,
	}
}

func NewAuthHandler(authUseCase *authusecase.AuthUseCase) *handlers.AuthHandler {
	return &handlers.AuthHandler{AuthUseCase: authUseCase}
}

func InitializeAuthHandler() *handlers.AuthHandler {
	wire.Build(providerSet)
	return &handlers.AuthHandler{}
}
