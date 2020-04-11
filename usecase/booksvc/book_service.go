package booksvc

import "github.com/beckojr/go-clean-architecture/domain/model"

// BookService input port
type BookService interface {
	AllBooks(b []*model.Book) ([]*model.Book, error)
	FindBook(b *model.Book) (*model.Book, error)
	CreateBook(b *model.Book) (*model.Book, error)
	UpdateBook(b *model.Book) (*model.Book, error)
	DeleteBook(b *model.Book) (*model.Book, error)
}

// BookRepository implemented in database package
type BookRepository interface {
	FindAllBooks(b []*model.Book) ([]*model.Book, error)
	FindBook(b *model.Book) (*model.Book, error)
	SaveBook(b *model.Book) (*model.Book, error)
	EditBook(b *model.Book) (*model.Book, error)
	RemoveBook(b *model.Book) (*model.Book, error)
}
