package router

import (
	"github.com/beckojr/go-clean-architecture/adapter"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New create a new router
func New(e *echo.Echo, c adapter.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", func(context echo.Context) error { return c.ListBooks(context) })
	e.GET("/books/:id", func(context echo.Context) error { return c.GetBook(context) })
	e.POST("/books", func(context echo.Context) error { return c.NewBook(context) })
	e.PUT("/books/:id", func(context echo.Context) error { return c.UpdateBook(context) })
	e.DELETE("/books/:id", func(context echo.Context) error { return c.DeleteBook(context) })

	return e
}
