package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetData)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	router.GET("/cars", controllers.GetDataAll)

	return router
}
