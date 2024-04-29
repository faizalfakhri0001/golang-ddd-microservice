package middleware

import (
	gohttp "golang-gin-ddd/pkg/http"
	"golang-gin-ddd/pkg/http/wrapper"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				exception, ok := r.(error)
				if ok {
					logrus.Error(exception)
					wrapper.Translate(ctx, gohttp.Response{
						Error: exception,
					})
				}
			}
		}()
		ctx.Next()
	}
}
