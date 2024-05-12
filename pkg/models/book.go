package models

import (
	"github.com/Yash-sudo-web/gosql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int) Book {
	var book Book
	db.Where("id = ?", id).Delete(book)
	return book
}
