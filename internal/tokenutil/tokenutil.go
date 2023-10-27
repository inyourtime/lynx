package tokenutil

// import (
// 	"lynx/model"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// )

// func CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
// 	claims := AccessClaims{
// 		ID:    user.UserID,
// 		Email: user.Email,
// 		Name:  user.Firstname + " " + user.Lastname,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			Issuer:    "access_token",
// 			Subject:   "users_access_token",
// 			ID:        uuid.NewString(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	t, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return t, err
// }

// func CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
// 	claims := RefreshClaims{
// 		ID: user.UserID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			Issuer:    "refresh_token",
// 			Subject:   "users_refresh_token",
// 			ID:        uuid.NewString(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	rt, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return rt, err
// }

// func ExtractIDFromToken(requestToken string, secret string) (string, error) {
// 	claim := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(requestToken, &claim, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return claim["user_id"].(string), nil
// }
