package game

import (
	"sync"
)

// CharaManager 角色管理器
// 管理在线角色的生命周期
type CharaManager struct {
	charaMap   sync.Map // map[string]*Chara key=Gid
	charaIdMap sync.Map // map[int32]*Chara key=ID
}

var (
	charaManagerInstance *CharaManager
	charaManagerOnce     sync.Once
)

// Instance 返回 CharaManager 单例
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
