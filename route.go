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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Authentication routes
	e.POST("/register", auth_controller.Register)
	e.POST("/login", auth_controller.Login)
	e.GET("/validate-token", auth_controller.ValidateToken, auth_controller.JWTMiddleware)
	e.POST("/refresh-token", auth_controller.RefreshToken)

	e.GET("/user/:userId", user_controller.GetUser, auth_controller.JWTMiddleware)
	e.POST("/user", user_controller.CreateUser, auth_controller.JWTMiddleware)

	e.GET("/category", category_controller.GetAllCategory, auth_controller.JWTMiddleware)
	e.POST("/category", category_controller.CreateCategory, auth_controller.JWTMiddleware)

	// store for seller
	e.GET("/stores/:storeId", store_controller.GetMyStore, auth_controller.JWTMiddleware)
	e.POST("/stores", store_controller.CreateStore, auth_controller.JWTMiddleware)
	e.PUT("/stores/:storeId", store_controller.UpdateStore, auth_controller.JWTMiddleware)
	e.DELETE("/stores/:storeId", store_controller.CloseStore, auth_controller.JWTMiddleware)

	// store for buyer
	e.GET("/home/:storeId", store_controller.GetStore, auth_controller.JWTMiddleware)

	// product for seller
	e.GET("/stores/:storeId/products", product_controller.GetAllMyProduct, auth_controller.JWTMiddleware)
	e.GET("/stores/:storeId/products/:productId", product_controller.GetMyProductDetail, auth_controller.JWTMiddleware)
	e.POST("/stores/:storeId/products", product_controller.CreateProduct, auth_controller.JWTMiddleware)
	e.PUT("/stores/:storeId/products/:productId", product_controller.UpdateProduct, auth_controller.JWTMiddleware)
	e.DELETE("/stores/:storeId/products/:productId", product_controller.DeleteProduct, auth_controller.JWTMiddleware)

	// product for buyer
	e.GET("/products", product_controller.GetAllProduct, auth_controller.JWTMiddleware)

	e.GET("/storePickupPlace/:storeId", store_pickup_place_controller.GetAllStorePickupPlace, auth_controller.JWTMiddleware)
	e.POST("/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace, auth_controller.JWTMiddleware)

<<<<<<< HEAD
=======
	// pickup place 
	e.GET("stores/:storeId/storePickupPlace", store_pickup_place_controller.GetAllStorePickupPlace)
	e.POST("stores/:storeId/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace)
	e.PUT("stores/:storeId/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.UpdateStorePickupPlace)
	e.DELETE("stores/:storeId/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.DeleteStorePickupPlace)
	
>>>>>>> 61aa2773c90cc2ffbb12d8fbbeca84cf752828b4
	// batches
	e.POST("/stores/:storeId/products/:productId", batch_controller.CreateBatch, auth_controller.JWTMiddleware)
	e.PUT("/stores/:storeId/products/:productId/:batchId", batch_controller.UpdateBatch, auth_controller.JWTMiddleware)

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
