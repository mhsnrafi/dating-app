package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mitte-task/models"
	"mitte-task/services"
	"net/http"
)

// CreateUser godoc
// @Summary      CreateUser
// @Description  Create random user
// @Tags         create user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /user/create [post]
func CreateUser(c *gin.Context) {
	var requestBody models.RegisterRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// create user record
	user, err := services.CreateUser()
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{
		"user": user,
	}
	response.SendResponse(c)
}
