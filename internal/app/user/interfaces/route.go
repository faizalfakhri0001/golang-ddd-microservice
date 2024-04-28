package interfaces

import (
	db "golang-gin-ddd/internal/app/shared/interfaces"
	"golang-gin-ddd/internal/app/user/infrastructure"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, db db.IDatabase) {
	repository := infrastructure.NewUserRepository(db.GetIntance())
	service := infrastructure.NewUserService(repository)
	controller := NewUserController(service)
	router.POST("users/", controller.Create())
	router.GET("users/:id", controller.GetByID())
}
