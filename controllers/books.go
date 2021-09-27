package controllers

import (
	"net/http"

	"echo-app/models"
	"echo-app/service"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(bookService service.Service) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func SayHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello homies!")
}

func (handler *bookHandler) GetAllBooks(c echo.Context) error {
	books, err := handler.bookService.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   books,
	})
}
