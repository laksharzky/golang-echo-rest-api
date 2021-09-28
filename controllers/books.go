package controllers

import (
	"net/http"
	"strconv"

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

func (handler *bookHandler) StoreBookHandler(c echo.Context) error {
	var bookRequest models.BookRequest
	err := c.Bind(&bookRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "bad request body",
		})
	}

	book, err := handler.bookService.Create(bookRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.WebResponse{
		Code: http.StatusOK,
		Data: book,
	})
}

func (handler *bookHandler) FindBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
		})
	}

	book, err := handler.bookService.FindById(id)
	return c.JSON(http.StatusOK, models.WebResponse{
		Code: http.StatusOK,
		Data: book,
	})
}

func (handler *bookHandler) UpdateHandler(c echo.Context) error {
	var bookRequest models.BookRequest

	err := c.Bind(&bookRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
		})
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := handler.bookService.Update(id, bookRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request Body",
		})
	}

	return c.JSON(http.StatusOK, models.WebResponse{
		Code: http.StatusOK,
		Data: book,
	})

}

func (handler *bookHandler) DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
		})
	}
	b, err := handler.bookService.Delete(int(id))
	return c.JSON(http.StatusOK, models.WebResponse{
		Code: http.StatusOK,
		Data: b,
	})
}
