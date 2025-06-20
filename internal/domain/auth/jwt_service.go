package auth

//トークン検証のためのドメインサービス

type JWTService interface {
	// Validateはトークンを検証し、クレームを返す。
	Validate(token string) (map[string]interface{}, error)
}
