package auth

import (
	domain "Mussica-Backend/internal/domain/auth"
)

// Output DTO
type GetCurrentUserOutput struct {
	User *domain.User
}
