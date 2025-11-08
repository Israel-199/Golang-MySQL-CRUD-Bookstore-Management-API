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

// CreateBook adds a new book record to the database
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById retrieves a single book by its ID
func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID = ?", id).First(&getBook)
	return &getBook, result
}

// DeleteBook deletes a book by ID
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}
