package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
)

func Sign(userId uint, expirationTime time.Duration) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(expirationTime).Unix()

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Verify(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid userId type in token")
	}

	return uint(userIdFloat), nil
}
