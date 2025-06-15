package service

import (
	"vijju/user/model"
	"vijju/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*model.User, error) {
	user := &model.User{Name: name, Email: email}
	err := s.repo.Create(user)
	return user, err
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(id uint, name, email string) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	err = s.repo.Update(user)
	return user, err
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
