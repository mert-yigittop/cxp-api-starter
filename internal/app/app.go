package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mert-yigittop/cxp-api-starter/internal/todo"
	"github.com/mert-yigittop/cxp-api-starter/internal/user"
	"github.com/mert-yigittop/cxp-api-starter/pkg/database"
	"log"
)

func Start() {
	// Initial Database
	db, err := database.ConnectToPostgresql()
	if err != nil {
		log.Fatal(err)
	}

	// New Fiber App
	app := fiber.New()

	// Cors Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Logger
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} ${latency} ${method} ${path}\n",
	}))

	// Metrics
	app.Get("/metrics", monitor.New(monitor.Config{Title: "CXP API Starter Monitor"}))

	// Handler Version
	version := app.Group("/api/v1") // V1

	// User
	user.Setup(version, db)

	// Todos
	todo.Setup(version, db)

	// Listening http port on :8080
	log.Fatal(app.Listen(":8080"))
}
