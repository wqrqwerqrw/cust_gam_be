package model

import (
	"time"
)

type TblService struct {
	ID         int       `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`              // 服务项目名称
	Cost       float32   `json:"cost" gorm:"column:cost"`              // 服务费用,元/小时
	Desc       string    `json:"desc" gorm:"column:desc"`              // 服务项目详细说明
	Tag        string    `json:"tag" gorm:"column:tag"`                // 标签
	Extra2     string    `json:"extra2" gorm:"column:extra2"`          // 扩展字段2
	Extra3     string    `json:"extra3" gorm:"column:extra3"`          // 扩展字段3
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"` // 更新日期
	Deleted    int8      `json:"deleted" gorm:"column:deleted"`        // 状态(可用0/已删除1状态)
}

func (m *TblService) TableName() string {
	return "tbl_service"
}

type TblUserService struct {
	ID          int       `json:"id" gorm:"column:id"`
	UserName    string    `json:"username" gorm:"column:user_name"`       // 用户名
	ServiceName string    `json:"serviceName" gorm:"column:service_name"` // 服务项目名称
	ServiceCost float32   `json:"serviceCost" gorm:"column:service_cost"` // 服务费用,元/小时
	ServiceTime int       `json:"serviceTime" gorm:"column:service_time"` // 使用时间,向上取整
	IsUsed      int8      `json:"isUsed" gorm:"column:is_used"`           // 0:没有使用，1:使用过
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`   // 创建时间
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time"`   // 更新日期
	Deleted     int8      `json:"deleted" gorm:"column:deleted"`          // 状态(可用0/已删除1状态)
}

func (m *TblUserService) TableName() string {
	return "tbl_user_service"
}

// APIService 前端传入的
type APIService struct {
	Name string  `json:"name"`
	Tag  string  `json:"tag"`
	Cost float32 `json:"cost"`
	Desc string  `json:"desc"`
}

// APIServiceWithId 前端传入的
type APIServiceWithId struct {
	Id   int     `json:"id"`
	Tag  string  `json:"tag"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Desc string  `json:"desc"`
}

// APIUserService 前端传入
type APIUserService struct {
	UserName    string  `json:"username" gorm:"column:user_name"`       // 用户名
	ServiceName string  `json:"serviceName" gorm:"column:service_name"` // 服务项目名称
	ServiceCost float32 `json:"serviceCost" gorm:"column:service_cost"` // 服务费用,元/小时
	ServiceTime int     `json:"serviceTime" gorm:"column:service_time"` // 使用时间,向上取整
	IsUsed      int8    `json:"isUsed" gorm:"column:is_used"`           // 0:没有使用，1:使用过
}

// APIUserServiceWithId 前端传入
type APIUserServiceWithId struct {
	Id          int     `json:"id"`
	UserName    string  `json:"username"`
	ServiceName string  `json:"serviceName"`
	ServiceCost float32 `json:"serviceCost"`
	ServiceTime int     `json:"serviceTime"`
	IsUsed      int8    `json:"isUsed"`
}

// APIPayMoney 付钱，前端传入
type APIPayMoney struct {
	UserName string  `json:"username"`
	Money    float64 `json:"money"`
}
