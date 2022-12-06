package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
	"mitte-task/models"
	db "mitte-task/models/db"
	"mitte-task/services"
	"net/http"
)

// CreateProfile godoc
// @Summary      CreateProfile
// @Description  Create profile for a user
// @Tags         create profile
// @Accept       json
// @Produce      json
// @Param        req  body      models.ProfileRequest true "Profile Request"
// @Success      200  {object}  models.Response
// @Router       /profile/create [post]
func CreateProfile(c *gin.Context) {
	var req models.ProfileRequest
	_ = c.ShouldBindBodyWith(&req, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// create user record
	user, err := services.CreateProfile(req.UserId, req.Age, req.AgeFilterMin, req.AgeFilterMax, req.InterestedIn, req.Gender, req.PreferredLocation, req.Swipes)
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

// Profiles godoc
// @Summary      Profiles
// @Description List down the match profiles
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /profile/:user_id [get]
func Profiles(c *gin.Context) {
	userid := c.Request.URL.Query().Get("user_id")

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// create user record
	matchedProfiles, err := services.MatchProfiles(userid)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"match_profiles": matchedProfiles,
	}
	response.SendResponse(c)
}

// FilteredProfiles godoc
// @Summary      FilteredProfiles
// @Description Filter profile by age or gender
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Response
// @Router       /profile/filter/:age/:gender [get]
func FilteredProfiles(c *gin.Context) {
	age := c.Request.URL.Query().Get("age")
	gender := c.Request.URL.Query().Get("gender")

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	var err error
	var filteredProfiles []db.Profile
	// fetch filtered profiles
	filteredProfiles, err = services.FilteredProfiles(cast.ToInt64(age), gender)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"filtered_profiles": filteredProfiles,
	}
	response.SendResponse(c)
}
