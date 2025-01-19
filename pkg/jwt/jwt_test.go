package jwt_test

import (
	"github.com/mert-yigittop/cxp-api-starter/pkg/jwt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignAndVerify(t *testing.T) {
	mockSecretKey := "mockSecretKey123"
	os.Setenv("JWT_SECRET_KEY", mockSecretKey)

	mockUserId := uint(12345)
	mockExpiration := 1 * time.Hour

	token, err := jwt.Sign(mockUserId, mockExpiration)
	assert.NoError(t, err, "Sign should not return an error")
	assert.NotEmpty(t, token, "Generated token should not be empty")

	parsedUserId, err := jwt.Verify(token)
	assert.NoError(t, err, "Verify should not return an error for a valid token")
	assert.Equal(t, mockUserId, parsedUserId, "Parsed userId should match the original userId")

	expiredToken, err := jwt.Sign(mockUserId, -1*time.Second)
	assert.NoError(t, err, "Sign should not return an error for expired token")
	_, err = jwt.Verify(expiredToken)
	assert.Error(t, err, "Verify should return an error for an expired token")
}

func TestVerifyInvalidToken(t *testing.T) {
	mockSecretKey := "mockSecretKey123"
	os.Setenv("JWT_SECRET_KEY", mockSecretKey)

	invalidToken := "invalid.token.string"
	_, err := jwt.Verify(invalidToken)
	assert.Error(t, err, "Verify should return an error for an invalid token")
}
