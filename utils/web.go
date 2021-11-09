package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrorInternalError = 1
	ErrorWrongAttr     = 2
)

func MakeErrorResp(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": message,
		"data":    nil,
	})
}

func MakeOKResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "success",
		"data":    data,
	})
}
