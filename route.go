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

	// Rute for store
	e.GET("/stores/:storeId", store_controller.GetMyStore)
	e.POST("/stores", store_controller.CreateStore)

	// Rute for product
	e.GET("/stores/:storeId/products", product_controller.GetAllMyProduct)
	e.GET("/stores/:storeId/products/:productId", product_controller.GetMyProductDetail)
	e.PUT("/stores/:storeId/products/:productId", product_controller.UpdateProduct)
	e.POST("/stores/:storeId/products", product_controller.CreateProduct)

	e.GET("/storePickupPlace/:storeId", store_pickup_place_controller.GetAllStorePickupPlace)
	e.POST("/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace)
	
	// Create batches
	e.POST("/stores/:storeId/products/:productId", batch_controller.CreateBatch)
	
	// product
	e.GET("/products", product_controller.GetAllProduct)


	// e.GET("/user/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	return e
}
