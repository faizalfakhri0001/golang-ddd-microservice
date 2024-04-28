package model

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Address   string
	City      string
}
