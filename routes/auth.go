package routes

import (
	"github.com/gin-gonic/gin"
	"mitte-task/controllers"
	"mitte-task/middlewares/validators"
)

func AuthRoute(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST(
			"/login",
			validators.LoginValidator(),
			controllers.Login,
		)

		auth.POST(
			"/refresh",
			validators.RefreshValidator(),
			controllers.Refresh,
		)
	}
}
