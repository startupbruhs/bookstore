package repository

import (
	"startupbruhs.github.io/bookstore/db/model/entity"

	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	Conn *gorm.DB
}

func (repo *BookRepo) All() []entity.Book {
	var books []entity.Book
	repo.Conn.Find(&books)
	return books
}

func (repo *BookRepo) Create(book *entity.Book) {
	repo.Conn.Create(&book)
}

func InitBookRepo(database *gorm.DB) *UserRepo {
	return &UserRepo{Conn: database}
}
