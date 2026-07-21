package model

// AutoShuaguai 自动刷怪表。
//
// 字段 state/pkings/bossNum/gid/line/npc 在 Java 中标注 @Transient，
// taskName/maxKillNum/consumeItems 虽标注 @Column 但未出现在
// wd-game-18.sql 建表语句中，按非持久化字段处理（gorm:"-"）。
// startTime/endTime 在 SQL 中为 time 类型，Java 中按 String 处理，这里保持 string。
type AutoShuaguai struct {
	// ID 数据库 id
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name boss 名字
	Name string `gorm:"size:255;not null;column:name" json:"name"`

	// TaskName 归属活动名字（非持久化：SQL 无此列）
	TaskName string `gorm:"-" json:"taskName"`

	// Type 类型（0 怪物,1 boss 单杀,2 npc 多人杀）
	Type int32 `gorm:"column:type" json:"type"`

	// Icon 模型
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Mapid 地图 id
	Mapid int32 `gorm:"column:mapid" json:"mapid"`

	// Mapname 地图名称
	Mapname string `gorm:"size:255;column:mapname" json:"mapname"`

	// Content 内容
	Content string `gorm:"size:255;column:content" json:"content"`

	// Btns 按钮列表
	Btns string `gorm:"size:255;column:btns" json:"btns"`

	// Guaiwus 怪物列表
	Guaiwus string `gorm:"size:255;column:guaiwus" json:"guaiwus"`

	// Reward 奖励列表
	Reward string `gorm:"size:3000;column:reward" json:"reward"`

	// Xys 随机的地图坐标
	Xys string `gorm:"size:255;column:xys" json:"xys"`

	// Resttime 刷新间隔（秒）
	Resttime int32 `gorm:"column:resttime" json:"resttime"`

	// Pklevel 挑战需要的等级
	Pklevel int32 `gorm:"column:pklevel" json:"pklevel"`

	// Pkdaohang 挑战需要的道行（年）
	Pkdaohang int32 `gorm:"column:pkdaohang" json:"pkdaohang"`

	// MaxKillNum 最大击杀数（非持久化：SQL 无此列）
	MaxKillNum int32 `gorm:"-" json:"maxKillNum"`

	// ConsumeItems 消耗物品（非持久化：SQL 无此列）
	ConsumeItems string `gorm:"-" json:"consumeItems"`

	// StartTime 开始时间（SQL 列 startTime 为 time 类型，按字符串处理）
	StartTime string `gorm:"column:startTime" json:"startTime"`

	// EndTime 结束时间（SQL 列 endTime 为 time 类型，按字符串处理）
	EndTime string `gorm:"column:endTime" json:"endTime"`

	// State 1 正常出现,0 隐藏或死亡（非持久化：@Transient）
	State int32 `gorm:"-" json:"state"`

	// Pkings 被 PK 的次数（非持久化：@Transient）
	Pkings int32 `gorm:"-" json:"pkings"`

	// BossNum boss 数量（非持久化：@Transient）
	BossNum int32 `gorm:"-" json:"bossNum"`

	// Gid bossid（非持久化：@Transient）
	Gid int32 `gorm:"-" json:"gid"`

	// Line 抢怪模式线路（非持久化：@Transient）
	Line int32 `gorm:"-" json:"line"`

	// Npc 运行时构造的 NPC 对象（非持久化：@Transient）
	Npc interface{} `gorm:"-" json:"npc"`
}

// TableName 显式指定表名
func (AutoShuaguai) TableName() string {
	return "auto_shuaguai"
}
