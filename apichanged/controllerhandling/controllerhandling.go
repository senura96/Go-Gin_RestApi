package controllerhandling

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  Handling Functions which Returning json object or status
func ReturnObject(c *gin.Context, err error, obj interface{}) {

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, obj)
	}

}

// Handling Functions which Returning  message or stauts
func ReturnMessage(c *gin.Context, err error, obj interface{}) {

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Post not Found"})
	}

}
