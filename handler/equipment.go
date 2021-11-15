package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
	"strconv"
)

// CreateEquipment 前端传json为model.APIEquipment
func CreateEquipment(c *gin.Context) {
	req := model.APIEquipment{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}
	err := store.CreatEquipment(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}

// QueryEquipment 查询设备详情
func QueryEquipment(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	dbEquip, err := store.QueryEquipment(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbEquip)
}

// QueryAllEquipment 查询所有设备详情
func QueryAllEquipment(c *gin.Context) {
	dbEquip, err := store.QueryAllEquipment()

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbEquip)
}
