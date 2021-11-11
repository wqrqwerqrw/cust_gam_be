package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wqrqwerqrw/cust_gam_backend/handler"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
)

func main() {
	store.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userGroup := r.Group("/user")
	userGroup.POST("/add", handler.CreateUser)
	userGroup.POST("/change", handler.ChangeUser)
	userGroup.DELETE("/delete", handler.DeleteUser)
	userGroup.GET("/query", handler.QueryUser)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
