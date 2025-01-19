package user

import (
	"github.com/gofiber/fiber/v2"
	usersHandler "github.com/mert-yigittop/cxp-api-starter/internal/user/controller/v1"
	userEntities "github.com/mert-yigittop/cxp-api-starter/internal/user/entity"
	usersRepository "github.com/mert-yigittop/cxp-api-starter/internal/user/repository"
	usersUsecase "github.com/mert-yigittop/cxp-api-starter/internal/user/usecase"
	"github.com/mert-yigittop/cxp-api-starter/pkg/middleware"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {
	// Migration
	err := db.AutoMigrate(&userEntities.User{})
	if err != nil {
		panic("USER Migration failed")
	}

	// Dependency Injection
	repo := usersRepository.New(db)
	usecase := usersUsecase.New(repo)
	handler := usersHandler.New(usecase)

	// Routes
	route := router.Group("/auth")

	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
	route.Get("/logout", middleware.AuthRequired(), handler.Logout)
}
