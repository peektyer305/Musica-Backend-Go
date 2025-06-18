package auth0

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"sync"

	"github.com/golang-jwt/jwt"
)

type Auth0JWTService struct {
	domain     string
	clientID   string
	jwks       *JWKS
	keyCache   map[string]*rsa.PublicKey
	cacheMutex sync.RWMutex
}

func NewAuth0JWTService(domain, clientID string, jwks *JWKS) *Auth0JWTService {
	return &Auth0JWTService{
		domain:   domain,
		clientID: clientID,
		jwks:     jwks,
		keyCache: make(map[string]*rsa.PublicKey),
	}
}

func (s *Auth0JWTService) Validate(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, s.keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}
	if !strings.HasPrefix(claims["iss"].(string), "https://"+s.domain+"/") {
		return nil, jwt.NewValidationError("invalid issuer", jwt.ValidationErrorIssuer)
	}
	if claims["aud"] != s.clientID {
		return nil, jwt.NewValidationError("invalid audience", jwt.ValidationErrorAudience)
	}
	return claims, nil
}

func (s *Auth0JWTService) keyFunc(token *jwt.Token) (interface{}, error) {
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, jwt.NewValidationError("missing kid in token header", jwt.ValidationErrorMalformed)
	}
	s.cacheMutex.RLock()
	if key, exists := s.keyCache[kid]; exists {
		s.cacheMutex.RUnlock()
		return key, nil
	}
	s.cacheMutex.RUnlock()
	//キャッシュミス→JWKSからキーを取得
	for _, jwk := range s.jwks.Keys {
		if jwk.KeyId == kid {
			pubKey, err := jwk.DecodePublicKey()
			if err != nil {
				return nil, err
			}
			s.cacheMutex.Lock()
			s.keyCache[kid] = pubKey
			s.cacheMutex.Unlock()
			return pubKey, nil
		}
	}
	return nil, fmt.Errorf("key with kid %s not found", kid)
}
