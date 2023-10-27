package tokenutil

import (
	"lynx/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshClaims struct {
	ID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshClaims(u *model.User, expiry int) (rc *RefreshClaims) {
	now := time.Now()
	rc = &RefreshClaims{
		ID: u.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expiry) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "refresh_token",
			Subject:   "users_refresh_token",
			ID:        uuid.NewString(),
		},
	}
	return
}

func (rc *RefreshClaims) IsExpired() bool {
	now := time.Now()
	exp := rc.ExpiresAt.Time
	return now.After(exp)
}

func (rc *RefreshClaims) JwtString(secret string) (refreshTokenString string, err error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshTokenString, err = refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return
}
