package main

import (
	"fmt"
	"interview/handlers"
	"interview/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	r := InitializeRouter()

	PORT := fmt.Sprintf(":%s", os.Getenv("WEB_PORT"))
	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	db := models.SetupDatabase()

	uh := handlers.NewUserHandler(db)

	users := r.Group("/users")
	{
		users.GET("/", uh.GetUsers)
		users.POST("/", uh.CreateUser)
		users.GET("/:id", uh.GetUser)
		// users.PUT("/:id", uh.UpdateUser)
		// users.DELETE("/:id", uh.DeleteUser)
	}

	bh := handlers.NewBookHandler(db)

	books := r.Group("/books")
	{
		books.GET("/", bh.GetBooks)
		books.POST("/", bh.CreateBook)
		books.GET("/:id", bh.GetBook)
		// books.PUT("/:id", bh.UpdateBook)
		// books.DELETE("/:id", bh.DeleteBook)
	}

	ubh := handlers.NewUserBookHandler(db)

	borrow := r.Group("/borrow")
	{
		borrow.GET("/", ubh.GetUserBooks)
		borrow.POST("/", ubh.CreateUserBook)
	}

	return r
}
