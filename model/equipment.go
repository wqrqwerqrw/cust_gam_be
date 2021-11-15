package model

import (
	"time"
)

type Equipment struct {
	ID         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`              // 健身设备名称
	Usable     int8      `json:"usable" gorm:"column:usable"`          // 是否可以使用(可用1/不可用0)
	Desc       string    `json:"desc" gorm:"column:desc"`              // 健身设备详细说明
	Stock      int       `json:"stock" gorm:"column:stock"`            // 库存
	Extra2     string    `json:"extra2" gorm:"column:extra2"`          // 扩展字段2
	Extra3     string    `json:"extra3" gorm:"column:extra3"`          // 扩展字段3
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"` // 更新日期
	Deleted    int8      `json:"deleted" gorm:"column:deleted"`        // 状态(可用0/已删除1状态)
}

func (m *Equipment) TableName() string {
	return "tbl_equipment"
}

type APIEquipment struct {
	Name   string `json:"name"`
	Usable int8   `json:"usable"`
	Desc   string `json:"desc"`
	Stock  int    `json:"stock"`
}

type APIEquipmentWithId struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Usable int8   `json:"usable"`
	Desc   string `json:"desc"`
	Stock  int    `json:"stock"`
}
