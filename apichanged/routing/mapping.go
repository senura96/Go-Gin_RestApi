package routing

import (
	"apichanged/controllers"

	"github.com/gin-gonic/gin"
)

func RouterMappings() *gin.Engine {

	Router := gin.Default()

	v1 := Router.Group("api/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUserDetail)
		v1.POST("/users", controllers.PostUser)
		v1.PUT("/users", controllers.Updateuser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}

	v2 := Router.Group("api/v2")
	{

		v2.GET("/posts", controllers.GetPosts)
		v2.GET("/posts/:id", controllers.GetPostDetail)
		v2.POST("/posts/:id", controllers.GetPostsDetailByUserID)
		v2.POST("/posts", controllers.CreateaPost)
		v2.PUT("/posts", controllers.UpdatePost)
		v2.DELETE("/posts/:id", controllers.DeletePost)

	}

	v3 := Router.Group("api/v3")
	{

		v3.GET("/login", controllers.LoginController)

	}
	return Router

}
