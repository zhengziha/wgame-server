package game

import (
	"testing"
)

// TestAOIEnter 测试玩家进入 AOI 系统
func TestAOIEnter(t *testing.T) {
	aoi := NewAOI(20, 20, 30) // 20x20格子，每格30像素

	// 玩家A进入
	appearA := aoi.Enter("playerA", 40, 40)
	if len(appearA) != 0 {
		t.Fatalf("first player should have no appear list, got %d", len(appearA))
	}

	// 玩家B进入同一九宫格区域
	appearB := aoi.Enter("playerB", 50, 50)
	if len(appearB) != 1 || appearB[0] != "playerA" {
		t.Fatalf("playerB should see playerA, got %v", appearB)
	}

	// 玩家C进入不同区域
	appearC := aoi.Enter("playerC", 200, 200)
	if len(appearC) != 0 {
		t.Fatalf("playerC in different area should have no appear list, got %v", appearC)
	}
}

// TestAOIMove 测试玩家移动
func TestAOIMove(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	// 放置两个玩家在不同格子
	aoi.Enter("playerA", 40, 40)   // 格子(1,1)
	aoi.Enter("playerB", 150, 150) // 格子(5,5)

	// 玩家A移动到玩家B附近
	appear, disappear := aoi.Move("playerA", 40, 40, 160, 160)
	if len(appear) != 1 || appear[0] != "playerB" {
		t.Fatalf("playerA should appear to playerB, got appear=%v", appear)
	}
	if len(disappear) != 0 {
		t.Fatalf("no one should disappear, got %v", disappear)
	}
}

// TestAOIMoveSameCell 测试玩家在同一个格子内移动
func TestAOIMoveSameCell(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	aoi.Enter("playerA", 40, 40)

	// 在同一格子内移动（30像素范围内）
	appear, disappear := aoi.Move("playerA", 40, 40, 50, 50)
	if len(appear) != 0 || len(disappear) != 0 {
		t.Fatalf("move within same cell should have empty lists, got appear=%v disappear=%v", appear, disappear)
	}
}

// TestAOILeave 测试玩家离开
func TestAOILeave(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	aoi.Enter("playerA", 40, 40)
	aoi.Enter("playerB", 50, 50)

	// 玩家A离开，玩家B应该收到disappear
	disappear := aoi.Leave("playerA", 40, 40)
	if len(disappear) != 1 || disappear[0] != "playerB" {
		t.Fatalf("playerB should see playerA disappear, got %v", disappear)
	}

	// 验证玩家A已被移除
	nearby := aoi.GetNearby(50, 50)
	for _, gid := range nearby {
		if gid == "playerA" {
			t.Fatal("playerA should not be in nearby list after leave")
		}
	}
}

// TestAOIGetNearby 测试获取附近玩家
func TestAOIGetNearby(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	aoi.Enter("playerA", 40, 40)
	aoi.Enter("playerB", 50, 50)
	aoi.Enter("playerC", 200, 200)

	// 在玩家A附近查询
	nearby := aoi.GetNearby(45, 45)
	if len(nearby) != 2 {
		t.Fatalf("should find 2 players (A and B), got %d: %v", len(nearby), nearby)
	}
}

// TestAOIGetCellKeyLargeCoordinate 测试坐标 >= 10 的场景（修复后的getCellKey）
func TestAOIGetCellKeyLargeCoordinate(t *testing.T) {
	aoi := NewAOI(100, 100, 30)

	// 玩家在大坐标位置（格子坐标 >= 10）
	aoi.Enter("playerA", 350, 350) // 格子(11,11)
	aoi.Enter("playerB", 360, 360) // 格子(12,12)

	// 验证两个玩家能互相看到
	appear := aoi.Enter("playerC", 355, 355) // 格子(11,11)
	if len(appear) != 2 {
		t.Fatalf("playerC should see playerA and playerB, got %v", appear)
	}

	// 验证附近查询
	nearby := aoi.GetNearby(350, 350)
	if len(nearby) != 3 {
		t.Fatalf("should find 3 players, got %d: %v", len(nearby), nearby)
	}
}

// TestAOIMoveCrossBoundary 测试玩家跨越九宫格边界
func TestAOIMoveCrossBoundary(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	// 放置三个玩家形成三角形
	aoi.Enter("playerA", 10, 10)   // 格子(0,0)
	aoi.Enter("playerB", 200, 10)  // 格子(6,0)
	aoi.Enter("playerC", 100, 200) // 格子(3,6)

	// 玩家A从格子(0,0)移动到格子(10,10)，跨越多个九宫格
	appear, disappear := aoi.Move("playerA", 10, 10, 320, 320)

	// 玩家A离开格子(0,0)区域，应该没人看到（只有自己）
	if len(disappear) != 0 {
		t.Fatalf("no one should see playerA disappear, got %v", disappear)
	}

	// 玩家A进入格子(10,10)区域，应该看到玩家C（如果在视野内）
	// 格子(10,10)的九宫格范围是(9-11, 9-11)
	// 玩家C在格子(3,6)，不在视野范围内
	if len(appear) != 0 {
		t.Fatalf("playerA should not appear to anyone, got %v", appear)
	}
}

// TestAOIMoveOverlapCells 测试移动时重叠格子场景
func TestAOIMoveOverlapCells(t *testing.T) {
	aoi := NewAOI(20, 20, 30)

	// playerA在格子(1,1)
	aoi.Enter("playerA", 40, 40)

	// playerB在格子(2,2)，与playerA的九宫格重叠
	aoi.Enter("playerB", 70, 70)

	// playerC在格子(4,4)，与playerA的旧九宫格不重叠，但在新九宫格内
	aoi.Enter("playerC", 130, 130)

	// playerA从(1,1)移动到(3,3)
	// 旧九宫格: (0-2, 0-2)
	// 新九宫格: (2-4, 2-4)
	// playerB在格子(2,2)，始终在九宫格内，不应该收到消息
	// playerC在格子(4,4)，在新九宫格内但不在旧九宫格内，应该收到appear
	appear, disappear := aoi.Move("playerA", 40, 40, 100, 100)

	if len(disappear) != 0 {
		t.Fatalf("no one should disappear, got %v", disappear)
	}
	if len(appear) != 1 || appear[0] != "playerC" {
		t.Fatalf("playerC should appear, got %v", appear)
	}
}
