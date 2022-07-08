package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler  {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(ctx *gin.Context)  {
	books, err := h.bookService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(ctx *gin.Context)  {
	idString := ctx.Param("id")

	id, _ := strconv.Atoi(idString)
	bookRes, err := h.bookService.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	bookResponse := convertToBookResponse(bookRes)

	ctx.JSON(http.StatusOK, gin.H {
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(ctx *gin.Context)  {
	var bookRequest book.BookRequest	
	// err := ctx.BindJSON(&BookRequest)
	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
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

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
	// ctx.IndentedJSON(http.StatusCreated, bookInput)
}

func (h *bookHandler) UpdateBook(ctx *gin.Context)   {
	var bookRequest book.BookRequest

	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)


	bookRes, err := h.bookService.Update(bookRequest, id)
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H {
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book":convertToBookResponse(bookRes),
	})
}

func (h *bookHandler) DeleteBook(ctx *gin.Context)  {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	bookRes, err := h.bookService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H {
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookRes,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse  {
	return book.BookResponse{
		ID: b.ID,
		Title: b.Title,
		Price: b.Price,
		Decription: b.Decription,
		Rating: b.Rating,
	}
}