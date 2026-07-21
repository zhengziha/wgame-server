package model

import "time"

// LifeDeathRecord 生死（生死战）记录表。
// status、time 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// Java 实体部分字段使用 snake_case 命名（bet_type/bet_num/att_info_members/def_info_members），
// Go 字段转为驼峰。
type LifeDeathRecord struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Uuid 对战 uuid
	Uuid string `gorm:"size:64;not null;column:uuid" json:"uuid"`

	// AttGid 挑战方 gid
	AttGid string `gorm:"size:64;column:att_gid" json:"attGid"`

	// AttIcon 挑战方头像
	AttIcon int32 `gorm:"column:att_icon" json:"attIcon"`

	// AttName 挑战方名字
	AttName string `gorm:"size:255;column:att_name" json:"attName"`

	// AttLevel 挑战方等级
	AttLevel int32 `gorm:"column:att_level" json:"attLevel"`

	// DefGid 应战方 gid
	DefGid string `gorm:"size:64;column:def_gid" json:"defGid"`

	// DefIcon 应战方头像
	DefIcon int32 `gorm:"column:def_icon" json:"defIcon"`

	// DefName 应战方名字
	DefName string `gorm:"size:255;column:def_name" json:"defName"`

	// DefLevel 应战方等级
	DefLevel int32 `gorm:"column:def_level" json:"defLevel"`

	// Status 状态（列名 status 为 SQL 保留字）
	// atk_raise 待应战方接受 / def_accept 应战方接受 / def_refuse 应战方拒绝 / over_time 超时
	Status string `gorm:"size:32;column:status" json:"status"`

	// Result 结果（no_start 双方均败 / atk 挑战方胜 / def 应战方胜 / draw 平）
	Result string `gorm:"size:32;column:result" json:"result"`

	// Time 战斗时间戳（列名 time 为 SQL 保留字）
	Time int32 `gorm:"column:time" json:"time"`

	// Mode 模式
	Mode string `gorm:"size:32;column:mode" json:"mode"`

	// BetType 押注类型
	BetType string `gorm:"size:32;column:bet_type" json:"betType"`

	// BetNum 押注数量
	BetNum int32 `gorm:"column:bet_num" json:"betNum"`

	// Server 所在线路
	Server string `gorm:"size:255;column:server" json:"server"`

	// AttInfoMembers 挑战方成员信息 JSON
	AttInfoMembers string `gorm:"type:text;column:att_info_members" json:"attInfoMembers"`

	// DefInfoMembers 应战方成员信息 JSON
	DefInfoMembers string `gorm:"type:text;column:def_info_members" json:"defInfoMembers"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (LifeDeathRecord) TableName() string {
	return "life_death_record"
}
