package api

import (
	"github.com/nurmuhammaddeveloper/API/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/nurmuhammaddeveloper/API/api/docs" // for swagger
)

type handler struct {
	storage *storage.DBManager
}

// @title           Swagger for Student project api
// @version         1.0
// @description     This is a student project api.
// @host      localhost:8000
func NewServer(storage *storage.DBManager) *gin.Engine {
	router := gin.Default()
	hand := handler{
		storage: storage,
	}
	router.GET("/students/:id", hand.Get)
	router.GET("/students", hand.GetAll)
	router.POST("/students", hand.Create)
	router.PUT("/students/:id", hand.Udate)
	router.DELETE("/students/:id", hand.Delete)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
