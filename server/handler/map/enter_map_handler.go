package map_handler

import (
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/core"
	"wgame-server/server/game"
	map_msg "wgame-server/server/msg/map"
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
func EnterMap(ctx context.MyCmdContext, chara *game.Chara, gameMap *game.GameMap, x, y int32) {
	log.Info("[map] EnterMap chara=%s from_map=%d to_map=%d x=%d y=%d",
		chara.Name, chara.MapId, gameMap.ID, x, y)

	if chara.MapId > 0 {
		oldMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
		if oldMap != nil {
			game.PlayerLeaveMap(chara, oldMap)
		}
	}

	chara.MapId = gameMap.ID
	chara.MapName = gameMap.Name
	chara.X = x
	chara.Y = y

	game.PlayerEnterMap(chara, gameMap)

	ctx.Write(&map_msg.MsgMapInfo{
		MapID:   gameMap.ID,
		MapName: gameMap.Name,
	})

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

	log.Info("[map] EnterMap done chara=%s map=%d", chara.Name, gameMap.ID)
}

func init() {
	handler.Register(32768, "CmdTeleport", CmdTeleportHandler)
	handler.Register(4144, "CmdEnterRoom", CmdEnterRoomHandler)
}
