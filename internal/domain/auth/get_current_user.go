package auth

import (
	"errors"

	"Musica-Backend/internal/domain/auth"
)

type GetCurrentUserUseCase interface {
	Execute(in *GetCurrentUserInput) (*GetCurrentUserOutput, error)
}

type getCurrentUserInteractor struct {
	jwtService     auth.JWTService
	userRepository auth.IUserRepository
}

func NewGetCurrentUserUsecase(
	jwtSvc auth.JWTService,
	userRepo auth.IUserRepository,
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
