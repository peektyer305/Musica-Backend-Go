package auth0

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
)

type JWK struct {
	KeyId string `json:"kid"`
	Alg   string `json:"alg"`
	Kty   string `json:"kty"`
	N     string `json:"n"`
	E     string `json:"e"`
}
type JWKS struct {
	Keys []JWK `json:"keys"`
}

func FetchJWKS(domain string) (*JWKS, error) {
	url := "https://musica-auth.jp.auth0.com/.well-known/jwks.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch JWKS: received status code %d", resp.StatusCode)
	}
	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %w", err)
	}
	return &jwks, nil
}

func (j *JWK) DecodePublicKey() (*rsa.PublicKey, error) {
	if j.Kty != "RSA" {
		return nil, errors.New("unsupported key type: " + j.Kty)
	}

	// Base64URL からバイナリに戻す
	nBytes, err := base64.RawURLEncoding.DecodeString(j.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(j.E)
	if err != nil {
		return nil, err
	}

	// eBytes は小さいので big.Int→int へ変換
	eInt := 0
	for _, b := range eBytes {
		eInt = eInt<<8 + int(b)
	}

	pub := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: eInt,
	}

	// オプションで ASN.1 DER にエンコード／デコードして検証する場合
	if _, err := x509.MarshalPKIXPublicKey(pub); err != nil {
		return nil, err
	}

	return pub, nil
}
