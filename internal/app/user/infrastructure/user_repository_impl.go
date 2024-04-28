package infrastructure

import (
	"golang-gin-ddd/internal/app/user/domain"

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
	result := rp.db.Find(&user).Where("id=?", id)

	return &user, result.Error
}
