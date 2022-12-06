package router

import (
	"demo/controller"
	accountController "demo/controller/account"
	middleware2 "demo/service/middleware"
	"github.com/gin-gonic/gin"
)
func LoadRouter() *gin.Engine {
	r := gin.Default()
	/** 全局中间件 */
	r.Use(middleware2.CommonMidReq())
	/** 用户相关路由 */
	accountGroup := r.Group("/account", middleware2.AccountMidReq())
	{
		accountGroup.POST("/getAccountInfo",accountController.GetAccountInfo)

	}


	/** 测试接口 */
	r.GET("/test1",controller.Test)
	r.GET("/test2",controller.TestB)
	r.GET("/test3",controller.TestC)
	r.GET("/test4",controller.TestD)
	return r
}