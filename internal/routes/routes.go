package routes

import (
	"log"
	"net/http"
	"sps-backend/internal/config"
	"sps-backend/internal/controllers"
	"sps-backend/internal/utils"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine,
	parkingController *controllers.ParkingController,
	homeController *controllers.HomeController,
	sessionsController *controllers.SessionsController,
	config *config.Config,
) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Cors},
		MaxAge:           86400,
		AllowMethods:     []string{"GET,POST,PUT,DELETE, OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Origin", "Authorization", "Cookie"},
		ExposeHeaders:    []string{"Accept", "Content-Type", "Origin", "Authorization"},
		AllowCredentials: false,
	}))

	// Public routes
	public := router.Group("/api/v1")
	{
		public.GET("/", homeController.Home)
		public.GET("", homeController.Home)
		public.GET("/home", homeController.Health)
		public.GET("/get-token", sessionsController.GetToken)
	}

	protected := router.Group("/api/v1")
	protected.Use(AuthenticateSession(config))
	{
		// User routes
		protected.POST("/parking-inq", parkingController.ParkingInquiry)
	}

	// Protected routes
	// protected := router.Group("/api/v1")
	// protected.Use(middleware.AuthMiddleware(config))
	// {
	// 	// User routes
	// 	protected.POST("/logout", userController.Logout) //
	// 	protected.PUT("/user/geo-location", userController.UpdateGeoLocation)
	// 	protected.PUT("/user/safe-mode", userController.UpdateSafeMode)

	// 	// Account routes
	// 	protected.GET("/accounts", accountController.GetAccounts)

	// 	//Transaction History
	// 	protected.GET("/transactions-history", transactionHistoryController.GetTransactionHistory) //

	// 	// Banks
	// 	protected.GET("/banks", transferController.GetBanks) //

	// 	// Transfer routes
	// 	protected.POST("/transfer/inquiry", transferController.TransferInquiry) //
	// 	protected.POST("/transfer", transferController.Transfer)                //

	// 	// Payment routes
	// 	protected.POST("/payment/va/inquiry", paymentController.PaymentVAInquiry) //
	// 	protected.POST("/payment/va", paymentController.PaymentVA)                //
	// 	protected.POST("/payment/qris", paymentController.PaymentQRIS)            //

	// 	// Merchants
	// 	protected.GET("/merchants", paymentController.GetMerchants) //
	// }
}

func AuthenticateSession(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(ctx, http.StatusUnauthorized, utils.ErrUnauthorized)
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.Error(ctx, http.StatusUnauthorized, utils.ErrUnauthorized)
			ctx.Abort()
			return
		}

		session, err := utils.ValidateSession(tokenString, cfg.JWTSecret)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}

		log.Println(session)
		ctx.Next()
	}
}
