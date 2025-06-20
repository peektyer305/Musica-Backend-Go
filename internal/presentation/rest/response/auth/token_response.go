package auth

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewTokenResponse(accessToken, idToken, tokenType string, expiresIn int) TokenResponse {
	return TokenResponse{
		AccessToken: accessToken,
		IDToken:     idToken,
		TokenType:   tokenType,
		ExpiresIn:   expiresIn,
	}
}
