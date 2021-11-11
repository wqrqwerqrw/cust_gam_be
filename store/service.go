package store

import (
	"github.com/wqrqwerqrw/cust_gam_backend/model"
)

func QueryByUserName(username string) (*[]model.TblUserService, error) {
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
