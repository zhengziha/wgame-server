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

// CmdMultiMoveToHandler 处理 CMD_MULTI_MOVE_TO (cmd=61634)
// 自动任务寻路/移动
func CmdMultiMoveToHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	id, _ := reader.ReadInt()
	mapID, _ := reader.ReadInt()
	_, _ = reader.ReadInt() // map_index
	count, _ := reader.ReadShort()

	var x, y int32
	for i := int16(0); i < count; i++ {
		curX, _ := reader.ReadShort()
		curY, _ := reader.ReadShort()
		x = int32(curX)
		y = int32(curY)
	}

	dir, _ := reader.ReadShort()
	_, _ = reader.ReadInt() // send_time

	log.Info("[move] CMD_MULTI_MOVE_TO id=%d map=%d x=%d y=%d dir=%d", id, mapID, x, y, dir)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[move] chara not found gid=%s", gid)
		return nil
	}

	if mapID != chara.MapId {
		log.Error("[move] map mismatch chara=%s expected=%d actual=%d", chara.Name, chara.MapId, mapID)
		return nil
	}

	gameMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
	if gameMap == nil {
		return nil
	}

	chara.Dir = int32(dir)

	appearList, disappearList := game.PlayerMove(chara, gameMap, x, y)

	log.Info("[move] appear=%d disappear=%d", len(appearList), len(disappearList))

	// 发送移动消息给自己
	ctx.Write(&map_msg.MsgMoved{
		ID:  chara.ID,
		X:   int16(x),
		Y:   int16(y),
		Dir: int8(dir),
	})

	// 发送移动消息给视野内的其他玩家（排除新进入视野的玩家）
	sendMoveMessages(chara, gameMap, gid, appearList)

	// 通知新进入视野的玩家
	sendAppearMessages(chara, appearList)

	// 通知离开视野的玩家
	sendDisappearMessages(chara.ID, disappearList)

	return nil
}

// CmdOtherMoveToHandler 处理 CMD_OTHER_MOVE_TO (cmd=16558)
// 队伍成员移动
func CmdOtherMoveToHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	id, _ := reader.ReadInt()
	mapID, _ := reader.ReadInt()
	x, _ := reader.ReadShort()
	y, _ := reader.ReadShort()
	dir, _ := reader.ReadShort()

	log.Info("[move] CMD_OTHER_MOVE_TO id=%d map=%d x=%d y=%d dir=%d", id, mapID, x, y, dir)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[move] chara not found gid=%s", gid)
		return nil
	}

	if mapID != chara.MapId {
		return nil
	}

	gameMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
	if gameMap == nil {
		return nil
	}

	chara.Dir = int32(dir)

	appearList, disappearList := game.PlayerMove(chara, gameMap, int32(x), int32(y))

	// 发送移动消息给自己
	ctx.Write(&map_msg.MsgMoved{
		ID:  chara.ID,
		X:   int16(chara.X),
		Y:   int16(chara.Y),
		Dir: int8(chara.Dir),
	})

	// 发送移动消息给视野内的其他玩家（排除新进入视野的玩家）
	sendMoveMessages(chara, gameMap, gid, appearList)

	// 通知新进入视野的玩家
	sendAppearMessages(chara, appearList)

	// 通知离开视野的玩家
	sendDisappearMessages(chara.ID, disappearList)

	return nil
}

func init() {
	handler.Register(61634, "CmdMultiMoveTo", CmdMultiMoveToHandler)
	handler.Register(16558, "CmdOtherMoveTo", CmdOtherMoveToHandler)
}

// sendMoveMessages 发送移动消息给视野内的其他玩家（排除自己和新进入视野的玩家）
func sendMoveMessages(chara *game.Chara, gameMap *game.GameMap, gid string, appearList []string) {
	if gameMap.AOI == nil {
		return
	}

	// 将 appearList 转为 map 方便快速查找
	appearSet := make(map[string]bool)
	for _, g := range appearList {
		appearSet[g] = true
	}

	nearbyGids := gameMap.AOI.GetNearby(chara.X, chara.Y)
	for _, otherGid := range nearbyGids {
		if otherGid != gid && !appearSet[otherGid] { // 不发给自己，也不发给新进入视野的玩家
			broadcaster.SendToGid(&map_msg.MsgMoved{
				ID:  chara.ID,
				X:   int16(chara.X),
				Y:   int16(chara.Y),
				Dir: int8(chara.Dir),
			}, otherGid)
		}
	}
}

// sendAppearMessages 通知新进入视野的玩家
func sendAppearMessages(chara *game.Chara, appearList []string) {
	for _, otherGid := range appearList {
		otherChara := game.CharaManagerInstance().GetCharaByGid(otherGid)
		if otherChara != nil {
			log.Info("[move] notify %s about appear of %s", otherChara.Name, chara.Name)
			broadcaster.SendToGid(&map_msg.MsgAppear{
				ID:                 chara.ID,
				X:                  int16(chara.X),
				Y:                  int16(chara.Y),
				Dir:                int16(chara.Dir),
				Icon:               0,
				WeaponIcon:         0,
				Type:               int16(chara.Sex),
				SubType:            0,
				OwnerID:            0,
				LeaderID:           0,
				Name:               chara.Name,
				Level:              int16(chara.Level),
				Title:              "",
				Family:             "",
				Party:              "",
				Status:             0,
				SpecialIcon:        0,
				OrgIcon:            0,
				SuitIcon:           0,
				SuitLight:          0,
				GuardIcon:          0,
				PetIcon:            0,
				ShadowIcon:         0,
				ShelterIcon:        0,
				MountIcon:          0,
				AliName:            "",
				Gid:                chara.Gid,
				Camp:               "",
				VipType:            0,
				IsHide:             0,
				MoveSpeed:          0,
				Score:              0,
				Opacity:            0,
				Masquerade:         0,
				UpgradeState:       0,
				UpgradeType:        0,
				Obstacle:           0,
				EffectCount:        0,
				Effects:            []int32{},
				ShareMountIcon:     0,
				ShareMountLeaderID: 0,
				ShareMountShadow:   0,
				GatherCount:        0,
				GatherMountIcons:   []int32{},
				GatherNameNum:      0,
				GatherNames:        []string{},
				Portrait:           0,
				CustomIcon:         "",
				TeamIcon:           0,
				ExtraScale:         0,
				GatherSuitIcons:    0,
				BanRule:            "",
				TitleBanRule:       "",
				XOffset:            0,
				YOffset:            0,
				MoveType:           0,
				FlyType:            0,
				MoveIDCount:        0,
				MoveIDs:            []int32{},
			}, otherGid)
		}
	}
}

// sendDisappearMessages 通知离开视野的玩家
func sendDisappearMessages(charID int32, disappearList []string) {
	for _, otherGid := range disappearList {
		otherChara := game.CharaManagerInstance().GetCharaByGid(otherGid)
		if otherChara != nil {
			log.Info("[move] notify %s about disappear", otherChara.Name)
			broadcaster.SendToGid(&map_msg.MsgDisappear{
				CharID: charID,
				Type:   1,
			}, otherGid)
		}
	}
}
