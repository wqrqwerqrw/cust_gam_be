package store

import (
	"fmt"
	"time"

	"github.com/wqrqwerqrw/cust_gam_backend/model"
)

// CreateUser 添加用户
func CreateUser(name, phone, password string, isVip int) (*model.User, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := &model.User{
		UserName:   name,
		Phone:      phone,
		UserPwd:    password,
		IsVip:      isVip,
		CreateTime: now,
		UpdateTime: now,
	}

	err = db.Debug().Create(user).Error

	if err != nil {
		return nil, err
	}

	fmt.Print(user)

	return user, nil
}

// QueryUser 查询用户
func QueryUser(id int) (*model.User, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	user := &model.User{}

	err = db.Debug().Table("tbl_user").Where("id = ? and deleted = 0", id).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
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
func UpdateUser(id int, name, phone, password string, isVip int) (*model.User, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	attr := map[string]interface{}{
		"phone":     phone,
		"user_pwd":  password,
		"is_vip":    isVip,
		"user_name": name,
	}

	user := &model.User{}

	err = db.Debug().
		Model(user).
		Where("id = ? and deleted = 0", id).
		Updates(attr).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil

}
