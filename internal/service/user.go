package service

import (
	"go-app-arch/internal/entity"
	"go-app-arch/internal/repository"
)

type User struct {
	RepoUser repository.User
}

func NewUserService(r repository.User) *User {
	return &User{RepoUser: r}
}

func (s *User) FindOneByToken(token string) (*entity.User, error) {
	user, err := s.RepoUser.FindOneByToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
