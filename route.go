package main

import (
	auth_controller "backend/auth_service"
	batch_controller "backend/batch_service"
	cart_controller "backend/cart_service"
	category_controller "backend/category_service"
	product_controller "backend/product_service"
	store_pickup_place_controller "backend/store_pickup_place_service"
	store_controller "backend/store_service"
	transaction_controller "backend/transaction_service"
	user_controller "backend/user_service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Replace with your frontend origin
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Authentication routes
	e.POST("/register", auth_controller.Register)
	e.POST("/login", auth_controller.Login)
	e.GET("/validate-token", auth_controller.ValidateToken, auth_controller.JWTMiddleware)
	e.POST("/refresh-token", auth_controller.RefreshToken)

	e.GET("/profile", user_controller.GetProfile, auth_controller.JWTMiddleware)
	// Routes
	e.PUT("/update-password", user_controller.UpdatePassword, auth_controller.JWTMiddleware)
	e.PUT("/update-whatsapp", user_controller.UpdateWhatsappNumber, auth_controller.JWTMiddleware)
	e.PUT("/admin/update-whatsapp", user_controller.UpdateWhatsappNumber, auth_controller.JWTMiddleware)
	e.PUT("/admin/update-password", user_controller.UpdatePassword, auth_controller.JWTMiddleware)
	e.PUT("/admin/update-is-admin", user_controller.UpdateIsAdmin, auth_controller.JWTMiddleware)
	e.GET("/admin/users", user_controller.GetAllUsers, auth_controller.JWTMiddleware)

	e.GET("/user/:userId", user_controller.GetUser, auth_controller.JWTMiddleware)
	e.POST("/user", user_controller.CreateUser, auth_controller.JWTMiddleware)

	e.GET("/category", category_controller.GetAllCategory, auth_controller.JWTMiddleware)
	e.POST("/category", category_controller.CreateCategory, auth_controller.JWTMiddleware)

	// store for seller
	e.GET("/store/check/:userId", store_controller.CheckUserHasStore, auth_controller.JWTMiddleware)
	e.GET("/stores/:storeId", store_controller.GetMyStore, auth_controller.JWTMiddleware)
	e.POST("/stores", store_controller.CreateStore, auth_controller.JWTMiddleware)
	e.PUT("/stores/:storeId", store_controller.UpdateStore, auth_controller.JWTMiddleware)
	e.DELETE("/stores/:storeId", store_controller.CloseStore, auth_controller.JWTMiddleware)

	// store for buyer
	e.GET("/home/:storeId", store_controller.GetStore, auth_controller.JWTMiddleware)

	// product for seller
	e.GET("/store/products", product_controller.GetAllMyProduct, auth_controller.JWTMiddleware)
	e.GET("/store/products/:productId", product_controller.GetMyProductDetail, auth_controller.JWTMiddleware)
	e.POST("/store/products", product_controller.CreateProduct, auth_controller.JWTMiddleware)
	e.PUT("/store/products/:productId", product_controller.UpdateProduct, auth_controller.JWTMiddleware)
	e.DELETE("/store/products/:productId", product_controller.DeleteProduct, auth_controller.JWTMiddleware)

	// product for buyer
	e.GET("/products", product_controller.GetAllProduct, auth_controller.JWTMiddleware)
	e.GET("/store/:id/products", product_controller.GetAllProduct, auth_controller.JWTMiddleware)

	// pickup place
	e.GET("/store/storePickupPlace", store_pickup_place_controller.GetAllStorePickupPlace)
	e.POST("/store/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace)
	e.PUT("/store/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.UpdateStorePickupPlace)
	e.DELETE("/store/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.DeleteStorePickupPlace)

	// batches
	e.POST("/store/products/:productId/batch", batch_controller.CreateBatch, auth_controller.JWTMiddleware)
	e.PUT("/store/products/:productId/batch/:batchId", batch_controller.UpdateBatch, auth_controller.JWTMiddleware)

	// Transaction routes
	e.GET("/transaction/:transactionID", transaction_controller.GetTransactionByID, auth_controller.JWTMiddleware)
	e.GET("/transactions/user/:userID", transaction_controller.GetTransactionsByUserID, auth_controller.JWTMiddleware)
	e.GET("/transactions/product/:productID", transaction_controller.GetTransactionsByProductID, auth_controller.JWTMiddleware)
	e.GET("/transactions/batch/:batchID", transaction_controller.GetTransactionsByBatchID, auth_controller.JWTMiddleware)
	e.POST("/transaction", transaction_controller.CreateTransaction, auth_controller.JWTMiddleware)
	e.PUT("/transaction/:transactionID", transaction_controller.UpdateTransaction, auth_controller.JWTMiddleware)
	e.DELETE("/transaction/:transactionID", transaction_controller.DeleteTransaction, auth_controller.JWTMiddleware)

	e.GET("/cart/:cartID", cart_controller.GetCartByID, auth_controller.JWTMiddleware)
	e.GET("/cart/user", cart_controller.GetCartsByUserID, auth_controller.JWTMiddleware)
	e.POST("/cart", cart_controller.CreateCart, auth_controller.JWTMiddleware)
	e.PUT("/cart/:cartID", cart_controller.UpdateCart, auth_controller.JWTMiddleware)
	e.DELETE("/cart/:cartID", cart_controller.DeleteCart, auth_controller.JWTMiddleware)
	return e
}
