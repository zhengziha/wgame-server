package game

import "sync"

// GameMap 表示游戏地图。
// 参考 Java wd-server-fl core/GameMap.java。
//
// 包含：
//   - 静态地图配置（从 map_info 加载）
//   - 当前在线玩家列表（sessionList）
//   - AOI 九宫格视野管理
type GameMap struct {
	Index    int32  // 数据库索引 map_info.id
	ID       int32  // 地图编号 map_info.map_id
	Name     string // 地图名称
	ShowName string // 显示名称
	X, Y     int32  // 默认出生点坐标
	MapType  int32  // 0=普通, 1=组队限制, 2=副本

	LineNum int32 // 所属线路号（1-based）

	// 当前地图的玩家列表 key=gid
	sessionList *sync.Map // map[string]struct{}

	// AOI 九宫格管理器
	AOI *AOI

	mu sync.RWMutex
}

// NewGameMap 创建一个新地图实例
func NewGameMap(index, id, x, y, lineNum, mapType int32, name, showName string) *GameMap {
	return &GameMap{
		Index:       index,
		ID:          id,
		X:           x,
		Y:           y,
		LineNum:     lineNum,
		MapType:     mapType,
		Name:        name,
		ShowName:    showName,
		sessionList: &sync.Map{},
		AOI:         NewAOI(100, 100, 30), // 默认 100x100 格子，每格 30 像素
	}
}

// Join 玩家进入地图
func (m *GameMap) Join(gid string) {
	m.sessionList.Store(gid, struct{}{})
}

// Leave 玩家离开地图
func (m *GameMap) Leave(gid string) {
	m.sessionList.Delete(gid)
}

// SessionCount 返回当前地图玩家数量
func (m *GameMap) SessionCount() int {
	count := 0
	m.sessionList.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// IsZone 判断是否为区域地图（地图类型 > 0）
func (m *GameMap) IsZone() bool {
	return m.MapType > 0
}

// IsDugeon 判断是否为副本地图（地图类型 > 1）
func (m *GameMap) IsDugeon() bool {
	return m.MapType > 1
}

// IsZhengDaoDianMap 判断是否为证道殿地图
func (m *GameMap) IsZhengDaoDianMap() bool {
	return m.ID == 29002
}

// IsXcwq 判断是否为玉露仙池地图
func (m *GameMap) IsXcwq() bool {
	return m.ID == 37013 || m.Name == "玉露仙池"
}

// CanSee 判断两个角色是否互相可见（同地图、同线路、距离范围内）
// 视野范围：x <= 36, y <= 30
func CanSee(c1, c2 *Chara) bool {
	return c1.MapName == c2.MapName &&
		c1.Line == c2.Line &&
		abs(c1.X-c2.X) <= 36 &&
		abs(c1.Y-c2.Y) <= 30
}

func abs(v int32) int32 {
	if v < 0 {
		return -v
	}
	return v
}

// RangeSessionList 遍历当前地图所有玩家
func (m *GameMap) RangeSessionList(fn func(gid string) bool) {
	m.sessionList.Range(func(key, value interface{}) bool {
		return fn(key.(string))
	})
}
