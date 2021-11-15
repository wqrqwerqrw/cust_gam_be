package store

import (
	"github.com/wqrqwerqrw/cust_gam_backend/model"
	"time"
)

// CreatEquipment 创建设备
func CreatEquipment(apiEquipment model.APIEquipment) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	now := time.Now()

	equipment := &model.Equipment{
		Name:       apiEquipment.Name,
		Usable:     apiEquipment.Usable,
		Desc:       apiEquipment.Desc,
		Stock:      apiEquipment.Stock,
		CreateTime: now,
		UpdateTime: now,
		Deleted:    0,
	}

	err = db.Debug().Create(equipment).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryEquipment 查询设备
func QueryEquipment(id int) (*model.APIEquipmentWithId, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	equipment := &model.APIEquipmentWithId{}

	err = db.Debug().Table("tbl_equipment").Where("id = ? and deleted = 0", id).First(equipment).Error

	if err != nil {
		return nil, err
	}

	return equipment, nil
}

// QueryAllEquipment 查询所有设备
func QueryAllEquipment() (*[]model.APIEquipmentWithId, error) {

	db, err := DBConn()
	if err != nil {
		return nil, err
	}

	equipments := &[]model.APIEquipmentWithId{}

	err = db.Debug().Table("tbl_equipment").Where("deleted = 0").Find(equipments).Error

	if err != nil {
		return nil, err
	}

	return equipments, nil
}
