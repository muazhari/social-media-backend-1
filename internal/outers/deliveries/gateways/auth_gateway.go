package gateways

import (
	"context"
	"encoding/json"
	"gopkg.in/square/go-jose.v2"
	"io"
	"net/http"
	"social-media-backend-1/internal/outers/configs"
)

type AuthGateway struct {
	AuthConfig *configs.AuthConfig
}

func NewAuthGateway(authConfig *configs.AuthConfig) *AuthGateway {
	return &AuthGateway{
		AuthConfig: authConfig,
	}
}

func (authGateway *AuthGateway) GetJwks() (*jose.JSONWebKeySet, error) {
	response, err := http.Get(authGateway.AuthConfig.JwksUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jwks := &jose.JSONWebKeySet{}
	err = json.NewDecoder(response.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	return jwks, nil
}

func (authGateway *AuthGateway) GetJwksPrivateKey(ctx context.Context) (string, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		authGateway.AuthConfig.JwksPrivateKeyUrl,
		nil,
	)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	keyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(keyBytes), nil
}

func (authGateway *AuthGateway) GetJwksPublicKey(ctx context.Context) (string, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		authGateway.AuthConfig.JwksPublicKeyUrl,
		nil,
	)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	keyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(keyBytes), nil
}
