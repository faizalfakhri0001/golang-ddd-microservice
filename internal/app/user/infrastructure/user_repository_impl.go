package infrastructure

import (
	"golang-gin-ddd/internal/app/user/domain"
	"golang-gin-ddd/pkg/errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (rp *UserRepositoryImpl) Create() error {
	return nil
}

func (rp *UserRepositoryImpl) GetById(id string) (*domain.User, error) {
	var user domain.User
	if err := rp.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
	}

	return &user, nil
}
