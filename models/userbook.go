package models

import "time"

type UserBook struct {
	UserID  int `json:"userId" gorm:"primaryKey"`
  	BookID int `json:"bookId" gorm:"primaryKey"`
	NumberOfDays uint `json:"numberOfDays"`
	BorrowDate time.Time `json:"borrowDate"`
	ReturnDate time.Time `json:"returnDate"`
}

