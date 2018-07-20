package main

import (
	"net/http"
	"web_demo/users"
	"gopkg.in/gin-gonic/gin.v1"
)

func subServ()  {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}

func main() {
	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	router.Static("/static", "./public")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World the gin")
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1/userinfo")
	{
		v1.POST("/", users.CreateUser)
		v1.GET("/", users.FetchAllUsers)
		v1.GET("/:id", users.FetchSingleUser)
		v1.PUT("/:id", users.UpdateUser)
		v1.DELETE("/:id", users.DeleteUser)
	}

	router.Run(":8000")
}