package auth

import (
	auth "Musica-Backend/internal/domain/auth"
	"context"
	"errors"

	"gorm.io/gorm"
)

type AuthUserRepository struct {
	Db *gorm.DB
}

func FindByEmail(ctx context.Context, db *gorm.DB, email string) (*auth.User, error) {
	var user auth.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // ユーザーが見つからない場合は nil を返す
		}
		return nil, err // その他のエラーはそのまま返す
	}
	return &user, nil
}
