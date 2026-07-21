package model

// MailboxRefresh 邮箱刷新表。
// status 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// createTime/expiredTime 在 SQL 中为 int(11) 时间戳，按 Java 类型映射保留为 int32。
type MailboxRefresh struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Count 数量
	Count int32 `gorm:"default:0;column:count" json:"count"`

	// Gid 我方 gid
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// ToGid 接收人的 gid
	ToGid string `gorm:"size:255;column:to_gid" json:"toGid"`

	// Type 邮件类型
	Type int32 `gorm:"column:type" json:"type"`

	// Sender 发件人
	Sender string `gorm:"size:255;column:sender" json:"sender"`

	// Title 标题
	Title string `gorm:"size:255;column:title" json:"title"`

	// Msg 消息内容
	Msg string `gorm:"size:255;column:msg" json:"msg"`

	// Attachment 附件
	Attachment string `gorm:"size:255;default:'';column:attachment" json:"attachment"`

	// Status 状态（0:未查看,1:已查看,2:已领取,3:删除）
	Status int32 `gorm:"column:status" json:"status"`

	// CreateTime 创建时间（int 时间戳）
	CreateTime int32 `gorm:"column:create_time" json:"createTime"`

	// ExpiredTime 过期时间（int 时间戳）
	ExpiredTime int32 `gorm:"column:expired_time" json:"expiredTime"`
}

// TableName 显式指定表名
func (MailboxRefresh) TableName() string {
	return "mailbox_refresh"
}
