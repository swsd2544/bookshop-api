package handlers

import (
	"interview/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserBookHandler struct {
	DB *gorm.DB
}

func NewUserBookHandler(db *gorm.DB) *UserBookHandler {
	return &UserBookHandler{db}
}

func (h *UserBookHandler) GetUserBooks(c *gin.Context) {
	userId := c.Query("userId")
	bookId := c.Query("bookId")
	UserBooks := []models.UserBook{}

	if userId != "" && bookId != "" {
		h.DB.Find(&UserBooks, "user_id = ? AND book_id = ?", userId, bookId)
	} else if userId != "" {
		h.DB.Find(&UserBooks, "user_id = ?", userId)
	} else if bookId != "" {
		h.DB.Find(&UserBooks, "book_id = ?", bookId)
	} else {
		h.DB.Find(&UserBooks)
	}

	c.JSON(http.StatusOK, gin.H{"data": UserBooks})
}

func (h *UserBookHandler) CreateUserBook(c *gin.Context) {
	userId := c.Query("userId")
	bookId := c.Query("bookId")
	UserBook := models.UserBook{}

	var convErr error
	if UserBook.UserID, convErr = strconv.Atoi(userId); convErr != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if UserBook.BookID, convErr = strconv.Atoi(bookId); convErr != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&UserBook); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	borrowDate := currentTime
	returnDate := currentTime.AddDate(0, 0, int(UserBook.NumberOfDays))

	UserBook.BorrowDate = borrowDate
	UserBook.ReturnDate = returnDate

	if err := h.DB.Save(&UserBook).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, UserBook)
}
