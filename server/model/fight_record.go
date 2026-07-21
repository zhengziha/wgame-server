package model

import "time"

// FightRecord 战斗记录表。
// time 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type FightRecord struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// CombatId 战斗 id
	CombatId string `gorm:"size:128;column:combat_id" json:"combatId"`

	// Dist 分布标识
	Dist string `gorm:"size:16;column:dist" json:"dist"`

	// CombatType 战斗类型
	CombatType string `gorm:"size:64;column:combat_type" json:"combatType"`

	// AtkName 进攻方名字
	AtkName string `gorm:"size:32;column:atk_name" json:"atkName"`

	// DefName 防守方名字
	DefName string `gorm:"size:32;column:def_name" json:"defName"`

	// Time 战斗时间戳（列名 time 为 SQL 保留字）
	Time int32 `gorm:"column:time" json:"time"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// RecordDataMap 战斗记录 JSON
	RecordDataMap string `gorm:"type:mediumtext;column:record_data_map" json:"recordDataMap"`
}

// TableName 显式指定表名
func (FightRecord) TableName() string {
	return "fight_record"
}
