package model

import (
	"time"
)

type TblEquipment struct {
	ID         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`               // 健身设备名称
	PhotoUrl   string    `json:"photo_url" gorm:"column:photo_url"`     // 图片url地址
	Desc       string    `json:"desc" gorm:"column:desc"`               // 健身设备详细说明
	Extra1     string    `json:"extra1" gorm:"column:extra1"`           // 扩展字段1
	Extra2     string    `json:"extra2" gorm:"column:extra2"`           // 扩展字段2
	Extra3     string    `json:"extra3" gorm:"column:extra3"`           // 扩展字段3
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新日期
	Deleted    int8      `json:"deleted" gorm:"column:deleted"`         // 状态(可用0/已删除1状态)
}

func (m *TblEquipment) TableName() string {
	return "tbl_equipment"
}
