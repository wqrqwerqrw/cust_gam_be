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

// DeleteUserService 删除用户服务
func DeleteUserService(id int) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	return db.Debug().Table("tbl_user_service").Where("id = ? and deleted = 0", id).Update("deleted", 1).Error
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

// QueryAllService 查询所有服务
func QueryAllService() (*[]model.TblService, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	services := &[]model.TblService{}

	err = db.Debug().Table("tbl_service").Where("deleted = 0").Find(services).Error

	if err != nil {
		return nil, err
	}

	return services, nil
}

// QueryServiceByName 查询服务
func QueryServiceByName(name string) (*model.TblService, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	service := &model.TblService{}

	err = db.Debug().Table("tbl_service").Where("name = ? and deleted = 0", name).First(service).Error

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
		Tag:        ser.Tag,
		CreateTime: now,
		UpdateTime: now,
	}
	err = db.Debug().Table("tbl_service").Create(&service).Error

	if err != nil {
		return err
	}

	return nil
}

// UpdateUserService 更新用户与服务
func UpdateUserService(ser model.APIUserServiceWithId) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	attr := map[string]interface{}{
		"user_name":    ser.UserName,
		"service_name": ser.ServiceName,
		"service_cost": ser.ServiceCost,
		"service_time": ser.ServiceTime,
		"is_used":      ser.IsUsed,
	}

	err = db.Debug().Table("tbl_user_service").Where("id = ? and deleted = 0", ser.Id).
		Updates(attr).
		Error

	if err != nil {
		return err
	}
	return nil
}

// UpdateService 更新服务
func UpdateService(ser model.APIServiceWithId) error {
	db, err := DBConn()
	if err != nil {
		return err
	}

	attr := map[string]interface{}{
		"name": ser.Name,
		"cost": ser.Cost,
		"tag":  ser.Tag,
		"desc": ser.Desc,
	}
	err = db.Debug().Table("tbl_service").Where("id = ? and deleted = 0", ser.Id).
		Updates(attr).
		Error

	if err != nil {
		return err
	}
	return nil
}

// QueryTotalMoney 查询用户服务总金额
func QueryTotalMoney(username string) (float64, error) {
	db, err := DBConn()
	if err != nil {
		return 0, err
	}

	total := 0.0
	err = db.Debug().Table("tbl_user_service").Select("ROUND(SUM(`service_cost` * `service_time`), 2)").
		Where("deleted = 0 and user_name = ? and is_used = 1", username).Find(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}

// PayMoney 付钱
func PayMoney(apiPay model.APIPayMoney) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	err = db.Debug().Table("tbl_user_service").Where("user_name = ? and deleted = 0 and is_used = 1", apiPay.UserName).Update("deleted", 1).Error
	if err != nil {
		return err
	}
	err = db.Debug().Table("tbl_user").Where("user_name = ? and deleted = 0", apiPay.UserName).Update("money", apiPay.Money).Error
	if err != nil {
		return err
	}
	return nil
}
