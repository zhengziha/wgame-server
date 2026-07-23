package game

import (
	"sync"
)

// CharaManager 角色管理器
// 管理在线角色的生命周期
//
// Java 对比：
//   - sync.Map 类似于 ConcurrentHashMap，并发安全的 map
//   - sync.Once 类似于使用 volatile + synchronized 实现的单例模式（双重检查锁定）
//   - Store/Load/Delete 对应 ConcurrentHashMap 的 put/get/remove
type CharaManager struct {
	charaMap   sync.Map // map[string]*Chara key=Gid（并发安全，无需额外加锁）
	charaIdMap sync.Map // map[int32]*Chara key=ID
}

var (
	charaManagerInstance *CharaManager
	charaManagerOnce     sync.Once // 确保单例只初始化一次，线程安全
)

// Instance 返回 CharaManager 单例
// sync.Once.Do 确保初始化逻辑只执行一次（类似 Java 的双重检查锁定单例）
func CharaManagerInstance() *CharaManager {
	charaManagerOnce.Do(func() {
		charaManagerInstance = &CharaManager{
			charaMap:   sync.Map{},
			charaIdMap: sync.Map{},
		}
	})
	return charaManagerInstance
}

// AddChara 添加角色到管理器
func (m *CharaManager) AddChara(chara *Chara) {
	if chara == nil {
		return
	}
	m.charaMap.Store(chara.Gid, chara)
	m.charaIdMap.Store(chara.ID, chara)
}

// RemoveChara 从管理器移除角色
func (m *CharaManager) RemoveChara(chara *Chara) {
	if chara == nil {
		return
	}
	m.charaMap.Delete(chara.Gid)
	m.charaIdMap.Delete(chara.ID)
}

// GetCharaByGid 按 Gid 获取角色
func (m *CharaManager) GetCharaByGid(gid string) *Chara {
	if gid == "" {
		return nil
	}
	if v, ok := m.charaMap.Load(gid); ok {
		return v.(*Chara)
	}
	return nil
}

// GetCharaById 按 ID 获取角色
func (m *CharaManager) GetCharaById(id int32) *Chara {
	if id <= 0 {
		return nil
	}
	if v, ok := m.charaIdMap.Load(id); ok {
		return v.(*Chara)
	}
	return nil
}

// Range 遍历所有角色
func (m *CharaManager) Range(fn func(*Chara) bool) {
	m.charaMap.Range(func(key, value interface{}) bool {
		return fn(value.(*Chara))
	})
}

// Count 返回角色数量
func (m *CharaManager) Count() int {
	count := 0
	m.charaMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
