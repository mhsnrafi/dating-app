package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"mitte-task/models"
	"mitte-task/services"
	"net/http"
)

// Swipe godoc
// @Summary      Swipe
// @Description  Swipe will return the perfect match of the user
// @Tags         swipe
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /swipe/:user_id/:profile_id/:preference [get]
func Swipe(c *gin.Context) {
	userId := c.Request.URL.Query().Get("user_id")
	profileId := c.Request.URL.Query().Get("profile_id")
	preference := c.Request.URL.Query().Get("preference")

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// create user record
	isMatch, err := services.SwipeProfile(userId, profileId, preference)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"match": cast.ToString(isMatch),
	}
	response.SendResponse(c)
}
