package test

import (
	"errors"
	"golang-gin-ddd/internal/app/user/domain"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock *mock.Mock
}

// Create implements domain.UserRepository.
func (repository UserRepositoryMock) Create() error {
	panic("unimplemented")
}

// GetById implements domain.UserRepository.
func (repository UserRepositoryMock) GetById(id string) (*domain.User, error) {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, errors.New("user not found")
	} else {
		user := args.Get(0).(domain.User)
		return &user, nil
	}
}
