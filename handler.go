package main

import (
	"buildgram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func successResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status": "success",
		"data":   data,
	})
}

func errorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}

func createUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "username and email are required")
		return
	}

	var bio *string
	if req.Bio != "" {
		bio = &req.Bio
	}

	user := models.User{
		ID:       nextUserID,
		Username: req.Username,
		Email:    req.Email,
		Bio:      bio,
	}

	nextUserID++
	users = append(users, user)

	successResponse(c, http.StatusCreated, user)
}

func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	for _, user := range users {
		if user.ID == id {
			successResponse(c, http.StatusOK, user)
			return
		}
	}

	errorResponse(c, http.StatusNotFound, "user not found")
}

func createPost(c *gin.Context) {
	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "userID, imageURL and caption are required")
		return
	}

	userExists := false
	for _, user := range users {
		if user.ID == req.UserID {
			userExists = true
			break
		}
	}

	if !userExists {
		errorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	post := models.Post{
		ID:         nextPostID,
		UserID:     req.UserID,
		ImageURL:   req.ImageURL,
		Caption:    req.Caption,
		Timestamp:  time.Now(),
		LikesCount: 0,
	}

	nextPostID++
	posts = append(posts, post)

	successResponse(c, http.StatusCreated, post)
}

func getAllPosts(c *gin.Context) {
	successResponse(c, http.StatusOK, posts)
}

func getPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	for _, post := range posts {
		if post.ID == id {

			var postComments []models.Comment

			for _, comment := range comments {
				if comment.PostID == id {
					postComments = append(postComments, comment)
				}
			}

			successResponse(c, http.StatusOK, gin.H{
				"post":     post,
				"comments": postComments,
			})
			return
		}
	}

	errorResponse(c, http.StatusNotFound, "post not found")
}

func likePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	for i := range posts {
		if posts[i].ID == id {
			posts[i].LikesCount++

			successResponse(c, http.StatusOK, gin.H{
				"id":         posts[i].ID,
				"likesCount": posts[i].LikesCount,
			})
			return
		}
	}

	errorResponse(c, http.StatusNotFound, "post not found")
}

func addComment(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var req models.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "userID and text are required")
		return
	}

	postExists := false
	for _, post := range posts {
		if post.ID == postID {
			postExists = true
			break
		}
	}

	if !postExists {
		errorResponse(c, http.StatusNotFound, "post not found")
		return
	}

	comment := models.Comment{
		ID:        nextCommentID,
		PostID:    postID,
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now(),
	}

	nextCommentID++
	comments = append(comments, comment)

	successResponse(c, http.StatusCreated, comment)
}