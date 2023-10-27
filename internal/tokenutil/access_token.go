package tokenutil

import (
	"lynx/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessClaims struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func NewAccessClaims(u *model.User, expiry int) (ac *AccessClaims) {
	now := time.Now()
	ac = &AccessClaims{
		ID:    u.UserID,
		Email: u.Email,
		Name:  u.Firstname + " " + u.Lastname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expiry) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
		},
	}
	return
}

func (ac *AccessClaims) IsExpired() bool {
	now := time.Now()
	exp := ac.ExpiresAt.Time
	return now.After(exp)
}

func (ac *AccessClaims) JwtString(secret string) (accessTokenString string, err error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessTokenString, err = accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return
}
