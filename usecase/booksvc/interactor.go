package booksvc

import (
	"log"

	"github.com/beckojr/go-clean-architecture/domain/model"
)

type bookInteractor struct {
	repository BookRepository
}

// New creates a new interactor
func New(r BookRepository) BookService {
	return &bookInteractor{repository: r}
}

func (i *bookInteractor) AllBooks(b []*model.Book) ([]*model.Book, error) {
	book, err := i.repository.FindAllBooks(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, nil
}

func (i *bookInteractor) FindBook(b *model.Book) (*model.Book, error) {
	book, err := i.repository.FindBook(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, err
}

func (i *bookInteractor) CreateBook(b *model.Book) (*model.Book, error) {
	book, err := i.repository.SaveBook(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, nil
}

func (i *bookInteractor) UpdateBook(b *model.Book) (*model.Book, error) {
	book, err := i.repository.EditBook(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, err
}
func (i *bookInteractor) DeleteBook(b *model.Book) (*model.Book, error) {
	book, err := i.repository.RemoveBook(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return book, nil
}
