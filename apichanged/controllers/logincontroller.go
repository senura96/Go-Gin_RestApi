package controllers

import (
	"apichanged/models"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {

	var user models.Users

	username := c.Params.ByName("username")
	password := c.Params.ByName("password")

	err := models.FindUserByUNPW(&user, username, password)

	if err == nil {

		c.JSON(200, user)

	} else {

		c.JSON(404, gin.H{" error": "user not  found"})

	}

}
