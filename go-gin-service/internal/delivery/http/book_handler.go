package http

import (
	"net/http"
	"strconv"

	"digital-library-go/internal/domain"
	"digital-library-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
}

func NewBookHandler(u *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: u}
}

func (h *BookHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/books", h.CreateBook)
	r.GET("/books", h.GetAllBooks)
	r.GET("/books/:id", h.GetBook)
	r.PUT("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book domain.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.usecase.Create(book)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Book creation in background",
	})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.usecase.GetAll())
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, ok := h.usecase.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.usecase.Update(id, book) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if !h.usecase.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
