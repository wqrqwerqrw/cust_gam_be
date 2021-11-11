package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
)

// QueryServiceByUserName 查询用户的服务情况
func QueryServiceByUserName(c *gin.Context) {
	req := c.Query("username")

	dbUser, err := store.QueryByUserName(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbUser)
}

// CreateService 添加服务
func CreateService(c *gin.Context) {

}

// UpDateService 修改服务信息
func UpDateService(c *gin.Context) {

}

// DeleteService 删除服务
func DeleteService(c *gin.Context) {

}
