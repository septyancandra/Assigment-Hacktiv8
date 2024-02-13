package routers

import (
	//"Mygram/controllers"
	//"Mygram/middlewares"

	"Mygram/controllers"
	"Mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.DeleteUser)
	}

	return r
}
