package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myGin/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Failed(ctx, fmt.Sprint(err), nil)
			}
		}()
		ctx.Next()
	}
}
