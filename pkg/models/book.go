package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
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

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID = ?", id).First(&getBook)
	return &getBook, result
}
func DeleteBook(id int64) (book Book, msg string) {
	db.Where("ID = ?", id).First(&book)
	db.Where("ID = ?", id).Delete(&book)
	msg = "Book deleted successfully"
	return book, msg
}
