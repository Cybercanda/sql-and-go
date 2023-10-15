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

// Mengambil buku dari ID
func GetBookById(ctx *gin.Context) {
	id := ctx.Param("bookID")
	var book Book

	err := db.QueryRow("SELECT * FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		} else {
			log.Fatal(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// Update Buku berdasarkan ID
func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("bookID")
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := db.Exec("UPDATE books SET title=$1, author=$2, description=$3 WHERE id=$4",
		updateBook.Title, updateBook.Author, updateBook.Description, id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully updated", id),
	})
}

// Menghapus Buku
func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("bookID")

	_, err := db.Exec("DELETE FROM books WHERE id=$1", id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully deleted", id),
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
