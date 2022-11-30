package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccountInfo(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"status":200,
		"msg":"success",
	})
}


/** 授权 */
func auth()  {
	
}


/** 解除授权 **/
func unAuth()  {
	
}


/** 获取账套 **/
func getGroupList() {

}