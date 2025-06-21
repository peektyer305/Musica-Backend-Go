package model

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"go/token"
)

// UserPrivate は、ユーザーのプライベート情報を表す構造体です。
// 認証方針決まったらnot null制約などフィールドにつける。
type UserPrivate struct {
	Id          valueobject.UserPrivateId `json:"id" gorm:"type:uuid;primaryKey"`
	UserId      valueobject.UserId        `json:"user_id" gorm:"type:uuid;not null"` // これが外部キー
	Password    string                    `json:"password" gorm:"type:varchar(255);"`
	MailAddress string                    `json:"mail_address" gorm:"type:varchar(255);"`
	CreatedAt   string                    `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt   string                    `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt   string                    `json:"deleted_at" gorm:"type:timestamp;not null"`
	token.Token `json:"token" gorm:"type:varchar(255);"`
}
