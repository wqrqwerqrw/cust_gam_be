package model

import (
	"time"
)

type TblService struct {
	ID         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`               // 服务项目名称
	Cost       float32   `json:"cost" gorm:"column:cost"`               // 服务费用,元/小时
	Desc       string    `json:"desc" gorm:"column:desc"`               // 服务项目详细说明
	Extra1     string    `json:"extra1" gorm:"column:extra1"`           // 扩展字段1
	Extra2     string    `json:"extra2" gorm:"column:extra2"`           // 扩展字段2
	Extra3     string    `json:"extra3" gorm:"column:extra3"`           // 扩展字段3
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新日期
	Deleted    int8      `json:"deleted" gorm:"column:deleted"`         // 状态(可用0/已删除1状态)
}

func (m *TblService) TableName() string {
	return "tbl_service"
}

type TblUserService struct {
	ID          int       `json:"id" gorm:"column:id"`
	UserName    string    `json:"user_name" gorm:"column:user_name"`       // 用户名
	ServiceName string    `json:"service_name" gorm:"column:service_name"` // 服务项目名称
	ServiceCost float32   `json:"service_cost" gorm:"column:service_cost"` // 服务费用,元/小时
	ServiceTime int       `json:"service_time" gorm:"column:service_time"` // 使用时间,向上取整
	IsUsed      int8      `json:"is_used" gorm:"column:is_used"`           // 0:没有使用，1:使用过
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`   // 创建时间
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`   // 更新日期
	Deleted     int8      `json:"deleted" gorm:"column:deleted"`           // 状态(可用0/已删除1状态)
}

func (m *TblUserService) TableName() string {
	return "tbl_user_service"
}
