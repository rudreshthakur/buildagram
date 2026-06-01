package main

import "buildgram/models"

var users []models.User
var posts []models.Post
var comments []models.Comment

var nextUserID = 1
var nextPostID = 1
var nextCommentID = 1