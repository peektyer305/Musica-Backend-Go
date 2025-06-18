package auth

import (
	"Musica-Backend/internal/domain/auth"
)

// Output DTO
type GetCurrentUserOutput struct {
	User *auth.User
}
