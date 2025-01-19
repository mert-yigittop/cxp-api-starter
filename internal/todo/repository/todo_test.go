package repository_test

import (
	"context"
	"testing"

	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupPostgresTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&entity.Todo{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func TestPostgresRepository_Create(t *testing.T) {
	db := setupPostgresTestDB()
	repo := repository.New(db)

	ctx := context.Background()
	todo := entity.Todo{
		Content:   "Test todo",
		Completed: false,
		UserID:    1,
	}

	resp, err := repo.Create(ctx, todo)

	assert.NoError(t, err)
	assert.Equal(t, todo.Content, resp.Content)
	assert.Equal(t, todo.Completed, resp.Completed)

	// Verify data in DB
	var savedTodo entity.Todo
	err = db.First(&savedTodo, resp.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Test todo", savedTodo.Content)
}
