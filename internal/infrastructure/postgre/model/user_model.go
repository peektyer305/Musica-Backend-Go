package model

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
    Id          valueobject.UserId `gorm:"type:uuid;primaryKey;schema:app"`
    Username    string             `gorm:"type:varchar(50);unique;not null"`
    UserIconUrl *string            `gorm:"type:varchar(255);"`
    Email       string             `gorm:"type:varchar(50);unique;not null"`
    Password    string             `gorm:"type:varchar(255);not null"`
    Posts       []Post             `gorm:"foreignKey:UserId;references:Id;schema:app"` // Specify schema here if needed
    CreatedAt   time.Time          `gorm:"type:timestamp;not null"`
    UpdatedAt   time.Time          `gorm:"type:timestamp;not null"`
}


func (user *User) SetPassword(password string) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
