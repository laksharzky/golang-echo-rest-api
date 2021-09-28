package service

import (
	"echo-app/models"
	"echo-app/repository"
)

type Service interface {
	FindAll() ([]models.Book, error)
	FindById(ID int) (models.Book, error)
	Create(bookRequest models.BookRequest) (models.Book, error)
	Update(ID int, bookRequest models.BookRequest) (models.Book, error)
	Delete(ID int) (models.Book, error)
}

type bookService struct {
	repository repository.Repository
}

func NewBookService(repository repository.Repository) *bookService {
	return &bookService{
		repository,
	}
}

func (s *bookService) FindAll() ([]models.Book, error) {
	return s.repository.FindAll()
}

func (s *bookService) Create(bookRequest models.BookRequest) (models.Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := models.Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *bookService) FindById(ID int) (models.Book, error) {
	return s.repository.FindById(ID)
}

func (s *bookService) Update(ID int, bookRequest models.BookRequest) (models.Book, error) {
	book, err := s.repository.FindById(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err

}

func (s *bookService) Delete(ID int) (models.Book, error) {
	book, err := s.repository.FindById(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
