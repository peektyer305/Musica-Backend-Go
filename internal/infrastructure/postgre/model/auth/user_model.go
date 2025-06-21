package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
)

type User struct {
	Id           valueobject.UserId `gorm:"type:uuid;primaryKey;schema:app"`
	UserName     string             `gorm:"type:varchar(50);not null"`
	UserInfo     string             `gorm:"type:text;"`
	UserIconUrl  string             `gorm:"type:varchar(255);"`
	UserClientId string             `gorm:"type:varchar(255);not null;unique"`
}
