package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mitte-task/models"
	"net/http"
)

func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginRequest models.LoginRequest
		_ = c.ShouldBindBodyWith(&loginRequest, binding.JSON)

		if err := loginRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func RefreshValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var refreshRequest models.RefreshRequest
		_ = c.ShouldBindBodyWith(&refreshRequest, binding.JSON)

		if err := refreshRequest.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}
