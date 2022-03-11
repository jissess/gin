package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range ctx.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("Access-Control-Allow-Origin,Access-Control-Allow-Headers,%s", headerStr)
		} else {
			headerStr = "Access-Control-Allow-Origin,Access-Control-Allow-Headers"
		}
		if origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, XMLHttpRequest, content-type, x-requested-with, Content-Length, X-CSRF-Token, Token, session")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			ctx.Header("Access-Control-Max-Age", "172800")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Set("Content-Type", "application/json; charset=utf-8")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		ctx.Next()
	}
}
