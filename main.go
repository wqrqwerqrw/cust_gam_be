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

	serviceGroup := r.Group("/service")
	serviceGroup.POST("/add", handler.CreateService)
	serviceGroup.DELETE("/delete", handler.DeleteService)
	serviceGroup.GET("/query", handler.QueryService)

	userService := r.Group("/userService")
	userService.POST("/add", handler.AddServiceWithUserName)
	userService.POST("/change", handler.UpDateUserService)
	userService.GET("/query", handler.QueryServiceByUserName)
	r.Run("localhost:9090") // 监听并在 0.0.0.0:8080 上启动服务
}
