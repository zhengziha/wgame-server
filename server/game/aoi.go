package game

import (
	"strconv"
	"sync"
)

// AOICell 表示 AOI 单元格（九宫格中的一格）
type AOICell struct {
	sessions *sync.Map // map[string]struct{} key=gid
}

// NewAOICell 创建一个新的 AOI 单元格
func NewAOICell() *AOICell {
	return &AOICell{
		sessions: &sync.Map{},
	}
}

// Add 添加玩家到单元格
func (c *AOICell) Add(gid string) {
	c.sessions.Store(gid, struct{}{})
}

// Remove 从单元格移除玩家
func (c *AOICell) Remove(gid string) {
	c.sessions.Delete(gid)
}

// Range 遍历单元格内所有玩家
func (c *AOICell) Range(fn func(gid string) bool) {
	c.sessions.Range(func(key, value interface{}) bool {
		return fn(key.(string))
	})
}

// Count 返回单元格内玩家数量
func (c *AOICell) Count() int {
	count := 0
	c.sessions.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// AOI 表示 AOI 九宫格管理器
// 每个地图有一个 AOI 管理器，负责维护玩家视野范围
type AOI struct {
	mapWidth  int32 // 地图宽度（格子数）
	mapHeight int32 // 地图高度（格子数）
	cellSize  int32 // 每个格子的大小（像素）

	cells *sync.Map // map[string]*AOICell，key 格式 "x,y"

	mu sync.RWMutex
}

// NewAOI 创建一个新的 AOI 管理器
// cellSize 表示每个格子的大小（像素），通常设置为视野范围的一半
func NewAOI(mapWidth, mapHeight, cellSize int32) *AOI {
	return &AOI{
		mapWidth:  mapWidth,
		mapHeight: mapHeight,
		cellSize:  cellSize,
		cells:     &sync.Map{},
	}
}

// getCellKey 获取单元格的键名
func (a *AOI) getCellKey(cellX, cellY int32) string {
	return strconv.Itoa(int(cellX)) + "," + strconv.Itoa(int(cellY))
}

// getCell 获取或创建单元格
func (a *AOI) getCell(cellX, cellY int32) *AOICell {
	key := a.getCellKey(cellX, cellY)
	if cell, ok := a.cells.Load(key); ok {
		return cell.(*AOICell)
	}

	newCell := NewAOICell()
	if existing, loaded := a.cells.LoadOrStore(key, newCell); loaded {
		return existing.(*AOICell)
	}
	return newCell
}

// posToCell 把坐标转换为格子坐标
func (a *AOI) posToCell(x, y int32) (cellX, cellY int32) {
	return x / a.cellSize, y / a.cellSize
}

// Enter 玩家进入 AOI 系统（移动到指定坐标）
// 返回需要通知的玩家 gid 列表（appear 消息）
func (a *AOI) Enter(gid string, x, y int32) []string {
	cellX, cellY := a.posToCell(x, y)

	var newCells []*AOICell
	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			nx := cellX + dx
			ny := cellY + dy
			if nx >= 0 && nx < a.mapWidth && ny >= 0 && ny < a.mapHeight {
				newCells = append(newCells, a.getCell(nx, ny))
			}
		}
	}

	for _, cell := range newCells {
		cell.Add(gid)
	}

	var appearList []string
	seen := make(map[string]bool)
	for _, cell := range newCells {
		cell.Range(func(otherGid string) bool {
			if otherGid != gid && !seen[otherGid] {
				appearList = append(appearList, otherGid)
				seen[otherGid] = true
			}
			return true
		})
	}

	return appearList
}

// Move 玩家在 AOI 系统中移动
// 返回需要 appear 的玩家 gid 列表和需要 disappear 的玩家 gid 列表
func (a *AOI) Move(gid string, oldX, oldY, newX, newY int32) ([]string, []string) {
	oldCellX, oldCellY := a.posToCell(oldX, oldY)
	newCellX, newCellY := a.posToCell(newX, newY)

	if oldCellX == newCellX && oldCellY == newCellY {
		return nil, nil
	}

	var oldCells []*AOICell
	var newCells []*AOICell

	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			nx := oldCellX + dx
			ny := oldCellY + dy
			if nx >= 0 && nx < a.mapWidth && ny >= 0 && ny < a.mapHeight {
				oldCells = append(oldCells, a.getCell(nx, ny))
			}

			nx = newCellX + dx
			ny = newCellY + dy
			if nx >= 0 && nx < a.mapWidth && ny >= 0 && ny < a.mapHeight {
				newCells = append(newCells, a.getCell(nx, ny))
			}
		}
	}

	for _, cell := range oldCells {
		cell.Remove(gid)
	}

	for _, cell := range newCells {
		cell.Add(gid)
	}

	var disappearList []string
	oldSet := make(map[string]bool)
	for _, cell := range oldCells {
		cell.Range(func(otherGid string) bool {
			if otherGid != gid {
				oldSet[otherGid] = true
			}
			return true
		})
	}

	newSet := make(map[string]bool)
	for _, cell := range newCells {
		cell.Range(func(otherGid string) bool {
			if otherGid != gid {
				newSet[otherGid] = true
			}
			return true
		})
	}

	for other := range oldSet {
		if !newSet[other] {
			disappearList = append(disappearList, other)
		}
	}

	var appearList []string
	for other := range newSet {
		if !oldSet[other] {
			appearList = append(appearList, other)
		}
	}

	return appearList, disappearList
}

// Leave 玩家离开 AOI 系统（下线或换图）
// 返回需要 disappear 的玩家 gid 列表
func (a *AOI) Leave(gid string, x, y int32) []string {
	cellX, cellY := a.posToCell(x, y)

	var cells []*AOICell
	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			nx := cellX + dx
			ny := cellY + dy
			if nx >= 0 && nx < a.mapWidth && ny >= 0 && ny < a.mapHeight {
				cells = append(cells, a.getCell(nx, ny))
			}
		}
	}

	var disappearList []string
	seen := make(map[string]bool)
	for _, cell := range cells {
		cell.Range(func(otherGid string) bool {
			if otherGid != gid && !seen[otherGid] {
				disappearList = append(disappearList, otherGid)
				seen[otherGid] = true
			}
			return true
		})
	}

	for _, cell := range cells {
		cell.Remove(gid)
	}

	return disappearList
}

// GetNearby 获取指定坐标附近的玩家 gid 列表（九宫格范围）
func (a *AOI) GetNearby(x, y int32) []string {
	cellX, cellY := a.posToCell(x, y)

	var result []string
	seen := make(map[string]bool)

	for dx := int32(-1); dx <= 1; dx++ {
		for dy := int32(-1); dy <= 1; dy++ {
			nx := cellX + dx
			ny := cellY + dy
			if nx >= 0 && nx < a.mapWidth && ny >= 0 && ny < a.mapHeight {
				a.getCell(nx, ny).Range(func(gid string) bool {
					if !seen[gid] {
						result = append(result, gid)
						seen[gid] = true
					}
					return true
				})
			}
		}
	}

	return result
}
