package chat_handler

import (
	"time"
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/core"
	"wgame-server/server/game"
	chat_msg "wgame-server/server/msg/chat"
	"wgame-server/server/network/handler"
)

// CHAT_CHANNEL 频道定义
const (
	CHAT_CHANNEL_MISC    = 0  // 混合频道
	CHAT_CHANNEL_CURRENT = 1  // 当前频道
	CHAT_CHANNEL_WORLD   = 2  // 世界频道
	CHAT_CHANNEL_TEAM    = 4  // 队伍频道
	CHAT_CHANNEL_PARTY   = 5  // 帮会频道
	CHAT_CHANNEL_FRIEND  = 9  // 好友频道
	CHAT_CHANNEL_TRADE   = 13 // 交易频道
	CHAT_CHANNEL_WHOOP   = 15 // 呐喊频道
	CHAT_CHANNEL_HORN    = 30 // 喇叭
	CHAT_CHANNEL_WEDDING = 31 // 婚礼弹幕
)

// CmdChatExHandler 处理 CMD_CHAT_EX (cmd=16482)
// 公共聊天消息
func CmdChatExHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	channel, _ := reader.ReadShort()
	compress, _ := reader.ReadShort()
	orgLength, _ := reader.ReadShort()

	var msg string
	if compress == 9999 {
		msg, _ = reader.ReadString()
	} else {
		msg, _ = reader.ReadString2()
	}

	cardCount, _ := reader.ReadShort()
	for i := int16(0); i < cardCount; i++ {
		_, _ = reader.ReadString2() // cardParams
	}

	voiceTime, _ := reader.ReadInt()
	token, _ := reader.ReadString2()
	_, _ = reader.ReadString() // para

	log.Info("[chat] 玩家聊天 channel=%d msg=%s", channel, msg)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[chat] chara not found gid=%s", gid)
		return nil
	}

	// 基础检查
	if chara.Jinyan != 0 {
		log.Error("[chat] 玩家已被禁言 gid=%s", gid)
		return nil
	}

	if chara.Level < 10 { // 最低发言等级限制
		log.Error("[chat] 等级不足 gid=%s level=%d", gid, chara.Level)
		return nil
	}

	// 创建聊天消息
	chatMsg := &chat_msg.MsgMessageEx{
		Channel:         int32(channel),
		ID:              chara.ID,
		Name:            chara.Name,
		Msg:             msg,
		Time:            int32(time.Now().Unix()),
		Privilege:       0, // chara.Privilege 暂未定义
		LineName:        "line1",
		ShowExtra:       0, // 文字消息
		Compress:        compress,
		OrgLength:       orgLength,
		CardCount:       cardCount,
		Cards:           []chat_msg.ChatCard{},
		VoiceTime:       voiceTime,
		Token:           token,
		Checksum:        0,
		FilteredMsg:     "",
		MapSize:         0,
		ItemCookieCount: 0,
		TipIndex:        0,
	}

	// 根据频道类型发送消息
	switch channel {
	case CHAT_CHANNEL_CURRENT:
		// 当前频道：发送给同地图所有玩家
		gameMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
		if gameMap != nil {
			// TODO: 广播给地图内所有玩家
			log.Info("[chat] 当前频道广播 chara=%s map=%d msg=%s", chara.Name, chara.MapId, msg)
			_ = chatMsg // 暂存，后续实现广播时使用
		}
	case CHAT_CHANNEL_WORLD:
		// 世界频道：发送给所有在线玩家
		log.Info("[chat] 世界频道广播 chara=%s msg=%s", chara.Name, msg)
		_ = chatMsg
	case CHAT_CHANNEL_TEAM:
		// 队伍频道：发送给队伍成员
		log.Info("[chat] 队伍频道 chara=%s msg=%s", chara.Name, msg)
		_ = chatMsg
	case CHAT_CHANNEL_PARTY:
		// 帮派频道：发送给帮派成员
		log.Info("[chat] 帮派频道 chara=%s msg=%s", chara.Name, msg)
		_ = chatMsg
	default:
		log.Info("[chat] 其他频道 channel=%d chara=%s msg=%s", channel, chara.Name, msg)
		_ = chatMsg
	}

	return nil
}

// CmdFriendTellExHandler 处理 CMD_FRIEND_TELL_EX (cmd=20590)
// 私聊消息
func CmdFriendTellExHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	_, _ = reader.ReadShort()      // flag
	name, _ := reader.ReadString() // 目标名称
	compress, _ := reader.ReadShort()
	orgLength, _ := reader.ReadShort()
	msg, _ := reader.ReadString2()
	cardCount, _ := reader.ReadShort()
	if cardCount != 0 {
		_, _ = reader.ReadString() // cardParam
	}
	voiceTime, _ := reader.ReadInt()
	token, _ := reader.ReadString2()
	receiveGid, _ := reader.ReadString() // 目标gid
	_, _ = reader.ReadString2()          // push_msg

	log.Info("[chat] 私聊 from=%s to=%s msg=%s", ctx.GetGid(), name, msg)

	gid := ctx.GetGid()
	if gid == "" {
		return nil
	}

	chara := game.CharaManagerInstance().GetCharaByGid(gid)
	if chara == nil {
		log.Error("[chat] chara not found gid=%s", gid)
		return nil
	}

	// 基础检查
	if chara.Jinyan != 0 {
		return nil
	}

	if chara.Level < 10 {
		return nil
	}

	// 查找目标玩家
	targetChara := game.CharaManagerInstance().GetCharaByGid(receiveGid)
	if targetChara == nil {
		log.Error("[chat] 目标玩家不存在 receiveGid=%s", receiveGid)
		return nil
	}

	// 创建聊天消息
	chatMsg := &chat_msg.MsgMessageEx{
		Channel:         CHAT_CHANNEL_FRIEND,
		ID:              chara.ID,
		Name:            chara.Name,
		Msg:             msg,
		Time:            int32(time.Now().Unix()),
		Privilege:       0, // chara.Privilege 暂未定义
		LineName:        "line1",
		ShowExtra:       0,
		Compress:        compress,
		OrgLength:       orgLength,
		CardCount:       cardCount,
		Cards:           []chat_msg.ChatCard{},
		VoiceTime:       voiceTime,
		Token:           token,
		Checksum:        0,
		FilteredMsg:     "",
		MapSize:         0,
		ItemCookieCount: 0,
		TipIndex:        0,
	}

	log.Info("[chat] 私聊成功 from=%s to=%s msg=%s", chara.Name, targetChara.Name, msg)
	_ = chatMsg // 暂存，后续实现私聊发送时使用

	return nil
}

func init() {
	handler.Register(16482, "CmdChatEx", CmdChatExHandler)
	handler.Register(20590, "CmdFriendTellEx", CmdFriendTellExHandler)
}
