package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
)

// CreateUser 前端传json
/*
{
    "user_name" : "sss",
    "phone" : "123123123",
    "password" : "123123123",
    "isVip" : 1
}
*/
func CreateUser(c *gin.Context) {
	req := model.APIUser{}
	if err := c.ShouldBind(&req); err != nil {
		utils.MakeErrorResp(c, utils.ErrorWrongAttr, "参数错误")
	}

	dbUser, err := store.CreateUser(req.UserName, req.Phone, req.Password, req.IsVip)

	if err != nil {
		utils.MakeErrorResp(c, utils.ErrorInternalError, "内部错误")
	}

	utils.MakeOKResp(c, dbUser)
}
