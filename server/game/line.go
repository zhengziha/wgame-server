package game

import (
	"sync"

	"wgame-server/server/model"
)

// GameLine 表示游戏线路。
// 参考 Java wd-server-fl core/GameLine.java。
//
// 包含：
//   - 线路号（lineNum）
//   - 静态地图列表（gameMapList）
//   - 动态地图列表（gameZoneList，如副本、试道、帮战）
type GameLine struct {
	LineNum int32

	// 静态地图列表（从 map_info 加载，只读）
	gameMapList []*GameMap

	// 地图编号到 GameMap 的映射（只读）
	gameMapIdMap map[int32]*GameMap

	// 地图名称到 GameMap 的映射（只读）
	gameRoomNameMap map[string]*GameMap

	// 动态地图列表（运行时创建，如副本、试道）
	gameZoneList []*GameZone

	// 动态地图 uid 到 GameZone 的映射
	gameZoneUidMap map[string]*GameZone

	mu sync.RWMutex
}

// NewGameLine 创建一个新线路
func NewGameLine(lineNum int32) *GameLine {
	return &GameLine{
		LineNum:         lineNum,
		gameMapList:     make([]*GameMap, 0),
		gameMapIdMap:    make(map[int32]*GameMap),
		gameRoomNameMap: make(map[string]*GameMap),
		gameZoneList:    make([]*GameZone, 0),
		gameZoneUidMap:  make(map[string]*GameZone),
	}
}

// Init 初始化线路，从数据库加载地图配置
func (l *GameLine) Init(mapInfos []*model.MapInfo) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.gameMapList = make([]*GameMap, 0, len(mapInfos))
	l.gameMapIdMap = make(map[int32]*GameMap, len(mapInfos))
	l.gameRoomNameMap = make(map[string]*GameMap, len(mapInfos))

	for _, mi := range mapInfos {
		gameMap := NewGameMap(
			mi.ID,
			mi.MapID,
			mi.X,
			mi.Y,
			l.LineNum,
			0, // mapType 默认为 0（普通地图）
			mi.Name,
			mi.Name, // showName 暂时用 name 代替
		)
		l.gameMapList = append(l.gameMapList, gameMap)
		l.gameMapIdMap[mi.ID] = gameMap
		l.gameRoomNameMap[mi.Name] = gameMap
	}
}

// GetGameMapById 按地图编号查找地图
func (l *GameLine) GetGameMapById(mapId int32) *GameMap {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.gameMapIdMap[mapId]
}

// GetGameMapByName 按地图名称查找地图
func (l *GameLine) GetGameMapByName(name string) *GameMap {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.gameRoomNameMap[name]
}

// CreateGameZone 创建动态地图
// 返回创建的 GameZone，如果地图不存在则返回 nil
func (l *GameLine) CreateGameZone(mapID int32, uid string) *GameZone {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 先查找基础地图配置（O(1) map 查找）
	baseMap := l.gameMapIdMap[mapID]
	if baseMap == nil {
		return nil
	}

	// 检查是否已存在
	if existing, ok := l.gameZoneUidMap[uid]; ok {
		return existing
	}

	zone := &GameZone{
		BaseMap: baseMap,
		Uid:     uid,
		MapType: 1,
	}

	l.gameZoneList = append(l.gameZoneList, zone)
	l.gameZoneUidMap[uid] = zone

	return zone
}

// DeleteGameZone 删除动态地图
func (l *GameLine) DeleteGameZone(uid string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.gameZoneUidMap, uid)
	for i, zone := range l.gameZoneList {
		if zone.Uid == uid {
			l.gameZoneList = append(l.gameZoneList[:i], l.gameZoneList[i+1:]...)
			break
		}
	}
}

// GetGameZone 按 uid 查找动态地图
func (l *GameLine) GetGameZone(uid string) *GameZone {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.gameZoneUidMap[uid]
}

// MapList 返回静态地图列表（只读）
func (l *GameLine) MapList() []*GameMap {
	l.mu.RLock()
	defer l.mu.RUnlock()
	result := make([]*GameMap, len(l.gameMapList))
	copy(result, l.gameMapList)
	return result
}
