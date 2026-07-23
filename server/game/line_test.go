package game

import (
	"testing"

	"wgame-server/server/model"
)

// TestGameLineInit 测试线路初始化
func TestGameLineInit(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "揽仙镇", MapID: 1000, X: 22, Y: 108},
		{ID: 2, Name: "天墉城", MapID: 1001, X: 100, Y: 200},
		{ID: 3, Name: "北海沙滩", MapID: 1002, X: 50, Y: 50},
	}

	line.Init(mapInfos)

	if len(line.MapList()) != 3 {
		t.Errorf("MapList length: got %d, want 3", len(line.MapList()))
	}
}

// TestGameLineGetGameMapById 测试按地图编号查找（验证 O(1) map 查找）
func TestGameLineGetGameMapById(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "揽仙镇", MapID: 1000, X: 22, Y: 108},
		{ID: 2, Name: "天墉城", MapID: 1001, X: 100, Y: 200},
		{ID: 3, Name: "北海沙滩", MapID: 1002, X: 50, Y: 50},
	}

	line.Init(mapInfos)

	// 测试存在的地图（使用 MapID 作为 key）
	gm := line.GetGameMapById(1000)
	if gm == nil {
		t.Error("GetGameMapById(1000) should return a map")
	}
	if gm.Name != "揽仙镇" {
		t.Errorf("Map name: got %s, want 揽仙镇", gm.Name)
	}

	gm = line.GetGameMapById(1001)
	if gm == nil {
		t.Error("GetGameMapById(1001) should return a map")
	}
	if gm.Name != "天墉城" {
		t.Errorf("Map name: got %s, want 天墉城", gm.Name)
	}

	// 测试不存在的地图
	gm = line.GetGameMapById(9999)
	if gm != nil {
		t.Error("GetGameMapById(9999) should return nil")
	}
}

// TestGameLineGetGameMapByName 测试按地图名称查找
func TestGameLineGetGameMapByName(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "揽仙镇", MapID: 1000, X: 22, Y: 108},
		{ID: 2, Name: "天墉城", MapID: 1001, X: 100, Y: 200},
	}

	line.Init(mapInfos)

	// 测试存在的地图
	gm := line.GetGameMapByName("揽仙镇")
	if gm == nil {
		t.Error("GetGameMapByName(揽仙镇) should return a map")
	}
	if gm.ID != 1000 { // MapID 是地图编号
		t.Errorf("Map ID: got %d, want 1000", gm.ID)
	}

	// 测试不存在的地图
	gm = line.GetGameMapByName("不存在的地图")
	if gm != nil {
		t.Error("GetGameMapByName(不存在的地图) should return nil")
	}
}

// TestGameLineCreateGameZone 测试创建动态地图
func TestGameLineCreateGameZone(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "副本地图", MapID: 2000, X: 0, Y: 0},
	}

	line.Init(mapInfos)

	// 创建动态地图（使用 MapID 2000）
	zone := line.CreateGameZone(2000, "zone1")
	if zone == nil {
		t.Error("CreateGameZone should return a zone")
	}
	if zone.Uid != "zone1" {
		t.Errorf("Zone Uid: got %s, want zone1", zone.Uid)
	}
	if zone.BaseMap == nil {
		t.Error("Zone BaseMap should not be nil")
	}
	if zone.BaseMap.ID != 2000 { // BaseMap.ID 是地图编号（MapID）
		t.Errorf("Zone BaseMap ID: got %d, want 2000", zone.BaseMap.ID)
	}

	// 创建已存在的动态地图（应该返回已有的）
	zone2 := line.CreateGameZone(2000, "zone1")
	if zone2 != zone {
		t.Error("CreateGameZone with existing uid should return existing zone")
	}

	// 创建不存在的基础地图（使用不存在的 MapID）
	zone3 := line.CreateGameZone(999, "zone2")
	if zone3 != nil {
		t.Error("CreateGameZone with non-existent base map should return nil")
	}
}

// TestGameLineDeleteGameZone 测试删除动态地图
func TestGameLineDeleteGameZone(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "副本地图", MapID: 2000, X: 0, Y: 0},
	}

	line.Init(mapInfos)

	// 创建动态地图（使用 MapID 2000）
	zone := line.CreateGameZone(2000, "zone1")
	if zone == nil {
		t.Fatal("CreateGameZone should return a zone")
	}

	// 验证创建成功
	if line.GetGameZone("zone1") == nil {
		t.Error("GetGameZone should find the zone")
	}

	// 删除动态地图
	line.DeleteGameZone("zone1")

	// 验证删除成功
	if line.GetGameZone("zone1") != nil {
		t.Error("GetGameZone should not find the zone after delete")
	}

	// 删除不存在的动态地图（不应 panic）
	line.DeleteGameZone("nonexistent")
}

// TestGameLineGetGameZone 测试获取动态地图
func TestGameLineGetGameZone(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "副本地图", MapID: 2000, X: 0, Y: 0},
	}

	line.Init(mapInfos)

	// 获取不存在的动态地图
	if line.GetGameZone("nonexistent") != nil {
		t.Error("GetGameZone(nonexistent) should return nil")
	}

	// 创建并获取（使用 MapID 2000）
	line.CreateGameZone(2000, "zone1")
	zone := line.GetGameZone("zone1")
	if zone == nil {
		t.Error("GetGameZone(zone1) should return the zone")
	}
	if zone.Uid != "zone1" {
		t.Errorf("Zone Uid: got %s, want zone1", zone.Uid)
	}
}

// TestGameLineMapList 测试返回地图列表
func TestGameLineMapList(t *testing.T) {
	line := NewGameLine(1)

	mapInfos := []*model.MapInfo{
		{ID: 1, Name: "地图1", MapID: 1000, X: 0, Y: 0},
		{ID: 2, Name: "地图2", MapID: 1001, X: 0, Y: 0},
		{ID: 3, Name: "地图3", MapID: 1002, X: 0, Y: 0},
	}

	line.Init(mapInfos)

	list := line.MapList()
	if len(list) != 3 {
		t.Errorf("MapList length: got %d, want 3", len(list))
	}

	// 验证返回的是切片副本（修改切片本身不影响原数据）
	list = append(list, nil)
	originalLen := len(line.MapList())
	if originalLen != 3 {
		t.Errorf("Original MapList length should remain 3 after append to copy, got %d", originalLen)
	}
}
