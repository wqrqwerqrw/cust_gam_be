package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
	"strconv"
)

// QueryServiceByUserName 查询用户的服务情况
func QueryServiceByUserName(c *gin.Context) {
	req := c.Query("username")

	user, err := store.QueryUserByUserName(req)

	if err != nil || user == nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "用户不存在")
		return
	}

	dbService, err := store.QueryServiceByUserName(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbService)
}

// CreateService 添加服务
func CreateService(c *gin.Context) {
	req := model.APIService{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}
	err := store.CreateService(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}

// UpDateUserService 修改用户服务信息
func UpDateUserService(c *gin.Context) {
	req := model.APIUserServiceWithId{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	user, err := store.QueryUserByUserName(req.UserName)
	if err != nil || user == nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "用户不存在")
		return
	}

	service, err := store.QueryServiceByName(req.ServiceName)
	if err != nil || service == nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "服务项目不存在")
		return
	}

	err = store.UpdateUserService(req)
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}

// DeleteService 删除服务
func DeleteService(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err = store.DeleteService(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, nil)
}

// QueryService 查询服务
func QueryService(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	dbService, err := store.QueryService(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbService)
}

// AddServiceWithUserName 给用户添加服务信息
func AddServiceWithUserName(c *gin.Context) {
	req := model.APIUserService{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	user, err := store.QueryUserByUserName(req.UserName)
	if err != nil || user == nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "用户不存在")
		return
	}

	service, err := store.QueryServiceByName(req.ServiceName)
	if err != nil || service == nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "服务项目不存在")
		return
	}

	err = store.AddServiceWithUserName(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}
