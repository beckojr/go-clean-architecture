package database

import (
	"github.com/beckojr/go-clean-architecture/domain/model"
	"github.com/beckojr/go-clean-architecture/usecase/booksvc"
	"github.com/jinzhu/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

// NewBookRepository create a new instance of bookRepository
func NewBookRepository(db *gorm.DB) booksvc.BookRepository {
	return &bookRepository{DB: db}
}

func (r bookRepository) FindAllBooks(b []*model.Book) ([]*model.Book, error) {
	err := r.DB.Find(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r bookRepository) FindBook(b *model.Book) (*model.Book, error) {
	err := r.DB.Find(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r bookRepository) SaveBook(b *model.Book) (*model.Book, error) {
	err := r.DB.Save(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r bookRepository) EditBook(b *model.Book) (*model.Book, error) {
	err := r.DB.Update(b).Error
	if err != nil {
		return nil, err
	}
	return b, err
}

func (r bookRepository) RemoveBook(b *model.Book) (*model.Book, error) {
	err := r.DB.Delete(b).Error
	if err != nil {
		return nil, err
	}
	return b, err
}
