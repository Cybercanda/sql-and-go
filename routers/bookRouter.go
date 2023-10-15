package routers

import (
	"github.com/gin-gonic/gin"
	"sql-and-go/controllers"
)

func StartServer() *gin.Engine {
	controllers.InitDB()

	router := gin.Default()

	router.POST("/books", controllers.AddBook)

	return router
}

func CloseDB() {
	controllers.CloseDB()
}

func main() {

}
