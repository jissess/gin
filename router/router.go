package router

import (
	"github.com/gin-gonic/gin"
	"myGin/controller"
	"myGin/middleware"
	"myGin/response"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.GET("/", func(ctx *gin.Context) {
		response.Success(ctx, nil, "成功")
	})
	apiRouters := r.Group("/api")
	{
		userRouters := apiRouters.Group("/user")
		{
			userRouters.POST("/register", controller.UserController{}.Register)
		}
	}
}
