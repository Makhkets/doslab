package repository

import (
	"doslab/internal/service/domain"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

type postRepository struct {
}

func (p postRepository) Posts() ([]domain.Post, error) {
	posts := make([]domain.Post, 256)

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p postRepository) Comments(postId int) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 256)

	response, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d/comments", postId))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func NewPostRepository() domain.PostRepository {
	return &postRepository{}
}
