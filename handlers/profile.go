package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	dbModel "github.com/rajch/2019-04-friendbook-profile/database/models"
)

func GetProfile(c *gin.Context) {
	userEmail := c.Query("userEmail")
	profile, err := dbModel.GetProfileByEmail(userEmail)
	if err != nil {
		log.Println("Error while getting profile: %s", err)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, profile)
	return
}

func CreateOrUpdateProfile(c *gin.Context) {
	userEmail := c.Query("userEmail")
	requestBody := dbModel.Profile{}
	c.BindJSON(&requestBody)
	log.Printf("%s", *requestBody.UserName)
	profile := dbModel.Profile{
		Email:        userEmail,
		City:         requestBody.City,
		DateOfBirth:  requestBody.DateOfBirth,
		Gender:       requestBody.Gender,
		LastName:     requestBody.LastName,
		MobileNumber: requestBody.MobileNumber,
		UserName:     requestBody.UserName,
		FirstName:    requestBody.FirstName,
	}

	err := dbModel.CreateOrUpdateProfile(&profile)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &profile)
}
