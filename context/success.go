package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, params ...interface{}) {
	json := gin.H{
		"code": 200,
	}
	for _, value := range params {
		switch value.(type) {
		case string:
			json["msg"] = value
		default:
			json["data"] = value
		}
	}
	c.JSON(http.StatusOK, json)
}
