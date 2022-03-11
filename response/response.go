package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 0, data, msg)
}

func Failed(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, 1, data, msg)
}
