package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, params ...interface{}) {
	json := gin.H{
		"code": 1000,
	}
	for _, value := range params {
		switch value.(type) {
		case int:
			json["code"] = value
		case string:
			json["msg"] = value
		default:
			json["data"] = value
		}
	}
	c.JSON(http.StatusOK, json)
}
