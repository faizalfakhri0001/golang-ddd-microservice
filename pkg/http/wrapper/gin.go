package wrapper

import (
	"golang-gin-ddd/pkg/errors"
	gohttp "golang-gin-ddd/pkg/http"

	"github.com/gin-gonic/gin"
)

// constants wrapper http
const (
	DataField    = "data"
	TokenField   = "token"
	TraceIDField = "trace_id"
	Error        = "errors"
	StatusField  = "status"
	CodeField    = "code"
	MessageField = "message"
)

// GinHandlerFn gin handler function
type GinHandlerFn func(c *gin.Context) gohttp.Response

func Wrap(fn GinHandlerFn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := fn(ctx)
		Translate(ctx, res)
	}
}

func Translate(ctx *gin.Context, res gohttp.Response) {
	var resCode int
	result := gin.H{}
	if _, ok := res.Error.(errors.CustomError); ok {
		code := int(errors.GetType(res.Error))

		if code > 1000 {
			resCode = 500
			result[CodeField] = resCode
			result[StatusField] = errors.GetStatus(resCode)
			result[MessageField] = res.Error.Error()
		} else if code == 500 {
			resCode = 500
			result[CodeField] = code
			result[StatusField] = errors.GetStatus(resCode)
			result[MessageField] = res.Error.Error()
		} else {
			resCode = code
			result[CodeField] = code
			result[StatusField] = errors.GetStatus(code)
			result[MessageField] = errors.GetMsg(code)
		}
	}

	parsedError, exist := ctx.Get("error")
	if exist {
		result[Error] = parsedError
		delete(result, MessageField)
	}

	// get data
	if res.Data != nil {
		result[DataField] = res.Data
	}

	if res.Token != nil {
		result[TokenField] = res.Token
	}

	ctx.JSON(resCode, result)
}
