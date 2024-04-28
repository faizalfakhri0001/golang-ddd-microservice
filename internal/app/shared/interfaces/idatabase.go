package interfaces

import "gorm.io/gorm"

type IDatabase interface {
	GetIntance() *gorm.DB
}
