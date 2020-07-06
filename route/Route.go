package route

import (
	"Github/go-crud/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	x := gin.Default()
	api := x.Group("/api")
	{
		api.GET("user", controller.GetUsers)
		api.GET("user/:id", controller.GetUserByUsername)
		api.POST("user", controller.SaveUser)
		api.PUT("user/:id", controller.UpdateUser)
		api.DELETE("user/:id", controller.DeleteUser)
	}
	return x
}
