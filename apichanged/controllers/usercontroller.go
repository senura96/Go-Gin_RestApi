package controllers

import (
	"apichanged/controllerhandling"
	"apichanged/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Retrieving all the user details
func GetUsers(c *gin.Context) {
	var user []models.Users
	err := models.GetAllUserDetails(&user)

	controllerhandling.ReturnObject(c, err, user)

}

//Retrieve a user by id
func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.Users
	err := models.GetAUser(&user, id)
	controllerhandling.ReturnObject(c, err, user)
}

//Inserting a user detail
func PostUser(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	log.Println(user)
	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
		err := models.CreateUser(&user)

		controllerhandling.ReturnObject(c, err, user)

	} else {

		c.AbortWithStatus(http.StatusPreconditionFailed)

	}

}

//Deleting a user
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	var user models.Users

	err := models.DeleteAUser(&user, id)

	controllerhandling.ReturnMessage(c, err, user)

}

// Updating A User

func Updateuser(c *gin.Context) {

	var user models.Users

	id := c.Params.ByName("id")

	err := models.GetAUser(&user, id)

	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&user)

	err = models.UpdateAUser(&user, id)
	controllerhandling.ReturnObject(c, err, user)

}
