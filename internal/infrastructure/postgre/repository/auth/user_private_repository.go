package auth

import (
	"Musica-Backend/internal/domain/auth"
	authmodel "Musica-Backend/internal/infrastructure/postgre/model/auth"
	"context"

	"gorm.io/gorm"
)

type UserPrivateRepository struct {
	Db *gorm.DB
}

func (r *UserPrivateRepository) FindMe(ctx context.Context, email string) (auth.UserPrivate, error) {
	conn := r.Db.WithContext(ctx)
	// ユーザープライベートテーブルからユーザーを取得
	var userPrivateInfo authmodel.UserPrivate
	err := conn.Where("email = ?", email).First(&userPrivateInfo).Error
	if err != nil {
		return auth.UserPrivate{}, err
	}
	// ユーザープライベートのUserIdを使って、ユーザーテーブルからユーザー情報を取得
	err = conn.Where("user_id = ?", userPrivateInfo.UserId).First(&userPrivateInfo).Error
	if err != nil {
		return auth.UserPrivate{}, err
	}
	return userPrivateInfo.ToDomain(), nil
}
