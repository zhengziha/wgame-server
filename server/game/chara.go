package game

import (
	"sync"
	"time"
)

// Chara 表示玩家运行时数据。
// 参考 Java wd-server-fl core/domain/Chara.java。
//
// 注意：这是内存态对象，不是数据库实体。数据库实体是 model.Characters。
// 登录时从 model.Characters 加载到 Chara，离线时写回。
type Chara struct {
	// 基础标识
	ID    int32  // 数据库主键 characters.id
	Gid   string // 全局唯一 id（客户端使用）
	Name  string
	Sex   int32
	Level int32
	Polar int32 // 门派/阵营

	// 位置信息
	MapId   int32
	MapName string
	Line    int32 // 线路号（1-based）
	X, Y    int32
	Dir     int32 // 朝向

	// 移动状态
	MoveType int32 // 0=正常行走, 1=飞行等

	// 队伍信息
	TeamId       int32
	IsTeamLeader bool

	// 房屋信息
	HouseType int32
	HouseId   int32

	// 显示相关
	IsHide      int32
	Opacity     int32
	Camp        int32
	Title       string
	ShieldOther bool // 是否屏蔽周围玩家

	// 宠物/跟宠
	FollowPet          int32
	FlowerChild        int32
	FlowerChildVisible int32
	GenchongIcon       int32

	// 坐骑/外观
	CscwQiaozhuang int32

	// 时间戳
	LastUpdate time.Time

	mu sync.RWMutex
}

// FollowPetId 返回跟宠 id（用于 disappear 消息）
func (c *Chara) FollowPetId() int32 {
	return c.FollowPet
}

// IsFight 判断是否在战斗中（占位，后续实现）
func (c *Chara) IsFight() bool {
	return false
}

// GetMoveType 返回移动类型
func (c *Chara) GetMoveType() int32 {
	return c.MoveType
}

// GetSetting 获取玩家设置（占位）
func (c *Chara) GetSetting(key string) int32 {
	switch key {
	case "sight_scope":
		return 0 // 默认正常视野
	default:
		return 0
	}
}
