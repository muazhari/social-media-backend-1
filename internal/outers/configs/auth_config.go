package configs

import (
	"os"
)

type AuthConfig struct {
	JwksUrl           string
	JwksPrivateKeyUrl string
	JwksPublicKeyUrl  string
}

func NewAuthConfig() *AuthConfig {
	jwksUrl := os.Getenv("AUTH_JWKS_URL")
	if jwksUrl == "" {
		panic("AUTH_JWKS_URL environment variable is not set")
	}

	jwksPrivateKeyUrl := os.Getenv("AUTH_JWKS_PRIVATE_KEY_URL")
	if jwksPrivateKeyUrl == "" {
		panic("AUTH_JWKS_PRIVATE_KEY_URL environment variable is not set")
	}

	jwksPublicKeyUrl := os.Getenv("AUTH_JWKS_PUBLIC_KEY_URL")
	if jwksPublicKeyUrl == "" {
		panic("AUTH_JWKS_PUBLIC_KEY_URL environment variable is not set")
	}

	authConfig := &AuthConfig{
		JwksUrl:           jwksUrl,
		JwksPrivateKeyUrl: jwksPrivateKeyUrl,
		JwksPublicKeyUrl:  jwksPublicKeyUrl,
	}

	return authConfig
}
