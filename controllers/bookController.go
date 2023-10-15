package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

func InitDB() {
	connectingString := "host=localhost port=5433 user=postgres password=databasebily dbname=sql-go sslmode=disable"
	var err error

	db, err = sql.Open("postgres", connectingString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")
}

func CloseDB() {
	db.Close()
	fmt.Println("Connection to database closed")
}

// Book Struct
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func AddBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := db.Exec("INSERT INTO books (title, author, description) VALUES ($1, $2, $3)",
		newBook.Title, newBook.Author, newBook.Description)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"book":    newBook,
	})
}

// Mengambil Semua data Buku
func GetBook(ctx *gin.Context) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}
		books = append(books, book)
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"books": books,
	})
}

// Mengambil semua buku
//func getBooks(c *gin.Context) {
//	rows, err := db.Query("SELECT * FROM books")
//	if err != nil {
//		log.Fatal(err)
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Internal Server Error",
//		})
//		return
//	}
//	defer rows.Close()
//
//	var books []Book
//	for rows.Next() {
//		var book Book
//		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
//
//		if err != nil {
//			log.Fatal(err)
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"error": "Internal Server Error",
//			})
//			return
//		}
//		books = append(books, book)
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"data": books,
//	})
//}

func main() {

}
