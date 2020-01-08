package controllers

import (
	"apichanged/controllerhandling"
	"apichanged/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Retrieving all the posts details
func GetPosts(c *gin.Context) {
	var post []models.Posts
	err := models.GetAllPosts(&post)
	controllerhandling.ReturnObject(c, err, post)

}

//Retrieve a post by id
func GetPostDetail(c *gin.Context) {
	id := c.Params.ByName("id")

	var post models.Posts
	err := models.GetAPost(&post, id)

	controllerhandling.ReturnObject(c, err, post)

}

//Retrive all posts of specific user by userid
func GetPostsDetailByUserID(c *gin.Context) {
	id := c.Params.ByName("userid")

	var post []models.Posts

	err := models.GetPostByUserId(&post, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	controllerhandling.ReturnObject(c, err, post)

}

//inserting a post

func CreateaPost(c *gin.Context) {

	var post models.Posts
	c.BindJSON(&post)

	if post.Title != "" && post.Body != "" {

		err := models.CreateAPost(&post)
		controllerhandling.ReturnObject(c, err, post)
	} else {
		c.AbortWithStatus(http.StatusPreconditionFailed)
	}

}

//deleting a post
func DeletePost(c *gin.Context) {

	id := c.Params.ByName("id")

	var post models.Posts

	err := models.DeleteAPost(&post, id)

	controllerhandling.ReturnMessage(c, err, post)

}

// Updating a post details
func UpdatePost(c *gin.Context) {
	id := c.Params.ByName("id")

	var post models.Posts

	err := models.GetAPost(&post, id)

	if err == nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&post)
	err = models.UpdateAPost(&post, id)
	controllerhandling.ReturnObject(c, err, post)

}
