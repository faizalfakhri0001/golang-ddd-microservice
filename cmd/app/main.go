package main

import (
	"golang-gin-ddd/internal/app/shared/mysql"
	userInterface "golang-gin-ddd/internal/app/user/interfaces"
	"golang-gin-ddd/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.NewDatabase()

	router := gin.Default()
	api := router.Group("/api")
	api.Use(middleware.Recovery())

	v1 := api.Group("/v1")
	userInterface.UserRoute(v1, db)

	_ = router.Run(":3028")
}
