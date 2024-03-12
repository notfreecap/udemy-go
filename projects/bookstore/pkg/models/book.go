package models

import (
	"github.com/jinzhu/gorm"
	"github.com/notfreecap/bookstore/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var b []Book
	db.Find(&b)
	return b
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var b Book
	db.Where("ID=?", id).Delete(b)
	return b
}
