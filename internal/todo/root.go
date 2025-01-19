package todo

import (
	"github.com/gofiber/fiber/v2"
	todoHandler "github.com/mert-yigittop/cxp-api-starter/internal/todo/controller/v1"
	todoEntity "github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	todoRepository "github.com/mert-yigittop/cxp-api-starter/internal/todo/repository"
	todoUsecase "github.com/mert-yigittop/cxp-api-starter/internal/todo/usecase"
	"github.com/mert-yigittop/cxp-api-starter/pkg/middleware"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {
	// Migration
	err := db.AutoMigrate(&todoEntity.Todo{})
	if err != nil {
		panic("TODO Migration failed")
	}

	// Dependency Injection
	repo := todoRepository.New(db)
	usecase := todoUsecase.New(repo)
	handler := todoHandler.New(usecase)

	// Routes
	route := router.Group("/todos", middleware.AuthRequired())

	route.Get("/", handler.GetList)
	route.Post("/", handler.Create)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
