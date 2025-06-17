package auth0

type JWK struct {
	KeyId string `json:"kid"`
	Alg   string `json:"alg"`
}
type JWKS struct {
	Keys []JWK `json:"keys"`
}
