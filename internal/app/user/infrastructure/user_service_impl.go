package infrastructure

import "golang-gin-ddd/internal/app/user/domain"

type UserServiceImpl struct {
	userRepository domain.UserRepository
}

func NewUserService(
	userRepository domain.UserRepository,
) domain.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (us *UserServiceImpl) Create() error {
	return nil
}

// GetById implements domain.UserService.
func (us *UserServiceImpl) GetById(id string) (*domain.User, error) {
	user, err := us.userRepository.GetById(id)
	return user, err
}
