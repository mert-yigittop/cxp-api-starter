package repository

import (
	"context"
	"testing"

	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.User{})
	assert.NoError(t, err)

	repo := New(db)

	t.Run("Register Success", func(t *testing.T) {
		user := entity.User{
			Username: "testuser",
			Password: "testpassword",
			Email:    "test@example.com",
		}

		resp, err := repo.Register(context.Background(), user)
		assert.NoError(t, err)
		assert.NotZero(t, resp.ID)
	})

	t.Run("Login Success", func(t *testing.T) {
		payload := dto.LoginRequest{
			Username: "testuser",
			Password: "testpassword",
		}

		userID, err := repo.Login(context.Background(), payload)
		assert.NoError(t, err)
		assert.NotZero(t, userID)
	})

	t.Run("Login Failure - Incorrect Password", func(t *testing.T) {
		payload := dto.LoginRequest{
			Username: "testuser",
			Password: "wrongpassword",
		}

		userID, err := repo.Login(context.Background(), payload)
		assert.Error(t, err)
		assert.Equal(t, uint(0), userID)
	})
}
