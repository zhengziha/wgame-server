package model

import "time"

// Accounts 账号表（登录账号主表）。
// 注意：supperPassword 在 Java 实体中声明但未出现在
// wd-auth-18.sql 建表语句中，按普通列保留以兼容 Java 逻辑。
type Accounts struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 登录账号（唯一）
	Name string `gorm:"uniqueIndex;size:128;not null;column:name" json:"name"`

	// Keyword 关键字
	Keyword string `gorm:"size:255;not null;column:keyword" json:"keyword"`

	// Password 密码
	Password string `gorm:"size:255;not null;column:password" json:"password"`

	// SupperPassword 超级密码（Java 实体声明，SQL 未建列）
	SupperPassword string `gorm:"size:255;column:supper_password" json:"supperPassword"`

	// Token 令牌
	Token string `gorm:"size:255;column:token" json:"token"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// RegisterCode 注册码
	RegisterCode string `gorm:"size:30;column:register_code" json:"registerCode"`

	// RegisterIp 注册 ip
	RegisterIp string `gorm:"size:30;column:register_ip" json:"registerIp"`

	// LastLoginIp 最后一次登录 ip
	LastLoginIp string `gorm:"size:30;column:last_login_ip" json:"lastLoginIp"`

	// Privilege 特权
	Privilege int32 `gorm:"column:privilege" json:"privilege"`

	// Mac 机器码
	Mac string `gorm:"size:255;column:mac" json:"mac"`

	// LastLoginMac 最后一次登录机器码
	LastLoginMac string `gorm:"size:255;column:last_login_mac" json:"lastLoginMac"`

	// SwitchServerData 切换服务器数据
	SwitchServerData string `gorm:"type:text;column:switch_server_data" json:"switchServerData"`
}

// TableName 显式指定表名
func (Accounts) TableName() string {
	return "accounts"
}
