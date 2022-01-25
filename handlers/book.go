package handlers

import (
	"interview/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	DB *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{db}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	Books := []models.Book{}

	h.DB.Find(&Books)

	c.JSON(http.StatusOK, gin.H{"data": Books})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	Book := models.Book{}

	if err := c.ShouldBindJSON(&Book); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&Book).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, Book)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	Book := models.Book{}

	if err := h.DB.First(&Book, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, Book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	Book := models.Book{}

	if err := h.DB.First(&Book, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&Book); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&Book).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, Book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	Book := models.Book{}

	if err := h.DB.First(&Book, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&Book).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}