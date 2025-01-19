package repository

import (
	"context"
	"testing"

	"github.com/mert-yigittop/cxp-api-starter/internal/todo/dto"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTodoRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Todo{})
	assert.NoError(t, err)

	repo := New(db)

	t.Run("Create Todo", func(t *testing.T) {
		todo := entity.Todo{
			Content:   "Test Todo",
			Completed: false,
			UserID:    1,
		}

		resp, err := repo.Create(context.Background(), todo)
		assert.NoError(t, err)
		assert.NotZero(t, resp.ID)
		assert.Equal(t, "Test Todo", resp.Content)
		assert.False(t, resp.Completed)
	})

	t.Run("Get Todo List", func(t *testing.T) {
		payload := dto.GetTodoListRequest{
			UserID: 1,
		}

		resp, err := repo.GetList(context.Background(), payload)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Todos)
		assert.Equal(t, 1, len(resp.Todos))
	})

	t.Run("Update Todo", func(t *testing.T) {
		payload := dto.UpdateTodoRequest{
			ID:        1,
			Content:   "Updated Todo",
			Completed: true,
		}

		resp, err := repo.Update(context.Background(), payload, 1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), resp.ID)
		assert.Equal(t, "Updated Todo", resp.Content)
		assert.True(t, resp.Completed)
	})

	t.Run("Update Todo - Unauthorized", func(t *testing.T) {
		payload := dto.UpdateTodoRequest{
			ID:        1,
			Content:   "Another Update",
			Completed: true,
		}

		_, err := repo.Update(context.Background(), payload, 2)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not authorized")
	})

	t.Run("Delete Todo", func(t *testing.T) {
		payload := dto.DeleteTodoRequest{
			ID: 1,
		}

		resp, err := repo.Delete(context.Background(), payload, 1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), resp.ID)
	})

	t.Run("Delete Todo - Unauthorized", func(t *testing.T) {
		todo := entity.Todo{
			Content:   "Unauthorized Test Todo",
			Completed: false,
			UserID:    1,
		}
		err := db.Create(&todo).Error
		assert.NoError(t, err)

		payload := dto.DeleteTodoRequest{
			ID: todo.ID,
		}

		_, err = repo.Delete(context.Background(), payload, 2)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not authorized")
	})

}
