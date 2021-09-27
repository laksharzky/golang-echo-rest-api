package repository

import (
	"echo-app/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Book, error)
	// FindById(ID int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	// Update(book models.Book) (models.Book, error)
	// Delete(book models.Book) (models.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Delete(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
