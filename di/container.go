//go:build wireinject
// +build wireinject

package di

import (
	"Musica-Backend/internal/domain"
	"Musica-Backend/internal/infrastructure/postgre"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(

	postgre.NewGormPostgres,

	NewTodoRepository,
	
)

func NewPostRepository(db *gorm.DB) domain.IPostRepository {
	return &postgre.PostRepository{Db: db}
}