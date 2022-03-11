package middleware

import (
	"github.com/gin-gonic/gin"
	"myGin/common"
	"myGin/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未登录"})
			ctx.Abort()
			return
		}
		parts := strings.SplitN(tokenString, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "Authorization格式有误"})
			ctx.Abort()
			return
		}
		token, claims, err := common.ParseToken(parts[1])
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token有误"})
			ctx.Abort()
			return
		}
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
