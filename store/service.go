package store

import (
	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"time"
)

// QueryServiceByUserName 查询用户的服务
func QueryServiceByUserName(username string) (*[]model.TblUserService, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	services := &[]model.TblUserService{}
	err = db.Debug().Table("tbl_user_service").Where("user_name = ? and deleted = 0", username).Find(&services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

// DeleteService 删除服务
func DeleteService(id int) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	return db.Debug().Table("tbl_service").Where("id = ? and deleted = 0", id).Update("deleted", 1).Error
}

// QueryService 查询服务
func QueryService(id int) (*model.TblService, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	service := &model.TblService{}

	err = db.Debug().Table("tbl_service").Where("id = ? and deleted = 0", id).First(service).Error

	if err != nil {
		return nil, err
	}

	return service, nil
}

// AddServiceWithUserName 给某个用户添加服务
func AddServiceWithUserName(userService model.APIUserService) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	now := time.Now()
	service := model.TblUserService{
		UserName:    userService.UserName,
		ServiceName: userService.ServiceName,
		ServiceCost: userService.ServiceCost,
		ServiceTime: userService.ServiceTime,
		IsUsed:      userService.IsUsed,
		CreateTime:  now,
		UpdateTime:  now,
		Deleted:     0,
	}
	err = db.Debug().Create(&service).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateService 创建服务
func CreateService(ser model.APIService) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	now := time.Now()

	service := model.TblService{
		Name:       ser.Name,
		Cost:       ser.Cost,
		Desc:       ser.Desc,
		CreateTime: now,
		UpdateTime: now,
	}
	err = db.Debug().Table("tbl_service").Create(&service).Error

	if err != nil {
		return err
	}

	return nil
}
