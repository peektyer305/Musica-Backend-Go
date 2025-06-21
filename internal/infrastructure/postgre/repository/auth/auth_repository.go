package auth

import (
	"Musica-Backend/internal/domain/auth"
	authmodel "Musica-Backend/internal/infrastructure/postgre/model/auth"
	"context"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func FindMe(ctx context.Context, db *gorm.DB, email string) (auth.User, error) {
	conn := db.WithContext(ctx)
	// ユーザープライベートテーブルからユーザーを取得
	var userPrivateInfo authmodel.UserPrivate
	err := conn.Where("email = ?", email).First(&userPrivateInfo).Error
	if err != nil {
		return auth.User{}, err
	}
	// ユーザープライベートのUserIdを使って、ユーザーテーブルからユーザー情報を取得
	err = conn.Where("user_id = ?", userPrivateInfo.UserId).First(&userPrivateInfo).Error
	if err != nil {
		return auth.User{}, err
	}
	return userPrivateInfo.ToDomain(), nil
}
