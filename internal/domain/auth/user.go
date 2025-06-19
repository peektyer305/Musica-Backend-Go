package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"time"
)

type User struct {
	ID        valueobject.UserPrivateId `json:"id"`
	UserId    valueobject.UserId        `json:"user_id" gorm:"type:uuid;not null"` // 外部キー
	Email     string                    `json:"email"`
	Name      string                    `json:"name"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}
