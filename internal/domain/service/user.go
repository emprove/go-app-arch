package service

import (
	"go-app-arch/internal/domain/entity"
	"go-app-arch/internal/domain/repository"
)

type UserServiceInterface interface {
	FindOneByToken(token string) (*entity.User, error)
}

type userService struct {
	RepoUser repository.User
}

func NewUserService(r repository.User) UserServiceInterface {
	return &userService{RepoUser: r}
}

func (s *userService) FindOneByToken(token string) (*entity.User, error) {
	user, err := s.RepoUser.FindOneByToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
