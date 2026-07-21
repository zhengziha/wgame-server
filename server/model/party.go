package model

import "time"

// Party 帮派表。
type Party struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// PartyId 帮派 id
	PartyId string `gorm:"size:50;column:party_id" json:"partyId"`

	// PartyName 帮派名字
	PartyName string `gorm:"size:255;column:party_name" json:"partyName"`

	// PartyBaseInfo 基本信息
	PartyBaseInfo string `gorm:"size:255;column:party_base_info" json:"partyBaseInfo"`

	// PartyAnnounce 公告帮派宗旨
	PartyAnnounce string `gorm:"size:255;column:party_announce" json:"partyAnnounce"`

	// Rights 权限
	Rights int32 `gorm:"column:rights" json:"rights"`

	// Construct 建设度
	Construct int32 `gorm:"column:construct" json:"construct"`

	// Money 资金
	Money int32 `gorm:"column:money" json:"money"`

	// Salary 俸禄
	Salary int32 `gorm:"column:salary" json:"salary"`

	// AutoAcceptLevel 自动同意等级
	AutoAcceptLevel int32 `gorm:"column:auto_accept_level" json:"autoAcceptLevel"`

	// MinTao 最低道行
	MinTao int32 `gorm:"column:min_tao" json:"minTao"`

	// Creator 创建人
	Creator string `gorm:"size:255;column:creator" json:"creator"`

	// Population 帮派人数（人口）
	Population int32 `gorm:"column:population" json:"population"`

	// PartyLevel 帮派等级
	PartyLevel int32 `gorm:"column:party_level" json:"partyLevel"`

	// Heir 继承人
	Heir string `gorm:"size:255;column:heir" json:"heir"`

	// State 状态
	State int32 `gorm:"column:state" json:"state"`

	// CreateTime 帮派创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// IconMd5 图标 MD5
	IconMd5 string `gorm:"size:255;column:icon_md5" json:"iconMd5"`

	// ReviewIconMd5 预览图标（blob）
	ReviewIconMd5 []byte `gorm:"column:review_icon_md5" json:"reviewIconMd5"`

	// Leader 领导 JSON
	Leader string `gorm:"type:text;column:leader" json:"leader"`
}

// TableName 显式指定表名
func (Party) TableName() string {
	return "party"
}
