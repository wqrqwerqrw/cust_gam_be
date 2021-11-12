package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
)

// CreateUser 前端传json为model.APIUser
func CreateUser(c *gin.Context) {
	req := model.APIUser{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err := store.CreateUser(req.UserName, req.Phone, req.Password, req.IsVip)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}

// ChangeUser 前端传json为model.APIUserWithId
func ChangeUser(c *gin.Context) {
	req := model.APIUserWithId{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err := store.UpdateUser(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, nil)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err = store.DeleteUser(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, nil)
}

// QueryUser 查询用户详情
func QueryUser(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	dbUser, err := store.QueryUser(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbUser)
}
