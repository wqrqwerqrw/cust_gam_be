package model

import "time"

// User 用户表实体类
type User struct {
	ID         int       `json:"id" gorm:"column:id"`
	UserName   string    `json:"username" gorm:"column:user_name"`     // 用户名
	Phone      string    `json:"phone" gorm:"column:phone"`            // 手机号
	IsVip      int       `json:"isVip" gorm:"column:is_vip"`           // 是否是会员(非会员0/会员1)
	Money      float64   `json:"money" gorm:"column:money"`            // 剩余余额
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"` // 更新日期
	Deleted    int       `json:"deleted" gorm:"column:deleted"`        // 状态(可用0/已删除1)
}

func (m *User) TableName() string {
	return "tbl_user"
}

// APIUser 前端注册传入
type APIUser struct {
	UserName string  `json:"username"`
	Phone    string  `json:"phone"`
	IsVip    int     `json:"isVip"`
	Money    float64 `json:"money"`
}

// APIUserWithId 带id的
type APIUserWithId struct {
	Id       int     `json:"id"`
	UserName string  `json:"username"`
	Phone    string  `json:"phone"`
	IsVip    int     `json:"isVip"`
	Money    float64 `json:"money"`
}
