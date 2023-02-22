package transactions

import (
	"context"
	"errors"
)

type Repository interface {
	GetRoleByName(name string) (*Role, error)
	SaveUser(user User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateUser(ctx context.Context, request UserRequest) (*UserResponse, error) {
	var user User = *request.toUser()
	_, err := s.Repository.GetUserByEmail(request.Email)
	if err == nil {
		return nil, errors.New("email is already being used")
	}

	roleUser, err := s.Repository.GetRoleByName("USER")
	if err != nil {
		return nil, errors.New("error on load user role")
	}

	user.Roles = []Role{*roleUser}

	savedUser, err := s.Repository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return savedUser.toResponse(), nil
}
