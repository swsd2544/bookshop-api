package models

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Phone string `json:"phone" gorm:"unique"`
	Books []Book `json:"-" gorm:"many2many:user_books;"`
}
