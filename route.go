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

	e.GET("/store/:storeId", store_controller.GetStoreById)
	e.POST("/store", store_controller.CreateStore)

	e.GET("/product/:storeId", product_controller.GetAllMyProduct)
	e.POST("/product", product_controller.CreateProduct)

	e.GET("/storePickupPlace/:storeId", store_pickup_place_controller.GetAllStorePickupPlace)
	e.POST("/storePickupPlace", store_pickup_place_controller.CreateStorePickupPlace)

	e.POST("/batch", batch_controller.CreateBatch)

	// e.GET("/user/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	return e
}
