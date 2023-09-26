package usecase

import "doslab/internal/service/domain"

type postUseCase struct {
	postRepository domain.PostRepository
}

func (p postUseCase) Posts() ([]domain.Post, error) {
	return p.postRepository.Posts()
}

func NewPostUseCase(authRep domain.PostRepository) domain.PostUseCase {
	return &postUseCase{authRep}
}

func (p postUseCase) Comments(postId int) ([]domain.Comment, error) {
	return p.postRepository.Comments(postId)
}
