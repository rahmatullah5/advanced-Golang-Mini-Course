package controllers

import (
	"Day2/lib/database"
	"Day2/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersControllers(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserControllers(c echo.Context) error {
	id := c.Param("id")

	user, e := database.GetUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")

	userValues := models.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	user, e := database.UpdateUser(id, userValues)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
func CreateUserControllers(c echo.Context) error {
	id := c.Param("id")

	userValues := models.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	user, e := database.CreateUser(id, userValues)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")

	user, e := database.DeleteUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
