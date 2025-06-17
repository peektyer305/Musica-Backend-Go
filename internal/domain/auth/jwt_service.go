package auth

type JWTService interface {
	//ValidateTokenを検証し、ペイロードを返す
	Validate(token string) (map[string]interface{}, error)
}
