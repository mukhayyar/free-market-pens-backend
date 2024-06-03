package main

import (
	batch_controller "backend/batch_service"
	category_controller "backend/category_service"
	product_controller "backend/product_service"
	store_pickup_place_controller "backend/store_pickup_place_service"
	store_controller "backend/store_service"
	user_controller "backend/user_service/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:userId", user_controller.GetUser)
	e.POST("/user", user_controller.CreateUser)

	e.GET("/category", category_controller.GetAllCategory)
	e.POST("/category", category_controller.CreateCategory)

	// store for seller
	e.GET("/stores/:storeId", store_controller.GetMyStore)
	e.POST("/stores", store_controller.CreateStore)
	e.PUT("/stores/:storeId", store_controller.UpdateStore)
	e.DELETE("/stores/:storeId", store_controller.CloseStore)
	
	// store for buyer
	e.GET("/home/:storeId", store_controller.GetStore)
	
	// product for seller
	e.GET("/stores/:storeId/products", product_controller.GetAllMyProduct)
	e.GET("/stores/:storeId/products/:productId", product_controller.GetMyProductDetail)
	e.POST("/stores/:storeId/products", product_controller.CreateProduct)
	e.PUT("/stores/:storeId/products/:productId", product_controller.UpdateProduct)
	e.DELETE("/stores/:storeId/products/:productId", product_controller.DeleteProduct)

	// product for buyer
	e.GET("/products", product_controller.GetAllProduct)

	// pickup place 
	e.GET("stores/:storeId/storePickupPlace", store_pickup_place_controller.GetAllStorePickupPlace)
	e.POST("stores/:storeId/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace)
	e.PUT("stores/:storeId/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.UpdateStorePickupPlace)
	e.DELETE("stores/:storeId/storePickupPlace/:storePickupPlaceId", store_pickup_place_controller.DeleteStorePickupPlace)
	
	// batches
	e.POST("/stores/:storeId/products/:productId", batch_controller.CreateBatch)
	e.PUT("/stores/:storeId/products/:productId/:batchId", batch_controller.UpdateBatch)

	// e.GET("/user/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	return e
}