package model

// ArenaRank 竞技场排名表。
// 对应 Java 实体 Arena_Rank，Go 结构体命名按规范去掉下划线改为驼峰。
// 注意：SQL 中列名 serverId / charaId 为驼峰命名（非常规 snake_case），
// 按真实列名写入 column tag。
type ArenaRank struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// ServerId 服务器 id（SQL 列名 serverId）
	ServerId string `gorm:"size:100;column:serverId" json:"serverId"`

	// Rankpm 排名
	Rankpm int32 `gorm:"column:rankpm" json:"rankpm"`

	// CharaId 角色 id（SQL 列名 charaId）
	CharaId int32 `gorm:"column:charaId" json:"charaId"`
}

// TableName 显式指定表名
func (ArenaRank) TableName() string {
	return "arena_rank"
}
