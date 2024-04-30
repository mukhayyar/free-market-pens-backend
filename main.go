package main

import (
	"net/http"

	user_controller "backend/user_service/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:id", user_controller.GetUser)
	e.Logger.Fatal(e.Start(":1323"))
}
