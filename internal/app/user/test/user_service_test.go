package test

import (
	"errors"
	"golang-gin-ddd/internal/app/user/domain"
	"golang-gin-ddd/internal/app/user/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserGetById_NotFound(t *testing.T) {
	userRepository := UserRepositoryMock{
		Mock: &mock.Mock{},
	}
	// Set up the mock to expect GetById("1") and return nil, indicating user not found
	userRepository.Mock.On("GetById", "1").Return(nil, errors.New("user not found"))

	var userService = infrastructure.NewUserService(userRepository)
	// Call the GetById method with argument "1"
	user, err := userService.GetById("1")

	// Assert that an error is returned and user is nil
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserGetById(t *testing.T) {
	userRepository := UserRepositoryMock{
		Mock: &mock.Mock{},
	}
	// Set up the mock to expect GetById("1") and return nil, indicating user not found
	user := &domain.User{
		ID:        1,
		FirstName: "Akbar",
		LastName:  "Jumali",
		Address:   "Tangerang Selatan",
		City:      "Tangerang",
	}
	userRepository.Mock.On("GetById", "1").Return(*user, nil)

	var userService = infrastructure.NewUserService(userRepository)
	// Call the GetById method with argument "1"
	result, err := userService.GetById("1")

	// Assert that an error is returned and user is nil
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
}
