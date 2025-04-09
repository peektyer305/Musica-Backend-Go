package model

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"go/token"

	"golang.org/x/crypto/bcrypt"
)

// UserPrivate は、ユーザーのプライベート情報を表す構造体です。
//認証方針決まったらnot null制約などフィールドにつける。
type UserPrivate struct {
	Id valueobject.UserPrivateId `json:"id" gorm:"type:uuid;primaryKey"`
	UserId      valueobject.UserId        `json:"user_id" gorm:"type:uuid;not null"` // これが外部キー
	Password string `json:"password" gorm:"type:varchar(255);"`
	MailAddress string `json:"mail_address" gorm:"type:varchar(255);"`
	CreatedAt string `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt string `json:"deleted_at" gorm:"type:timestamp;not null"`
	token.Token `json:"token" gorm:"type:varchar(255);"`
}

func (user *UserPrivate) SetPassword(password string) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	return nil
}

func (user *UserPrivate) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
