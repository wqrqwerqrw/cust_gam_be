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
	dbServices := &[]model.APIUserServiceWithId{}
	for _, service := range *dbService {
		ser := model.APIUserServiceWithId{
			Id:          service.ID,
			UserName:    service.UserName,
			ServiceName: service.ServiceName,
			ServiceCost: service.ServiceCost,
			ServiceTime: service.ServiceTime,
			IsUsed:      service.IsUsed,
		}
		*dbServices = append(*dbServices, ser)
	}

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, dbServices)
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

	service := model.APIServiceWithId{
		Id:   dbService.ID,
		Name: dbService.Name,
		Cost: dbService.Cost,
		Desc: dbService.Desc,
		Tag:  dbService.Tag,
	}
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, service)
}

// QueryAllService 查询所有服务
func QueryAllService(c *gin.Context) {
	dbService, err := store.QueryAllService()

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	dbServices := &[]model.APIServiceWithId{}
	for _, service := range *dbService {
		ser := model.APIServiceWithId{
			Id:   service.ID,
			Tag:  service.Tag,
			Name: service.Name,
			Cost: service.Cost,
			Desc: service.Desc,
		}
		*dbServices = append(*dbServices, ser)
	}

	utils.MakeOKResp(c, dbServices)
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

// UpDateService 修改服务信息
func UpDateService(c *gin.Context) {
	req := model.APIServiceWithId{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}
	err := store.UpdateService(req)
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}

	utils.MakeOKResp(c, nil)
}

// DeleteUserService 删除用户服务
func DeleteUserService(c *gin.Context) {
	req, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err = store.DeleteUserService(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, nil)
}

// QueryTotalMoney 查询用户服务总金额
func QueryTotalMoney(c *gin.Context) {
	req := c.Query("username")

	totalMoney, err := store.QueryTotalMoney(req)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, totalMoney)
}

// PayMoney 支付现金
func PayMoney(c *gin.Context) {
	req := model.APIPayMoney{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
		return
	}

	err := store.PayMoney(req)
	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
		return
	}
	utils.MakeOKResp(c, nil)
}
