package usecase

import (
	"digital-library-go/internal/domain"
	"digital-library-go/internal/repository"
	"digital-library-go/pkg/worker"
)

type BookUsecase struct {
	repo   *repository.BookRepository
	worker *worker.Worker
}

func NewBookUsecase(r *repository.BookRepository, w *worker.Worker) *BookUsecase {
	return &BookUsecase{
		repo:   r,
		worker: w,
	}
}

func (u *BookUsecase) Create(book domain.Book) {
	u.worker.AddJob(book) // async creation
}

func (u *BookUsecase) GetAll() []domain.Book {
	return u.repo.GetAll()
}

func (u *BookUsecase) GetByID(id int) (domain.Book, bool) {
	return u.repo.GetByID(id)
}

func (u *BookUsecase) Update(id int, book domain.Book) bool {
	return u.repo.Update(id, book)
}

func (u *BookUsecase) Delete(id int) bool {
	return u.repo.Delete(id)
}
