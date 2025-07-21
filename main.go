package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sps-backend/internal/clients"
	"sps-backend/internal/config"
	"sps-backend/internal/controllers"
	"sps-backend/internal/routes"
	"sps-backend/internal/services"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// 	endpointURL := "/riskService/auth/v1/access_token"
	// 	httpMethod := "POST"
	// 	requestBody := `{
	//     "grantType": "client_credentials",
	//     "additionalInfo": "{}"
	// }`

	// 	now := time.Now()

	// 	signature, timestamp, err := utils.GenerateSignatureForGetToken(cfg.ClientSecret, endpointURL, httpMethod, requestBody, now)
	// 	if err != nil {
	// 		fmt.Printf("Error: %v\n", err)
	// 		return
	// 	}

	// 	fmt.Println()
	// 	fmt.Println("X-TIMESTAMP:", timestamp)
	// 	fmt.Println("X-CLIENT-KEY", cfg.ClientKey)
	// 	fmt.Println("X-Signature:", signature)
	// 	fmt.Println()

	// 	utils.GenerateSignature2(now)

	// Initialize database connection

	// db, err := config.InitializeDB(cfg)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer func() {
	// 	if err := db.Close(); err != nil {
	// 		log.Printf("Error closing database connection: %v", err)
	// 	}
	// }()

	// FDS
	// authClient := clients.NewAuthClient(cfg.FDSBaseURL2, cfg.ClientKey, cfg.ClientSecret)
	spsClient := clients.NewSPSClient(cfg.SPSBaseURL, cfg.ClientKey, cfg.ClientSecret)

	// Initialize repositories
	// userRepo := repositories.NewUserRepository(db)
	// accountRepo := repositories.NewAccountRepository(db)
	// transferRepo := repositories.NewTransferRepository(db)
	// paymentRepo := repositories.NewPaymentRepository(db)
	// trxHistoryRepo := repositories.NewTransactionHistoryRepository(db)

	// Initialize services
	// userService := services.NewUserService(userRepo, cfg.JWTSecret, authClient, fdsClient)
	// accountService := services.NewAccountService(accountRepo)
	// transferService := services.NewTransferService(transferRepo, accountRepo, authClient, fdsClient)
	// paymentService := services.NewPaymentService(paymentRepo, accountRepo, fdsClient, authClient)
	// trxHistoryService := services.NewTransactionHistoryService(trxHistoryRepo)
	spsServices := services.NewParkingService(spsClient, cfg)

	// Initialize controllers
	// userController := controllers.NewUserController(userService)
	// accountController := controllers.NewAccountController(accountService)
	// transferController := controllers.NewTransferController(transferService)
	// paymentController := controllers.NewPaymentController(paymentService)
	// trxHistoryController := controllers.NewTransactionHistoryController(trxHistoryService)
	spsController := controllers.NewParkingController(spsServices)
	homeController := controllers.NewHomeController("")

	// Initialize Gin router
	// if !cfg.Debug {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	router := gin.Default()

	// router.Use(middleware.GlobalTimeoutMiddleware(15 * time.Second))

	// Set up routes
	routes.SetupRoutes(router, spsController, homeController, cfg)

	// Configure HTTP server with timeouts
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
		// Timeout configurations
		// ReadTimeout:  20 * time.Second, // Time to read request headers and body
		// WriteTimeout: 30 * time.Second, // Time to write response
		// IdleTimeout:  120 * time.Second, // Time to keep idle connections
	}

	// Run server in a goroutine so we can handle shutdown
	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal
	<-quit
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	} else {
		log.Println("Server shutdown complete")
	}
}
