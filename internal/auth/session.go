package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var ErrNoSessionOrExpired = errors.New("no session exists or session has expired")

// This can be replaced with a database table
var sessions = map[string]session{}

type session struct {
	userId uuid.UUID
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func SetSessionToken(w http.ResponseWriter, userId uuid.UUID, expiresIn time.Duration) {
	expiresAt := time.Now().Add(expiresIn)

	sessionToken := uuid.NewString()
	sessions[sessionToken] = session{userId, expiresAt}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}

func GetSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func GetSessionUserId(sessionToken string) (uuid.UUID, error) {
	s, ok := sessions[sessionToken]
	if !ok || s.isExpired() {
		return uuid.Nil, ErrNoSessionOrExpired
	}
	return s.userId, nil
}

func RefreshSessionToken(w http.ResponseWriter, r *http.Request, expiresIn time.Duration) error {
	sessionToken, err := GetSessionToken(r)
	if err != nil {
		return err
	}

	delete(sessions, sessionToken)

	userId, err := GetSessionUserId(sessionToken)
	if err != nil {
		return err
	}

	SetSessionToken(w, userId, expiresIn)
	return nil
}

func ClearSessionToken(w http.ResponseWriter, r *http.Request) error {
	sessionToken, err := GetSessionToken(r)
	if err != nil {
		return err
	}

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Second),
	})
	return nil
}
