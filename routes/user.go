package routes

import (
	"github.com/gin-gonic/gin"
	"mitte-task/controllers"
)

func User(router *gin.RouterGroup) {
	auth := router.Group("/")
	{
		auth.POST(
			"user/create",
			controllers.CreateUser,
		)
	}
}
