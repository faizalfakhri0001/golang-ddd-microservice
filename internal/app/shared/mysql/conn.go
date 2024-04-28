package mysql

import (
	"fmt"
	"golang-gin-ddd/config"
	"golang-gin-ddd/internal/app/shared/interfaces"
	"golang-gin-ddd/internal/app/shared/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	db *gorm.DB
}

func (db *MySQLDatabase) GetIntance() *gorm.DB {
	return db.db
}

func NewDatabase() interfaces.IDatabase {
	dbConfig := config.Config.Database

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.PanicIfError(err)

	return &MySQLDatabase{
		db: db,
	}
}
