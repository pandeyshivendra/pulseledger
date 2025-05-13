package main

import (
	"pulseledger/config"
	"pulseledger/db"
	"pulseledger/handlers"
	"pulseledger/repositories"
	"pulseledger/services"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func InitApp() *fiber.App {
	app := fiber.New()

	// Load configuration
	config.Load()

	// Initialize database
	database := db.Init()

	// Initialize repositories
	accountRepo := repositories.NewAccountRepository(database)
	transactionRepo := repositories.NewTransactionRepository(database)

	// Initialize services
	accountService := services.NewAccountService(accountRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	// Initialize handlers
	accountHandler := handlers.NewAccountHandler(accountService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Register routes
	api := app.Group("/api/v1")
	accountHandler.RegisterRoutes(api)
	transactionHandler.RegisterRoutes(api)

	return app
}

func main() {
	app := InitApp()

	// Start server
	port := config.GetConfig().GetPort()
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
