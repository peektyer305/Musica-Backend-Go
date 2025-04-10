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
	
	// まずユーザーを取得
	var user model.User
	if err := conn.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}

	// 次に、ユーザーIDに一致する全Postを取得
	var posts []model.Post
	if err := conn.Where("user_id = ?", id).Find(&posts).Error; err != nil {
		return domain.User{}, err
	}

	// 取得した posts をユーザーのフィールドに代入
	user.Posts = posts  // user.Posts が []model.Post 型である前提
	// 最後に、ユーザーをドメインモデルに変換して返す
	return user.ToDomain(), nil
}
