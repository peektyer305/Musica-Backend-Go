package auth

import (
	"errors"
)

type GetCurrentUserUseCase interface {
	Execute(in *GetCurrentUserInput) (*GetCurrentUserOutput, error)
}

type getCurrentUserInteractor struct {
	jwtService     JWTService
	userRepository IAuthRepository
}

func NewGetCurrentUserUsecase(
	jwtSvc JWTService,
	userRepo IAuthRepository,
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
	// Auth0からemailを取得
	if claims == nil {
		return nil, errors.New("invalid token claims")
	}
	// claimsからemailを取得
	if claims["email"] == nil {
		return nil, errors.New("email not found in token claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid token email")
	}
	// ユーザリポジトリからユーザを取得
	if email == "" {
		return nil, errors.New("email is empty")
	}
	// オプション: ユーザリポジトリから詳細情報を取得
	user, err := i.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return &GetCurrentUserOutput{User: user}, nil
}
