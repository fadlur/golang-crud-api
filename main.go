package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	
	dsn := "root:@tcp(127.0.0.1:3306)/pustakaapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// _ =db
	if err != nil {
		log.Fatal("Db Connection failed")
	}
	_ = db
	fmt.Println("DB Connected")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("books", bookHandler.CreateBook)
	v1.GET("books", bookHandler.GetBooks)
	v1.GET("books/:id", bookHandler.GetBook)
	v1.PUT("books/:id", bookHandler.UpdateBook)
	v1.DELETE("books/:id", bookHandler.DeleteBook)
	router.Run("localhost:9000")
}

