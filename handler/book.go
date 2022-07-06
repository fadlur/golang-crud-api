package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Fadlur",
		"title": "Buku",
	})
}

func HelloHandler(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Fadlur",
		"title": "Hello from the other side",
	})
}

func BooksHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
		"title":title,
	})
}

func QueryHandler(ctx *gin.Context)  {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func PostBooksHandler(ctx *gin.Context)  {
	var bookInput book.BookRequest

	
	// err := ctx.BindJSON(&bookInput)
	if err := ctx.ShouldBindJSON(&bookInput); err != nil {
		// log.Fatal(err)
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
	// ctx.IndentedJSON(http.StatusCreated, bookInput)
}