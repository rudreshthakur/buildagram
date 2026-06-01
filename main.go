package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Use(LoggerMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", createUser)
		v1.GET("/users/:id", getUserByID)

		v1.POST("/posts", createPost)
		v1.GET("/posts", getAllPosts)
		v1.GET("/posts/:id", getPostByID)

		v1.POST("/posts/:id/like", likePost)
		v1.POST("/posts/:id/comments", addComment)
	}

	router.Run(":8080")
}