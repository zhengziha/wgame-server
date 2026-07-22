package map_handler

import (
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/core"
	"wgame-server/server/game"
	map_msg "wgame-server/server/msg/map"
	"wgame-server/server/network/broadcaster"
	"wgame-server/server/network/handler"
)

// CmdTeleportHandler 处理 CMD_TELEPORT (cmd=32768)
// 传送（地图切换）
func CmdTeleportHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	mapID, _ := reader.ReadInt()
	x, _ := reader.ReadInt()
	y, _ := reader.ReadInt()
	isTaskWalk, _ := reader.ReadUByte()

	log.Info("[map] CMD_TELEPORT map=%d x=%d y=%d isTaskWalk=%d", mapID, x, y, isTaskWalk)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[map] chara not found gid=%s", gid)
		return nil
	}

	gameMap := core.Instance().GetGameMap(chara.Line, mapID)
	if gameMap == nil {
		log.Error("[map] map not found line=%d map=%d", chara.Line, mapID)
		return nil
	}

	if chara.MapId == mapID && chara.X == x && chara.Y == y {
		return nil
	}

	EnterMap(ctx, chara, gameMap, x, y)

	return nil
}

// CmdEnterRoomHandler 处理 CMD_ENTER_ROOM (cmd=4144)
// 通过地图名称进入地图
func CmdEnterRoomHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	roomName, _ := reader.ReadString()
	isTaskWalk, _ := reader.ReadUByte()

	log.Info("[map] CMD_ENTER_ROOM room_name=%s isTaskWalk=%d", roomName, isTaskWalk)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[map] chara not found gid=%s", gid)
		return nil
	}

	if chara.MapName == roomName {
		return nil
	}

	gameMap := core.Instance().GetGameMapByName(chara.Line, roomName)
	if gameMap == nil {
		log.Error("[map] map not found line=%d name=%s", chara.Line, roomName)
		return nil
	}

	x := gameMap.X
	y := gameMap.Y

	EnterMap(ctx, chara, gameMap, x, y)

	return nil
}

// EnterMap 处理玩家进入地图的核心逻辑
// 包括：离开旧地图、进入新地图、AOI同步、周围玩家通知
// 发包顺序：MSG_CLEAR_ALL_CHAR -> MSG_ENTER_ROOM_79 -> MSG_APPEAR
func EnterMap(ctx context.MyCmdContext, chara *game.Chara, gameMap *game.GameMap, x, y int32) {
	log.Info("[map] EnterMap chara=%s from_map=%d to_map=%d x=%d y=%d",
		chara.Name, chara.MapId, gameMap.ID, x, y)

	// 1. 离开旧地图（通知旧地图玩家该玩家消失）
	if chara.MapId > 0 && chara.MapId != gameMap.ID {
		oldMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
		if oldMap != nil {
			disappearList := game.PlayerLeaveMap(chara, oldMap)
			// 通知旧地图玩家该玩家消失
			sendDisappearToNearbyPlayers(chara, disappearList)
		}
	}

	// 2. 更新角色位置
	chara.MapId = gameMap.ID
	chara.MapName = gameMap.Name
	chara.X = x
	chara.Y = y

	// 3. 进入新地图，获取需要通知的周围玩家
	appearList := game.PlayerEnterMap(chara, gameMap)

	// 4. 发送消息给自己
	sendEnterMapMessages(ctx, chara, gameMap)

	// 5. 发送消息给周围玩家（他们看到新玩家出现）
	sendAppearToNearbyPlayers(chara, appearList)

	log.Info("[map] EnterMap done chara=%s map=%d", chara.Name, gameMap.ID)
}

// sendEnterMapMessages 发送玩家进入地图后的消息（发送给自己）
func sendEnterMapMessages(ctx context.MyCmdContext, chara *game.Chara, gameMap *game.GameMap) {
	// 1) MSG_CLEAR_ALL_CHAR - 清除所有角色
	ctx.Write(&map_msg.MsgClearAllChar{
		ID:    chara.ID,
		MapID: gameMap.ID,
	})

	// 2) MSG_ENTER_ROOM_79 - 进入房间
	ctx.Write(&map_msg.MsgEnterRoom79{
		MapName:          gameMap.Name,
		MapShowName:      gameMap.Name,
		MapID:            gameMap.ID,
		X:                chara.X,
		Y:                chara.Y,
		Dir:              chara.Dir,
		MapIndex:         0,
		CompactMapIndex:  0,
		FloorIndex:       0,
		WallIndex:        0,
		SafeZone:         0,
		IsTaskWalk:       0,
		EnterEffectIndex: 0,
	})

	// 3) MSG_APPEAR - 发送自己
	ctx.Write(&map_msg.MsgAppear{
		CharID:      chara.ID,
		Name:        chara.Name,
		Gid:         chara.Gid,
		Level:       chara.Level,
		Polar:       chara.Polar,
		Sex:         chara.Sex,
		X:           chara.X,
		Y:           chara.Y,
		Dir:         chara.Dir,
		Waiguan:     chara.Waiguan,
		Nice:        chara.Nice,
		FashionIcon: 0,
	})

	// 发送周围玩家列表
	if gameMap.AOI != nil {
		nearbyGids := gameMap.AOI.GetNearby(chara.X, chara.Y)
		for _, nearbyGid := range nearbyGids {
			nearbyChara := game.CharaManagerInstance().GetCharaByGid(nearbyGid)
			if nearbyChara != nil && nearbyChara.ID != chara.ID {
				ctx.Write(&map_msg.MsgAppear{
					CharID:      nearbyChara.ID,
					Name:        nearbyChara.Name,
					Gid:         nearbyChara.Gid,
					Level:       nearbyChara.Level,
					Polar:       nearbyChara.Polar,
					Sex:         nearbyChara.Sex,
					X:           nearbyChara.X,
					Y:           nearbyChara.Y,
					Dir:         nearbyChara.Dir,
					Waiguan:     nearbyChara.Waiguan,
					Nice:        nearbyChara.Nice,
					FashionIcon: 0,
				})
			}
		}
	}
}

// sendAppearToNearbyPlayers 通知周围玩家有新玩家出现
func sendAppearToNearbyPlayers(chara *game.Chara, appearList []string) {
	for _, otherGid := range appearList {
		otherChara := game.CharaManagerInstance().GetCharaByGid(otherGid)
		if otherChara != nil {
			log.Info("[map] notify %s about appear of %s", otherChara.Name, chara.Name)
			// 通过 broadcaster 发送消息给周围玩家
			broadcaster.SendToGid(&map_msg.MsgAppear{
				CharID:      chara.ID,
				Name:        chara.Name,
				Gid:         chara.Gid,
				Level:       chara.Level,
				Polar:       chara.Polar,
				Sex:         chara.Sex,
				X:           chara.X,
				Y:           chara.Y,
				Dir:         chara.Dir,
				Waiguan:     chara.Waiguan,
				Nice:        chara.Nice,
				FashionIcon: 0,
			}, otherGid)
		}
	}
}

// sendDisappearToNearbyPlayers 通知周围玩家有玩家消失
func sendDisappearToNearbyPlayers(chara *game.Chara, disappearList []string) {
	for _, otherGid := range disappearList {
		otherChara := game.CharaManagerInstance().GetCharaByGid(otherGid)
		if otherChara != nil {
			log.Info("[map] notify %s about disappear of %s", otherChara.Name, chara.Name)
			// 通过 broadcaster 发送消息给周围玩家
			broadcaster.SendToGid(&map_msg.MsgDisappear{
				CharID: chara.ID,
			}, otherGid)
		}
	}
}

func init() {
	handler.Register(32768, "CmdTeleport", CmdTeleportHandler)
	handler.Register(4144, "CmdEnterRoom", CmdEnterRoomHandler)
}
