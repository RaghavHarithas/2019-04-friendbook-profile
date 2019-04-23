package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	dbModel "github.com/rajch/2019-04-friendbook-profile/database/models"
	"github.com/rajch/2019-04-friendbook-profile/models"
)

func GetProfile(c *gin.Context) {
	userEmail := c.Query("userEmail")
	profile, err := dbModel.GetProfileByEmail(userEmail)
	if err != nil {
		log.Println("Error while getting profile: %s", err)
		c.Status(http.StatusNotFound)
		return
	}

	//	respProfile := models.Profile{
	//		City:         *profile.City,
	//		DateOfBirth:  *profile.DateOfBirth,
	//		FirstName:    *profile.FirstName,
	//		LastName:     *profile.LastName,
	//		Gender:       *profile.Gender,
	//		MobileNumber: *profile.MobileNumber,
	//		UserName:     *profile.UserName,
	//	}

	c.JSON(http.StatusOK, profile)
	return
}

func CreateOrUpdateProfile(c *gin.Context) {
	userEmail := c.Query("userEmail")
	requestBody := models.Profile{}
	c.Bind(&requestBody)
	profile := dbModel.Profile{
		Email: userEmail,
		City:  &requestBody.City,
	}

	err := dbModel.CreateOrUpdateProfile(&profile)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &requestBody)
}
