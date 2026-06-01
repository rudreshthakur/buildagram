package models

import "time"

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio,omitempty"`
}

type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userID"`
	ImageURL   string    `json:"imageURL"`
	Caption    string    `json:"caption"`
	Timestamp  time.Time `json:"timestamp"`
	LikesCount int       `json:"likesCount"`
}

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postID"`
	UserID    int       `json:"userID"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Bio      string `json:"bio"`
}

type CreatePostRequest struct {
	UserID   int    `json:"userID" binding:"required"`
	ImageURL string `json:"imageURL" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
}

type CreateCommentRequest struct {
	UserID int    `json:"userID" binding:"required"`
	Text   string `json:"text" binding:"required"`
}