package model

import "time"

// SysUser 系统账号管理表（后台用户）。
// 注意：主键 id 为 Long 类型，对应 bigint(20)，使用 int64。
type SysUser struct {
	// ID 主键
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// UserName 登录名（全网唯一，不可更改）
	UserName string `gorm:"uniqueIndex;size:255;not null;column:user_name" json:"userName"`

	// Password 登录密码
	Password string `gorm:"size:255;column:password" json:"password"`

	// UserType 用户类型（0:超级管理员 1:管理员）
	UserType int32 `gorm:"not null;column:user_type" json:"userType"`

	// HeadImgUrl 头像 URL
	HeadImgUrl string `gorm:"size:255;column:head_img_url" json:"headImgUrl"`

	// Sex 性别
	Sex int32 `gorm:"column:sex" json:"sex"`

	// NickName 昵称
	NickName string `gorm:"size:255;column:nick_name" json:"nickName"`

	// State 状态（0:启用 1:禁用）
	State int32 `gorm:"column:state" json:"state"`

	// FirstLoginUpdatePasswordFlag 第一次登录是否修改密码（0:要修改 1:不修改）
	FirstLoginUpdatePasswordFlag int32 `gorm:"column:first_login_update_password_flag" json:"firstLoginUpdatePasswordFlag"`

	// CreateUserId 创建用户 id
	CreateUserId int64 `gorm:"column:create_user_id" json:"createUserId"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// LastLoginTime 最后一次登录时间
	LastLoginTime time.Time `gorm:"column:last_login_time" json:"lastLoginTime"`

	// LastLoginIp 最后一次登录 ip 地址
	LastLoginIp string `gorm:"size:255;column:last_login_ip" json:"lastLoginIp"`

	// AllowLoginIp 允许登录的 ip 地址
	AllowLoginIp string `gorm:"size:255;column:allow_login_ip" json:"allowLoginIp"`

	// MobilePhone 手机号码
	MobilePhone string `gorm:"size:255;column:mobile_phone" json:"mobilePhone"`

	// Email 邮箱
	Email string `gorm:"size:255;column:email" json:"email"`
}

// TableName 显式指定表名
func (SysUser) TableName() string {
	return "sys_user"
}
