package core

import (
	"sync"

	"wgame-server/comm/log"
	"wgame-server/server/game"
	"wgame-server/server/model"

	"gorm.io/gorm"
)

// GameCore 全局游戏核心管理。
// 参考 Java wd-server-fl core/GameCore.java。
//
// 职责：
//   - 管理游戏线路（GameLine）
//   - 管理初始化状态
//   - 提供全局访问入口
type GameCore struct {
	gameLineList []*game.GameLine
	lineNum      int32 // 线路数量

	mu      sync.RWMutex
	isReady bool
}

var (
	instance *GameCore
	initOnce sync.Once
)

// Instance 返回 GameCore 单例
func Instance() *GameCore {
	initOnce.Do(func() {
		instance = &GameCore{
			gameLineList: make([]*game.GameLine, 0),
			lineNum:      1, // 默认 1 条线路
		}
	})
	return instance
}

// InitLineAndMap 初始化线路和地图
// 从数据库加载 map_info 表，为每条线路创建静态地图
func (c *GameCore) InitLineAndMap(db *gorm.DB) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Info("[core] 开始初始化线路和地图...")

	var mapInfos []*model.MapInfo
	if err := db.Find(&mapInfos).Error; err != nil {
		return err
	}

	c.gameLineList = make([]*game.GameLine, 0, c.lineNum)
	for i := int32(0); i < c.lineNum; i++ {
		line := game.NewGameLine(i + 1)
		line.Init(mapInfos)
		c.gameLineList = append(c.gameLineList, line)
	}

	log.Info("[core] 线路和地图初始化完成：%d 条线路，%d 张地图", c.lineNum, len(mapInfos))
	return nil
}

// SetLineNum 设置线路数量（必须在 InitLineAndMap 之前调用）
func (c *GameCore) SetLineNum(num int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.gameLineList) > 0 {
		log.Info("[core] 线路已初始化，SetLineNum 无效")
		return
	}
	c.lineNum = num
}

// GetGameLine 按线路号获取线路
func (c *GameCore) GetGameLine(lineNum int32) *game.GameLine {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, line := range c.gameLineList {
		if line.LineNum == lineNum {
			return line
		}
	}
	return nil
}

// GetGameMap 按线路号和地图编号获取地图
func (c *GameCore) GetGameMap(lineNum, mapID int32) *game.GameMap {
	line := c.GetGameLine(lineNum)
	if line == nil {
		return nil
	}
	return line.GetGameMapById(mapID)
}

// GetGameMapByName 按线路号和地图名称获取地图
func (c *GameCore) GetGameMapByName(lineNum int32, name string) *game.GameMap {
	line := c.GetGameLine(lineNum)
	if line == nil {
		return nil
	}
	return line.GetGameMapByName(name)
}

// CreateGameZone 创建动态地图
func (c *GameCore) CreateGameZone(lineNum, mapID int32, uid string) *game.GameZone {
	line := c.GetGameLine(lineNum)
	if line == nil {
		return nil
	}
	return line.CreateGameZone(mapID, uid)
}

// SetReady 设置初始化完成状态
func (c *GameCore) SetReady(ready bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.isReady = ready
}

// IsReady 返回游戏是否初始化完成
func (c *GameCore) IsReady() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.isReady
}

// GameLineList 返回所有线路列表（只读）
func (c *GameCore) GameLineList() []*game.GameLine {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result := make([]*game.GameLine, len(c.gameLineList))
	copy(result, c.gameLineList)
	return result
}
