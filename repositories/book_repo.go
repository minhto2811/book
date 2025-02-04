package repositories

import "book/models"

type BookRepo interface {
Create(book *models.Book) error				
	UpdateAllField(book *models.Book) error
	GetAll(books *[]models.Book) error
	GetById(id string, book *models.Book) error
	DeleteById(id string) error
}
