package utils

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const claimsContextKey = contextKey("jwtClaims")

// AddClaimsToContext adds JWT claims to the context
func AddClaimsToContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, claimsContextKey, claims)
}

// GetClaimsFromContext retrieves JWT claims from the context
func GetClaimsFromContext(ctx context.Context) (jwt.MapClaims, error) {
	claims, ok := ctx.Value(claimsContextKey).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no claims found in context")
	}
	return claims, nil
}
