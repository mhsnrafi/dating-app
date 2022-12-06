package routes

import (
	"github.com/gin-gonic/gin"
	"mitte-task/controllers"
	"mitte-task/middlewares"
)

func Profile(router *gin.RouterGroup) {
	auth := router.Group("/")
	{
		auth.POST(
			"profile/create",
			middlewares.JWTMiddleware(),
			controllers.CreateProfile,
		)
		auth.GET(
			"profile",
			middlewares.JWTMiddleware(),
			controllers.Profiles,
		)
		auth.GET(
			"/profile/filter",
			middlewares.JWTMiddleware(),
			controllers.FilteredProfiles,
		)
	}
}
