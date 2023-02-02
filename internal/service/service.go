package service

import (
	"forum/models"
)

type Authorization interface {
	CreateUser(models.User) error
	GenerateToken(string, string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	DeleteToken(string) error
}

type Post interface {
	CreatePost(*models.Post) error
	GetAllPost() (**[]models.Post, error)
	GetPostByCategory(string) (**[]models.Post, error)
	MyPosts(string) (*[]models.Post, error)
	MyFavourites(int) (*[]models.Post, error)
	GetPostByID(string) (*models.Post, error)
}

type Comment interface {
	CreateComment(models.Comment) error
	GetCommentByPostID(int) (*[]models.Comment, error)
}

type Service struct {
	Authorization
	Post
	Comment
}
