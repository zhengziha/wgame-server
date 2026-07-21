package model

import "time"

// HouseData 房屋数据表。
type HouseData struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// HouseId 居所 id
	HouseId string `gorm:"size:50;column:house_id" json:"houseId"`

	// Gid 玩家 gid
	Gid string `gorm:"size:50;column:gid" json:"gid"`

	// HouseClass 居所分类（1:私人 2:夫妻）
	HouseClass int32 `gorm:"column:house_class" json:"houseClass"`

	// HouseName 居所名称前缀
	HouseName string `gorm:"size:15;column:house_name" json:"houseName"`

	// Comfort 舒适度
	Comfort int32 `gorm:"column:comfort" json:"comfort"`

	// Cleanliness 清洁度
	Cleanliness int32 `gorm:"column:cleanliness" json:"cleanliness"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (HouseData) TableName() string {
	return "house_data"
}
