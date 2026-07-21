package model

// ChargePoint 充值积分表。
// 注意：SQL 中 awardStr 列为驼峰命名；status 列名为 SQL 保留字。
type ChargePoint struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 序号（从 0 开始）
	No int32 `gorm:"column:no" json:"no"`

	// Awardstr 奖励（SQL 列名为驼峰 awardStr）
	Awardstr string `gorm:"size:255;not null;column:awardStr" json:"awardstr"`

	// Point 消耗积分
	Point int32 `gorm:"column:point" json:"point"`

	// LeftNum 剩余数量
	LeftNum int32 `gorm:"column:left_num" json:"leftNum"`

	// Status 状态（列名 status 为 SQL 保留字）
	Status int32 `gorm:"column:status" json:"status"`
}

// TableName 显式指定表名
func (ChargePoint) TableName() string {
	return "charge_point"
}
