package bookctlr

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/beckojr/go-clean-architecture/adapter"
	"github.com/beckojr/go-clean-architecture/domain/model"
	"github.com/beckojr/go-clean-architecture/usecase/booksvc"
)

//BookController defines /books endpoint handlers
type BookController interface {
	GetBook(context adapter.Context) error
	ListBooks(context adapter.Context) error
	NewBook(context adapter.Context) error
	UpdateBook(context adapter.Context) error
	DeleteBook(context adapter.Context) error
}

type bookController struct {
	service booksvc.BookService
}

// New create a new book controller
func New(s booksvc.BookService) BookController {
	return &bookController{service: s}
}

// GetBook handler that retrieves a book
func (c *bookController) GetBook(context adapter.Context) error {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"message": "The book ID is malformed"})
	}
	book := &model.Book{ID: bookID}
	b, err := c.service.FindBook(book)
	if err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprint(err)})
	}
	return context.JSON(http.StatusOK, b)
}

// ListBooks handler that list all books
func (c *bookController) ListBooks(context adapter.Context) error {
	var books []*model.Book
	b, err := c.service.AllBooks(books)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, b)
}

// NewBook handler that add a new book to the store
func (c *bookController) NewBook(context adapter.Context) error {
	book := &model.Book{}
	err := context.Bind(book)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"message": "Your payload is malformed"})
	}
	b, err := c.service.CreateBook(book)
	if err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprint(err)})
	}
	return context.JSON(http.StatusCreated, b)

}

// UpdateBook update handler
func (c *bookController) UpdateBook(context adapter.Context) error {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"message": "The book ID is malformed"})
	}
	book := &model.Book{ID: bookID}
	err = context.Bind(book)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"message": "Your payload is malformed"})
	}
	b, err := c.service.UpdateBook(book)
	if err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprint(err)})
	}
	return context.JSON(http.StatusOK, b)
}

// DeleteBook delete handler
func (c *bookController) DeleteBook(context adapter.Context) error {
	bookID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"message": "The book ID is malformed"})
	}
	book := &model.Book{ID: bookID}
	_, err = c.service.DeleteBook(book)
	if err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprint(err)})
	}
	return context.NoContent(http.StatusNoContent)
}
