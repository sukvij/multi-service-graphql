package service

import (
	"vijju/post/model"
	"vijju/post/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(title, content string, userID uint) (*model.Post, error) {
	post := &model.Post{Title: title, Content: content, UserID: userID}
	err := s.repo.Create(post)
	return post, err
}

func (s *PostService) GetPost(id uint) (*model.Post, error) {
	return s.repo.FindByID(id)
}

func (s *PostService) GetPosts() ([]model.Post, error) {
	return s.repo.FindAll()
}

func (s *PostService) UpdatePost(id uint, title, content string) (*model.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Content = content
	err = s.repo.Update(post)
	return post, err
}

func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
