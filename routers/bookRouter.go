package routers

import (
	"github.com/gin-gonic/gin"
	"sql-and-go/controllers"
)

func StartServer() *gin.Engine {
	controllers.InitDB()

	router := gin.Default()

	router.POST("/books", controllers.AddBook)
	router.GET("/books", controllers.GetBook)
	router.GET("/book/:bookID", controllers.GetBookById)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}

func CloseDB() {
	controllers.CloseDB()
}

func main() {

}
