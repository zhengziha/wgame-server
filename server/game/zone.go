package game

// GameZone 表示动态地图（副本、试道、帮战等）。
// 参考 Java wd-server-fl core/GameZone.java。
//
// GameZone 通过组合方式持有 GameMap，增加了：
//   - uid：动态地图唯一标识
//   - gameDugeon：副本数据（后续扩展）
//   - isHouseZone：是否为房屋区域
type GameZone struct {
	// 基础地图配置（只读，不修改）
	BaseMap *GameMap

	// 动态地图属性（覆盖或扩展基础配置）
	Uid         string // 动态地图唯一标识
	MapType     int32  // 覆盖基础地图的 MapType
	GameDugeon  *GameDugeon
	IsHouseZone bool
	Forever     bool // 是否永久存在

	endTime int64 // 生命周期结束时间（-1 表示永久）
}

// GameDugeon 表示副本数据（占位，后续实现）
type GameDugeon struct {
	Type       string
	DugeonName string
}

// NewGameZone 创建一个新的动态地图
func NewGameZone(baseMap *GameMap, uid string) *GameZone {
	return &GameZone{
		BaseMap:     baseMap,
		Uid:         uid,
		MapType:     1, // 动态地图默认 map_type=1
		IsHouseZone: false,
		Forever:     false,
		endTime:     -1,
	}
}

// SetLifeTime 设置动态地图生命周期（单位：毫秒）
func (z *GameZone) SetLifeTime(lifeTimeMs int64) {
	if lifeTimeMs <= 0 {
		z.endTime = -1
		return
	}
	z.endTime = lifeTimeMs // TODO: 需要当前时间 + lifeTimeMs
}

// IsExpired 判断动态地图是否已过期
func (z *GameZone) IsExpired() bool {
	if z.endTime <= 0 {
		return false
	}
	return false // 暂未实现时间判断
}

// 以下方法代理到 BaseMap，保持接口兼容

func (z *GameZone) Index() int32     { return z.BaseMap.Index }
func (z *GameZone) ID() int32        { return z.BaseMap.ID }
func (z *GameZone) Name() string     { return z.BaseMap.Name }
func (z *GameZone) ShowName() string { return z.BaseMap.ShowName }
func (z *GameZone) X() int32         { return z.BaseMap.X }
func (z *GameZone) Y() int32         { return z.BaseMap.Y }
func (z *GameZone) LineNum() int32   { return z.BaseMap.LineNum }

func (z *GameZone) Join(sess interface{}) {
	_ = sess
}

func (z *GameZone) Leave(sess interface{}) {
	_ = sess
}
