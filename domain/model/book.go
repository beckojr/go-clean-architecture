package model

import "time"

// Book entity
type Book struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	Edition     string     `json:"edition"`
	CreatedAt   *time.Time `json:"created_at"`
	ModifiedAt  *time.Time `json:"modified_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// TableName returns the entity's table name
func (b Book) TableName() string {
	return "books"
}
