package game

import "wgame-server/comm/log"

// PlayerEnterMap 处理玩家进入地图的核心逻辑
// 返回需要收到 appear 消息的玩家 gid 列表
func PlayerEnterMap(chara *Chara, gameMap *GameMap) []string {
	if chara == nil || gameMap == nil {
		return nil
	}

	log.Info("[game] PlayerEnterMap chara=%s map=%d x=%d y=%d", chara.Name, gameMap.ID, chara.X, chara.Y)

	// 检查是否已经在地图中（断线重连场景）
	if gameMap.Has(chara.Gid) {
		log.Info("[game] PlayerEnterMap: chara %s already in map %d", chara.Name, gameMap.ID)
		return nil
	}

	gameMap.Join(chara.Gid)
	appearList := gameMap.AOI.Enter(chara.Gid, chara.X, chara.Y)

	log.Info("[game] PlayerEnterMap appearList size=%d", len(appearList))
	return appearList
}

// PlayerLeaveMap 处理玩家离开地图的核心逻辑
// 返回需要收到 disappear 消息的玩家 gid 列表
func PlayerLeaveMap(chara *Chara, gameMap *GameMap) []string {
	if chara == nil || gameMap == nil {
		return nil
	}

	disappearList := gameMap.AOI.Leave(chara.Gid, chara.X, chara.Y)
	gameMap.Leave(chara.Gid)

	log.Info("[game] PlayerLeaveMap chara=%s map=%d disappearList size=%d", chara.Name, gameMap.ID, len(disappearList))
	return disappearList
}

// PlayerMove 处理玩家移动的核心逻辑
// 返回需要 appear 和 disappear 的玩家 gid 列表
func PlayerMove(chara *Chara, gameMap *GameMap, newX, newY int32) ([]string, []string) {
	if chara == nil || gameMap == nil {
		return nil, nil
	}

	oldX, oldY := chara.X, chara.Y
	if oldX == newX && oldY == newY {
		return nil, nil
	}

	appearList, disappearList := gameMap.AOI.Move(chara.Gid, oldX, oldY, newX, newY)

	chara.X = newX
	chara.Y = newY

	return appearList, disappearList
}
