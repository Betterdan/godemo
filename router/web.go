package router

import (
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
	return r
}