package store

import (
	"time"

	"github.com/wqrqwerqrw/cust_gam_backend/model"
)

// CreateUser 添加用户
func CreateUser(apiUser model.APIUser) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	now := time.Now()

	user := &model.User{
		UserName:   apiUser.UserName,
		Phone:      apiUser.Phone,
		IsVip:      apiUser.IsVip,
		Money:      apiUser.Money,
		CreateTime: now,
		UpdateTime: now,
		Deleted:    0,
	}

	err = db.Debug().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryUser 查询用户
func QueryUser(id int) (*model.APIUserWithId, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	user := &model.APIUserWithId{}

	err = db.Debug().Table("tbl_user").Where("id = ? and deleted = 0", id).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// QueryAllUser 查询所有用户
func QueryAllUser() (*[]model.APIUserWithId, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	users := &[]model.APIUserWithId{}

	err = db.Debug().Table("tbl_user").Where("deleted = 0").Find(users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser 删除用户
func DeleteUser(id int) error {

	db, err := DBConn()
	if err != nil {
		return err
	}

	return db.Debug().Table("tbl_user").Where("id = ? and deleted = 0", id).Update("deleted", 1).Error

}

// UpdateUser 修改用户的基本信息
func UpdateUser(myUser model.APIUserWithId) error {

	db, err := DBConn()
	if err != nil {
		return err
	}

	attr := map[string]interface{}{
		"phone":     myUser.Phone,
		"is_vip":    myUser.IsVip,
		"user_name": myUser.UserName,
		"money":     myUser.Money,
	}

	err = db.Debug().Table("tbl_user").
		Where("id = ? and deleted = 0", myUser.Id).
		Updates(attr).
		Error

	if err != nil {
		return err
	}

	return nil
}

// QueryUserByUserName 查询用户
func QueryUserByUserName(username string) (*model.APIUserWithId, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	user := &model.APIUserWithId{}

	err = db.Debug().Table("tbl_user").Where("user_name = ? and deleted = 0", username).First(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
