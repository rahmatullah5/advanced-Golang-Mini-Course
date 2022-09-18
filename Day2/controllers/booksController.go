package controllers

import (
	"Day2/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var BookResponse models.BookResponse
var BooksResponse models.BooksResponse
var CreateBookResponse models.CreateBookResponse
var DeleteBookResponse models.DeleteBookResponse

func GetBooksController(c echo.Context) error {
	resp, e := http.Get("https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/book")

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if err := json.Unmarshal(body, &BooksResponse); err != nil { // Parse []byte to go struct pointer
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  BooksResponse,
	})
}

func CreateBookController(c echo.Context) error {
	req, e := http.NewRequest("POST", "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/book/", nil)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(bodyBytes, &CreateBookResponse); err != nil { // Parse []byte to go struct pointer
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   CreateBookResponse,
	})
}

func GetBookController(c echo.Context) error {
	resp, e := http.Get("https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/book/1")

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if err := json.Unmarshal(body, &BookResponse); err != nil { // Parse []byte to go struct pointer
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   BookResponse,
	})
}

func UpdateBookController(c echo.Context) error {
	req, e := http.NewRequest("PUT", "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/book/1", nil)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(bodyBytes, &BookResponse); err != nil { // Parse []byte to go struct pointer
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   BookResponse,
	})
}

func DeleteBookController(c echo.Context) error {
	req, e := http.NewRequest("DELETE", "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/book/1", nil)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(bodyBytes, &DeleteBookResponse); err != nil { // Parse []byte to go struct pointer
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   DeleteBookResponse,
	})
}
