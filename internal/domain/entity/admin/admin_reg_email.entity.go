// Package entities
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package entities

import (
	"time"
)

var _ = time.Time{}

// AdminRegEmail ENGINE InnoDB CHARSET utf8mb4 COMMENT '管理员-注册-邮箱'
type AdminRegEmail struct {
	Id          uint64    `gorm:"column:id;primaryKey" json:"id"`          // ID
	AdminEmail  string    `gorm:"column:admin_email" json:"admin_email"`   // 管理员手机号码
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"` // 最后修改时间
	AdminUuid   string    `gorm:"column:admin_uuid" json:"admin_uuid"`     // uuid
}

// TableName table name
func (s *AdminRegEmail) TableName() string {
	return "srv_admin_reg_email"
}