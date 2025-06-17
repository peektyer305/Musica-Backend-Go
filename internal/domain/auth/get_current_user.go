package auth

import (
	domain "Musica-Backend/internal/domain/auth"
)

type GetCurrentUserUseCase interface {
	Execute(in *GetCurrentUserInput) (*GetCurrentUserOutput, error)
}

type getCurrentUserInteractor struct {
	jwtService     domain.JWTService
	userRepository domain.IUserRepository
}
