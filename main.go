package main

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/repository"
	"echo-app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()

	//setup repo
	bookRepository := repository.NewRepo(db)

	//setup service
	bookService := service.NewBookService(bookRepository)

	//setup controller
	bookController := controllers.NewBookHandler(bookService)

	//routes
	v1 := e.Group("/v1")
	v1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1.GET("/books", bookController.GetAllBooks)
	v1.POST("/books", bookController.StoreBookHandler)
	v1.GET("/books/:id", bookController.FindBook)
	v1.PUT("/books/:id", bookController.UpdateHandler)
	v1.DELETE("/books/:id", bookController.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
