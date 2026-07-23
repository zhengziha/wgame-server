package game

import (
	"sync"
	"testing"
)

// newTestCharaManager 创建一个新的测试用 CharaManager
func newTestCharaManager() *CharaManager {
	return &CharaManager{
		charaMap:   sync.Map{},
		charaIdMap: sync.Map{},
	}
}

// TestCharaManagerAddAndGet 测试添加和获取角色
func TestCharaManagerAddAndGet(t *testing.T) {
	mgr := newTestCharaManager()

	// 添加角色
	chara := &Chara{ID: 1, Gid: "gid1", Name: "test1"}
	mgr.AddChara(chara)

	// 按 Gid 获取
	got := mgr.GetCharaByGid("gid1")
	if got == nil || got.ID != 1 {
		t.Errorf("GetCharaByGid(gid1): got %v, want chara with ID=1", got)
	}

	// 按 ID 获取
	got = mgr.GetCharaById(1)
	if got == nil || got.Gid != "gid1" {
		t.Errorf("GetCharaById(1): got %v, want chara with Gid=gid1", got)
	}

	// 获取不存在的角色
	if mgr.GetCharaByGid("nonexistent") != nil {
		t.Error("GetCharaByGid(nonexistent) should return nil")
	}
	if mgr.GetCharaById(999) != nil {
		t.Error("GetCharaById(999) should return nil")
	}
}

// TestCharaManagerRemove 测试移除角色
func TestCharaManagerRemove(t *testing.T) {
	mgr := newTestCharaManager()

	// 添加角色
	chara := &Chara{ID: 2, Gid: "gid2", Name: "test2"}
	mgr.AddChara(chara)

	// 验证添加成功
	if mgr.GetCharaByGid("gid2") == nil {
		t.Error("Chara should exist after AddChara")
	}

	// 移除角色
	mgr.RemoveChara(chara)

	// 验证移除成功
	if mgr.GetCharaByGid("gid2") != nil {
		t.Error("Chara should not exist after RemoveChara")
	}
	if mgr.GetCharaById(2) != nil {
		t.Error("Chara should not exist after RemoveChara")
	}
}

// TestCharaManagerNilChara 测试 nil 角色处理
func TestCharaManagerNilChara(t *testing.T) {
	mgr := newTestCharaManager()

	// 添加 nil 角色
	mgr.AddChara(nil)

	// 移除 nil 角色
	mgr.RemoveChara(nil)

	// 获取空 gid
	if mgr.GetCharaByGid("") != nil {
		t.Error("GetCharaByGid(\"\") should return nil")
	}

	// 获取 0 ID
	if mgr.GetCharaById(0) != nil {
		t.Error("GetCharaById(0) should return nil")
	}
}

// TestCharaManagerCount 测试角色数量
func TestCharaManagerCount(t *testing.T) {
	mgr := newTestCharaManager()

	// 先清空
	mgr.Range(func(chara *Chara) bool {
		mgr.RemoveChara(chara)
		return true
	})

	if mgr.Count() != 0 {
		t.Errorf("Count should be 0, got %d", mgr.Count())
	}

	// 添加角色
	mgr.AddChara(&Chara{ID: 100, Gid: "gid100"})
	mgr.AddChara(&Chara{ID: 101, Gid: "gid101"})

	if mgr.Count() != 2 {
		t.Errorf("Count should be 2, got %d", mgr.Count())
	}
}

// TestCharaManagerRange 测试遍历功能
func TestCharaManagerRange(t *testing.T) {
	mgr := newTestCharaManager()

	// 添加角色
	mgr.AddChara(&Chara{ID: 200, Gid: "gid200", Name: "char200"})
	mgr.AddChara(&Chara{ID: 201, Gid: "gid201", Name: "char201"})
	mgr.AddChara(&Chara{ID: 202, Gid: "gid202", Name: "char202"})

	var count int
	var names []string
	mgr.Range(func(chara *Chara) bool {
		count++
		names = append(names, chara.Name)
		return true
	})

	if count != 3 {
		t.Errorf("Range count: got %d, want 3", count)
	}

	// 测试提前退出
	count = 0
	mgr.Range(func(chara *Chara) bool {
		count++
		return count < 2 // 只遍历前2个
	})

	if count != 2 {
		t.Errorf("Range with early exit: got %d, want 2", count)
	}
}

// TestCharaManagerConcurrent 测试并发安全性
func TestCharaManagerConcurrent(t *testing.T) {
	mgr := newTestCharaManager()
	var wg sync.WaitGroup
	numGoroutines := 100

	// 并发添加角色
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			chara := &Chara{
				ID:   int32(id + 1000),
				Gid:  "gid_concurrent_" + string(rune(id+'0')),
				Name: "concurrent_" + string(rune(id+'0')),
			}
			mgr.AddChara(chara)
		}(i)
	}
	wg.Wait()

	// 验证添加数量
	if mgr.Count() != numGoroutines {
		t.Errorf("Concurrent add count: got %d, want %d", mgr.Count(), numGoroutines)
	}

	// 并发读取和修改
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			gid := "gid_concurrent_" + string(rune(id+'0'))
			chara := mgr.GetCharaByGid(gid)
			if chara != nil {
				chara.Name = "modified_" + string(rune(id+'0'))
			}
		}(i)

		go func(id int) {
			defer wg.Done()
			mgr.GetCharaById(int32(id + 1000))
		}(i)
	}
	wg.Wait()

	// 并发移除角色
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			chara := mgr.GetCharaById(int32(id + 1000))
			if chara != nil {
				mgr.RemoveChara(chara)
			}
		}(i)
	}
	wg.Wait()

	// 验证移除后数量为0
	if mgr.Count() != 0 {
		t.Errorf("Concurrent remove count: got %d, want 0", mgr.Count())
	}
}
