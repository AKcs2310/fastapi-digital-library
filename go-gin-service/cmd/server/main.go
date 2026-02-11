package main

import (
	"digital-library-go/internal/delivery/http"
	"digital-library-go/internal/repository"
	"digital-library-go/internal/usecase"
	"digital-library-go/pkg/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize layers
	repo := repository.NewBookRepository()
	w := worker.NewWorker(repo)
	u := usecase.NewBookUsecase(repo, w)
	h := http.NewBookHandler(u)

	// gin router
	r := gin.Default()

	// register routes
	h.RegisterRoutes(r)

	// run server
	r.Run(":8080")
}
