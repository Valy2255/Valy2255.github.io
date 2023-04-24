package models

import (
	"github.com/jinzhu/gorm"
)

//var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Image       []byte `json:"-"`
}



func (b *Book) CreateBook(db *gorm.DB) *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks(db *gorm.DB) []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(db *gorm.DB, Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db = db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(db *gorm.DB, ID int64) Book {
	var book Book
	db.Where("Id=?", ID).Delete(&book)
	return book
}
