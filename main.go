package main

import (
	"fmt"
	"log"
	"pustaka-api/book"

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
	// books, err := bookRepository.FindAll()

	// for _, item := range books {
	// 	fmt.Printf("Object itemnya %v\n", item)
	// }

	// book, err := bookRepository.FindByID(2)
	// fmt.Printf("Objectnya %v\n", book)
	// db.AutoMigrate(&book.Book{})

	/*
	CREATE DATA
	*/
	bookRequest := book.BookRequest{
		Title : "Judulnya",
		Price : "12000",
		// Decription : "Ini description",
		// Rating : 5,
	}
	

	bookService.Create(bookRequest)
	// fmt.Printf("Objectnya %v\n", book)

	
	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("Error createing book record ", err)
	// }

	/*
	GET DATA
	*/

	// var book book.Book
	// err = db.Debug().First(&book, 1).Error
	// if err != nil {
	// 	fmt.Println("Error finding book", err)
	// }

	// fmt.Println("Title : ", book.Title)
	// fmt.Printf("book objet %v\n", book)

	/*
	UPDATE DATA
	*/

	// var book book.Book

	// err = db.Debug().Where("id=?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding data: ", err)
	// }

	// book.Title = "Crouching tiger hidden dragon"
	
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error update data", err)
	// }

	/*
	DELETE DATA	
	*/

	// var book book.Book
	// err = db.Debug().Where("id=?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error finding data: ", err)
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error update data", err)
	// }

	// router := gin.Default()

	// v1 := router.Group("/v1")
	// router.GET("/", handler.RootHandler)
	// v1.GET("/hello", handler.HelloHandler)
	// v1.GET("/books/:id/:title", handler.BooksHandler)
	// v1.GET("/query", handler.QueryHandler)
	// v1.POST("books", handler.PostBooksHandler)
	// router.Run("localhost:9000")
}

