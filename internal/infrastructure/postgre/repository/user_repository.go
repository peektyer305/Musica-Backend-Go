package repository

import (
	domain "Musica-Backend/internal/domain/user"
	valueobject "Musica-Backend/internal/domain/value_object"
	model "Musica-Backend/internal/infrastructure/postgre/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (u *UserRepository) FindById(ctx context.Context, id valueobject.UserId) (domain.User, error) {
	conn := u.Db.WithContext(ctx)
	var user model.User
	err := conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

