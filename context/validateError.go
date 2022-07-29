package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*ValidateError : 用户登录*/
func ValidateError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  "Validate error",
	})
}

/* NeedLogin : 用户未登录 */
func NeedLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "用户未登录",
	})
}
