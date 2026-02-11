package worker

import (
	"digital-library-go/internal/domain"
	"digital-library-go/internal/repository"
)

type Job struct {
	Book domain.Book
}

type Worker struct {
	JobQueue chan Job
	Repo     *repository.BookRepository
}

func NewWorker(repo *repository.BookRepository) *Worker {
	w := &Worker{
		JobQueue: make(chan Job, 100),
		Repo:     repo,
	}

	go w.start()
	return w
}

func (w *Worker) start() {
	for job := range w.JobQueue {
		w.Repo.Create(job.Book)
	}
}

func (w *Worker) AddJob(book domain.Book) {
	w.JobQueue <- Job{Book: book}
}
