package use_cases

import (
	"context"
	"crypto/rsa"
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

func (uc *AuthUseCase) Register(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	account.ID = &id
	account.Scopes = []string{"user"}
	account.TotalPostLike = &[]float64{0}[0]
	account.TotalChatMessage = &[]float64{0}[0]

	createdAccount, err := uc.AccountRepository.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	return createdAccount, nil
}

func (uc *AuthUseCase) createToken(ctx context.Context, claims *value_objects.Claims) (string, error) {
	privateKey, err := uc.AuthGateway.GetJwksPrivateKey(ctx)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode([]byte(privateKey))
	parseResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parseResult.(*rsa.PrivateKey)
	opts := &jose.SignerOptions{}
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: key},
		opts.WithType("JWT").WithHeader("kid", "social-media-backend-key"),
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

func (uc *AuthUseCase) Login(ctx context.Context, email string, password string) (*value_objects.Session, error) {
	foundAccount, err := uc.AccountRepository.GetAccountByEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()

	accessTokenClaims := &value_objects.Claims{
		Subject:  foundAccount.ID.String(),
		IssuedAt: jwt.NewNumericDate(timeNow),
		Expiry:   jwt.NewNumericDate(timeNow.Add(15 * time.Minute)),
		Issuer:   "social-media-backend-1",
		Scope:    strings.Join(foundAccount.Scopes, " "),
		Audience: &jwt.Audience{"social-media-backend"},
	}
	accessToken, err := uc.createToken(ctx, accessTokenClaims)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := &value_objects.Claims{
		Subject:  foundAccount.ID.String(),
		IssuedAt: jwt.NewNumericDate(timeNow),
		Expiry:   jwt.NewNumericDate(timeNow.Add(24 * time.Hour)),
		Issuer:   "social-media-backend-1",
		Scope:    strings.Join(foundAccount.Scopes, " "),
		Audience: &jwt.Audience{"social-media-backend"},
	}
	refreshToken, err := uc.createToken(ctx, refreshTokenClaims)
	if err != nil {
		return nil, err
	}

	session := &value_objects.Session{
		Account:      foundAccount,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return session, nil
}
