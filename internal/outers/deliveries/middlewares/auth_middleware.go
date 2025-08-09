package middlewares

import (
	"context"
	"net/http"
	"social-media-backend-1/internal/inners/models/value_objects"
	"social-media-backend-1/internal/inners/use_cases"
	"strings"
)

var ClaimCtxKey = &contextKey{"claimContextKey"}

type contextKey struct {
	name string
}

type AuthMiddleware struct {
	AuthUseCase *use_cases.AuthUseCase
}

func NewAuthMiddleware(
	authUseCase *use_cases.AuthUseCase,
) *AuthMiddleware {
	return &AuthMiddleware{
		AuthUseCase: authUseCase,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := m.AuthUseCase.VerifyToken(r.Context(), tokenString)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimContext(ctx context.Context) *value_objects.Claims {
	raw, _ := ctx.Value(ClaimCtxKey).(*value_objects.Claims)
	return raw
}
