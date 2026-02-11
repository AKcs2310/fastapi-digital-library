package repository

import "digital-library-go/internal/domain"

type BookRepository struct {
	books map[int]domain.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		books: make(map[int]domain.Book),
	}
}

func (r *BookRepository) Create(book domain.Book) {
	r.books[book.ID] = book
}

func (r *BookRepository) GetAll() []domain.Book {
	result := []domain.Book{}
	for _, b := range r.books {
		result = append(result, b)
	}
	return result
}

func (r *BookRepository) GetByID(id int) (domain.Book, bool) {
	book, ok := r.books[id]
	return book, ok
}

func (r *BookRepository) Update(id int, book domain.Book) bool {
	if _, ok := r.books[id]; !ok {
		return false
	}
	r.books[id] = book
	return true
}

func (r *BookRepository) Delete(id int) bool {
	if _, ok := r.books[id]; !ok {
		return false
	}
	delete(r.books, id)
	return true
}
