package store

import (
	"time"

	"github.com/wqrqwerqrw/cust_gam_backend/model"
)

// CreateUser 添加用户
func CreateUser(name, phone, password string, isVip int) error {

	db, err := DBConn()
	if err != nil {
		return err
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

// DeleteUser 删除用户
func DeleteUser(id int) error {

	db, err := DBConn()
	if err != nil {
		return err
	}

	return db.Debug().Table("tbl_user").Where("id = ? and deleted = 0", id).Update("deleted", 1).Error

}

// UpdateUser 修改用户的基本信息
func UpdateUser(myUser model.APIUserWithId) (*model.User, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	attr := map[string]interface{}{
		"phone":     myUser.Phone,
		"user_pwd":  myUser.Password,
		"is_vip":    myUser.IsVip,
		"user_name": myUser.UserName,
	}

	user := &model.User{}

	err = db.Debug().
		Model(user).
		Where("id = ? and deleted = 0", myUser.Id).
		Updates(attr).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
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
