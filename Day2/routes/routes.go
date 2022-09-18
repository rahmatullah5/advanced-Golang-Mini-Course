package routes

import (
	"Day2/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", controllers.GetUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.GET("/users", controllers.GetUsersControllers)
	e.POST("/users", controllers.CreateUserControllers)

	return e
}
