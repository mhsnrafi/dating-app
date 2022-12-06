package routes

import (
	"github.com/gin-gonic/gin"
	"mitte-task/controllers"
	"mitte-task/middlewares"
)

func Swipe(router *gin.RouterGroup) {
	auth := router.Group("/")
	{
		auth.GET(
			"swipe",
			middlewares.JWTMiddleware(),
			controllers.Swipe,
		)
	}
}
