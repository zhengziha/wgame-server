package model

// Conn 连接记录表（清理缓存用的连接配置）。
// 注意：dist、port、dbname 在 Java 实体中声明但未出现在
// wd-auth-18.sql 建表语句中，按普通列保留以兼容 Java 逻辑。
type Conn struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Dist 区服标识（Java 实体声明，SQL 未建列）
	Dist string `gorm:"size:255;column:dist" json:"dist"`

	// IP 地址
	IP string `gorm:"size:255;column:ip" json:"ip"`

	// Port 端口（Java 实体声明，SQL 未建列）
	Port int32 `gorm:"column:port" json:"port"`

	// Username 用户名
	Username string `gorm:"size:255;column:username" json:"username"`

	// Password 密码
	Password string `gorm:"size:255;column:password" json:"password"`

	// DBName 数据库名（Java 实体声明，SQL 未建列）
	DBName string `gorm:"size:255;column:dbname" json:"dbname"`

	// URL 清理缓存用的 url
	URL string `gorm:"size:255;column:url" json:"url"`
}

// TableName 显式指定表名
func (Conn) TableName() string {
	return "conn"
}
