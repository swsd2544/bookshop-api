package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	PricePerDay float64 `json:"pricePerDay"`
	Users []User `json:"-" gorm:"many2many:user_books;"`
}
