package auth

import (
	domain "Musica-Backend/internal/domain/auth"
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

type UserPrivate struct {
	Id        valueobject.UserPrivateId `gorm:"type:uuid;primaryKey;schema:app"`
	UserId    valueobject.UserId        `gorm:"type:uuid;not null;unique"`
	Email     string                    `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time                 `gorm:"type:timestamp;not null"`
	User      User                      `gorm:"foreignKey:UserId;references:Id;schema:app"` // Specify schema here if needed
}

func (u *UserPrivate) ToDomain() domain.UserPrivate {
	return domain.UserPrivate{
		Id:           u.Id,
		UserId:       u.User.Id,
		UserName:     u.User.UserName,
		UserInfo:     u.User.UserInfo,
		UserIconUrl:  u.User.UserIconUrl,
		UserClientId: u.User.UserClientId,
		Email:        u.Email,
		CreatedAt:    u.CreatedAt,
	}
}
