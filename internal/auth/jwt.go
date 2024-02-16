package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

var ErrNoAuthHeader = errors.New("no authorization header included")
var ErrMalformedAuthHeader = errors.New("malformed authorization header")
var ErrNoAccessToken = errors.New("no access token included")

func MakeJWT(
	userID uuid.UUID,
	tokenSecret string,
	expiresIn time.Duration,
	tokenType TokenType,
) (string, error) {
	signingKey := []byte(tokenSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(tokenType),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   userID.String(),
	})
	return token.SignedString(signingKey)
}

func ValidateJWT(tokenString, tokenSecret string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		return "", err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", err
	}
	if issuer != string(TokenTypeAccess) {
		return "", errors.New("invalid issuer")
	}

	return userIDString, nil
}

func GetJWTToken(r *http.Request) (string, error) {
	token, err := GetBearerToken(r)
	if err == nil {
		return token, nil
	}

	token, err = GetCookieToken(r)
	if err == nil {
		return token, nil
	}

	return "", ErrNoAccessToken
}

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", ErrMalformedAuthHeader
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func SetCookieToken(w http.ResponseWriter, token string, expiresIn time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(expiresIn),
		Secure:   true,
		HttpOnly: true,
		// SameSite: http.SameSiteStrictMode,
	})
}

func GetCookieToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
