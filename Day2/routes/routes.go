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
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.GET("/users", controllers.GetUsersControllers)

	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books", controllers.GetBooksController)
	return e
}
