package use_cases

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"social-media-backend-1/internal/inners/models/entities"
	"social-media-backend-1/internal/inners/models/value_objects"
	"social-media-backend-1/internal/outers/deliveries/gateways"
	"social-media-backend-1/internal/outers/repositories"
	"strings"
	"time"
)

type AuthUseCase struct {
	AccountRepository *repositories.AccountRepository
	AuthGateway       *gateways.AuthGateway
}

func NewAuthUseCase(
	accountRepository *repositories.AccountRepository,
	authGateway *gateways.AuthGateway,
) *AuthUseCase {
	return &AuthUseCase{
		AccountRepository: accountRepository,
		AuthGateway:       authGateway,
	}
}

func (uc *AuthUseCase) Register(account *entities.Account) (*entities.Account, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	account.ID = &id
	account.Scopes = []string{"user"}

	createdAccount, err := uc.AccountRepository.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	return createdAccount, nil
}

func (uc *AuthUseCase) createToken(claims *value_objects.Claims) (string, error) {
	privateKey, err := uc.AuthGateway.GetJwksPrivateKey()
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode([]byte(privateKey))
	parseResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parseResult.(ed25519.PrivateKey)
	opts := &jose.SignerOptions{}
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.EdDSA, Key: key},
		opts.WithType("JWT"),
	)
	if err != nil {
		return "", err
	}

	token, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *AuthUseCase) Login(email string, password string) (*value_objects.Session, error) {
	account, err := uc.AccountRepository.GetAccountByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()

	accessTokenClaims := &value_objects.Claims{
		Subject:  account.ID.String(),
		IssuedAt: jwt.NewNumericDate(timeNow),
		Expiry:   jwt.NewNumericDate(timeNow.Add(15 * time.Minute)),
		Issuer:   "social-media-backend-1",
		Scopes:   strings.Join(account.Scopes, " "),
	}
	accessToken, err := uc.createToken(accessTokenClaims)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := &value_objects.Claims{
		Subject:  account.ID.String(),
		IssuedAt: jwt.NewNumericDate(timeNow),
		Expiry:   jwt.NewNumericDate(timeNow.Add(24 * time.Hour)),
		Issuer:   "social-media-backend-1",
		Scopes:   strings.Join(account.Scopes, " "),
	}
	refreshToken, err := uc.createToken(refreshTokenClaims)
	if err != nil {
		return nil, err
	}

	session := &value_objects.Session{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return session, nil
}
