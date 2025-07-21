package routes

import (
	"sps-backend/internal/config"
	"sps-backend/internal/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine,
	parkingController *controllers.ParkingController,
	config *config.Config,
) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Cors},
		MaxAge:           86400,
		AllowMethods:     []string{"GET,POST,PUT,DELETE, OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Origin", "Authorization", "Cookie"},
		ExposeHeaders:    []string{"Accept", "Content-Type", "Origin", "Authorization"},
		AllowCredentials: true,
	}))

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/parking-inq", parkingController.ParkingInquiry)
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
