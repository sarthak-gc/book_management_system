package models

import (
	"github.com/sarthak-gc/book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type BookSummary struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

var bookSummary BookSummary

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b).Scan(&bookSummary)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books).Scan(&bookSummary)
	return Books
}

func GetBookById(id int64) Book {
	var Book Book
	db.Where("ID=?", id).First(&Book)

	return Book
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book).Scan(&bookSummary)
	return book
}

func UpdateBook(id int64, updatedData Book) Book {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return Book{}
	}
	if updatedData.Author != "" {
		book.Author = updatedData.Author
	}
	if updatedData.Name != "" {
		book.Name = updatedData.Name
	}

	db.Save(&book)
	return book

}
