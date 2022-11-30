package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func CommonMidReq() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		/** 设置此次请求的参数 */
		context.Set("start",startTime)
		/** 执行该请求剩下的操作 */
		context.Next()

		useTime := time.Since(startTime)
		log.Println("req useTime:",useTime)
	}
}

func AccountMidReq() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}