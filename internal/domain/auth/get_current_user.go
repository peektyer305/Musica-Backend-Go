package auth

import (
	"errors"

	domain "Musica-Backend/internal/domain/auth"
)

type GetCurrentUserUseCase interface {
	Execute(in *GetCurrentUserInput) (*GetCurrentUserOutput, error)
}

type getCurrentUserInteractor struct {
	jwtService     domain.JWTService
	userRepository domain.IUserRepository
}

func NewGetCurrentUserUsecase(
	jwtSvc domain.JWTService,
	userRepo domain.IUserRepository,
) GetCurrentUserUseCase {
	return &getCurrentUserInteractor{
		jwtService:     jwtSvc,
		userRepository: userRepo,
	}
}

func (i *getCurrentUserInteractor) Execute(in *GetCurrentUserInput) (*GetCurrentUserOutput, error) {
	claims, err := i.jwtService.Validate(in.Token)
	if err != nil {
		return nil, err
	}
	// Auth0 の 'sub' をユーザーIDとみなす
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token subject")
	}
	// オプション: ユーザリポジトリから詳細情報を取得
	user, err := i.userRepository.FindByID(sub)
	if err != nil {
		return nil, err
	}
	return &GetCurrentUserOutput{User: user}, nil
}
