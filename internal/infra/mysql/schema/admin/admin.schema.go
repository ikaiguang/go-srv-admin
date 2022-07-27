// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	gorm "gorm.io/gorm"
	"time"
)

var _ = time.Time{}

// AdminSchema Admin
var AdminSchema Admin

// NewAdmin new schema
func NewAdmin() *Admin {
	return &Admin{}
}

// Admin ENGINE InnoDB CHARSET utf8mb4 COMMENT '管理员'
type Admin struct {
	Id              uint64     `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	AdminUuid       string     `gorm:"column:admin_uuid;unique;type:string;size:255;not null;default:'';comment:uuid" json:"admin_uuid"`
	CreatedTime     time.Time  `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime     time.Time  `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	DeletedTime     *time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
	IsDeleted       bool       `gorm:"column:is_deleted;index;type:bool;not null;default:0;comment:是否已删除" json:"is_deleted"`
	AdminEmail      string     `gorm:"column:admin_email;type:string;size:255;not null;default:'';comment:邮箱" json:"admin_email"`
	AdminNickname   string     `gorm:"column:admin_nickname;type:string;size:255;not null;default:'';comment:昵称" json:"admin_nickname"`
	AdminAvatar     string     `gorm:"column:admin_avatar;type:string;size:255;not null;default:'';comment:头像" json:"admin_avatar"`
	AdminGender     string     `gorm:"column:admin_gender;type:string;size:255;not null;default:'';comment:性别" json:"admin_gender"`
	AdminStatus     string     `gorm:"column:admin_status;type:string;size:255;not null;default:'';comment:状态" json:"admin_status"`
	ActiveBeginTime *time.Time `gorm:"column:active_begin_time;type:time;comment:激活开始时间" json:"active_begin_time"`
	ActiveEndTime   *time.Time `gorm:"column:active_end_time;type:time;comment:激活结束时间" json:"active_end_time"`
	DisableTime     *time.Time `gorm:"column:disable_time;type:time;comment:禁用时间" json:"disable_time"`
	BlacklistTime   *time.Time `gorm:"column:blacklist_time;type:time;comment:黑名单时间" json:"blacklist_time"`
	PasswordHash    string     `gorm:"column:password_hash;type:string;size:255;not null;default:'';comment:密码" json:"password_hash"`
	RegisterFrom    string     `gorm:"column:register_from;type:string;size:255;not null;default:'';comment:注册来源" json:"register_from"`
}

// TableName table name
func (s *Admin) TableName() string {
	return "srv_admin"
}

// CreateTableMigrator create table migrator
func (s *Admin) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewCreateTable(migrator, s)
}

// DropTableMigrator create table migrator
func (s *Admin) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewDropTable(migrator, s)
}
