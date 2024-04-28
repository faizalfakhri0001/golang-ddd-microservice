package interfaces

import "github.com/gin-gonic/gin"

type UserController interface {
	Create() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	// tambahkan metode lain yang diperlukan
}
