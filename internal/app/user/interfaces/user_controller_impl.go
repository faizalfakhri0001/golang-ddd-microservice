package interfaces

import (
	"golang-gin-ddd/internal/app/shared/utils"
	"golang-gin-ddd/internal/app/user/domain"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userService domain.UserService
}

func NewUserController(
	userService domain.UserService,
) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (ctrl *UserControllerImpl) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(201, gin.H{
			"code":   201,
			"status": "success create",
		})
	}
}

func (ctrl *UserControllerImpl) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idstr := ctx.Param("id")

		user, err := ctrl.userService.GetById(idstr)
		utils.PanicIfError(err)

		ctx.JSON(200, gin.H{
			"code":   200,
			"data":   user,
			"status": "success",
		})
	}
}
