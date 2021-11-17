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
	userGroup.GET("/queryAll", handler.QueryAllUser)
	userGroup.GET("/queryByUsername", handler.QueryUserByUsername)

	serviceGroup := r.Group("/service")
	serviceGroup.POST("/add", handler.CreateService)
	serviceGroup.DELETE("/delete", handler.DeleteService)
	serviceGroup.GET("/query", handler.QueryService)
	serviceGroup.GET("/queryAll", handler.QueryAllService)
	serviceGroup.POST("/change", handler.UpDateService)

	userServiceGroup := r.Group("/userService")
	userServiceGroup.POST("/add", handler.AddServiceWithUserName)
	userServiceGroup.POST("/change", handler.UpDateUserService)
	userServiceGroup.GET("/query", handler.QueryServiceByUserName)
	userServiceGroup.DELETE("/delete", handler.DeleteUserService)
	userServiceGroup.GET("/queryMoney", handler.QueryTotalMoney)
	userServiceGroup.POST("/payMoney", handler.PayMoney)

	equipmentGroup := r.Group("/equipment")
	equipmentGroup.POST("/add", handler.CreateEquipment)
	equipmentGroup.GET("/query", handler.QueryEquipment)
	equipmentGroup.GET("/queryAll", handler.QueryAllEquipment)
	equipmentGroup.DELETE("/delete", handler.DeleteEquipment)
	equipmentGroup.POST("/change", handler.ChangeEquipment)

	r.Run("localhost:9090") // 监听并在 0.0.0.0:8080 上启动服务
}
