package auth

import (
	"fmt"
	"time"
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/constant"
	"wgame-server/server/context"
	"wgame-server/server/core"
	"wgame-server/server/db"
	"wgame-server/server/game"
	map_handler "wgame-server/server/handler/map"
	"wgame-server/server/model"
	"wgame-server/server/msg/auth"
	"wgame-server/server/msg/system"
	"wgame-server/server/network/handler"
)

// CmdAccountHandler 处理 CMD_L_ACCOUNT (cmd=9040)
// 客户端请求登录账号
func CmdAccountHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	reqType, _ := reader.ReadString()
	account, _ := reader.ReadString()
	_, _ = reader.ReadString() // password
	mac, _ := reader.ReadString()
	_, _ = reader.ReadString() // oaid
	_, _ = reader.ReadString() // data
	_, _ = reader.ReadString() // lock
	dist, _ := reader.ReadString()
	_, _ = reader.ReadUByte()  // from3rdSdk
	_, _ = reader.ReadString() // channel
	_, _ = reader.ReadString() // os_ver
	_, _ = reader.ReadString() // term_info
	_, _ = reader.ReadString() // imei
	_, _ = reader.ReadString() // client_original_ver
	_, _ = reader.ReadUByte()  // not_replace
	_, _ = reader.ReadUByte()  // oper_type
	_, _ = reader.ReadString() // m_value

	log.Info("[auth] 客户端请求登录账号 type=%s account=%s mac=%s dist=%s", reqType, account, mac, dist)

	if account == "" {
		ctx.Write(&auth.MsgAuth{Msg: "验证不通过，请重新登录。"})
		ctx.Disconnect()
		return nil
	}

	var accounts model.Accounts
	account = account[6:]
	result := db.AuthGORM().Where("token = ?", account).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 账号不存在 account=%s", account)
		ctx.Write(&auth.MsgAuth{Msg: "验证不通过，请重新登录。"})
		ctx.Disconnect()
		return nil
	}

	if accounts.Deleted {
		log.Error("[auth] 账号被封 account=%s", account)
		ctx.Write(&auth.MsgAuth{Msg: "账号被封,无法登录。"})
		ctx.Disconnect()
		return nil
	}

	db.AuthGORM().Model(&model.Accounts{}).Where("id = ?", accounts.ID).Update("last_login_mac", mac)

	if reqType == "cmd_l_login_preview_player" {
		log.Info("[auth] 登录预览 type=%s", reqType)
		// 发送登录预览玩家列表和登录开始消息
		ctx.Write(&auth.MsgLoginPreviewPlayer{
			ServerTime: int32(time.Now().Unix()),
			TimeZone:   8,
			Account:    account,
			Gid:        "",
			Time:       int32(time.Now().Unix()),
			Cookie:     "47Q60635Q22",
			ServerName: dist,
			IP:         "192.168.1.63",
			Port:       8800,
		})
		ctx.Write(&auth.MsgStartLogin{
			Type:   "normal",
			Cookie: "47Q60635Q22", // 固定值，与 Java 一致
		})
		return nil
	}

	ctx.Write(&auth.MsgAgentResult{
		AuthKey:      0,
		Result:       1,
		Privilege:    0,
		IP:           "192.168.1.63",
		Port:         8800,
		Seed:         0,
		ID:           1,
		ServerName:   dist,
		ServerStatus: 1,
		Msg:          "允许该账号登录",
	})
	log.Info("[auth] 客户端请求登录账号，允许该账号登录 type=%s account=%s mac=%s dist=%s", reqType, account, mac, dist)
	return nil
}

// CmdLoginHandler 处理 CMD_LOGIN (cmd=12290)
// 角色登录，返回角色列表
func CmdLoginHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	user, _ := reader.ReadString()
	authKey, _ := reader.ReadInt()
	seed, _ := reader.ReadInt()
	_, _ = reader.ReadUByte()  // emulator
	_, _ = reader.ReadUByte()  // sight_scope
	_, _ = reader.ReadString() // version
	_, _ = reader.ReadString() // clientid
	_, _ = reader.ReadShort()  // netStatus
	_, _ = reader.ReadUByte()  // adult
	_, _ = reader.ReadString() // signature
	_, _ = reader.ReadString() // clientname
	_, _ = reader.ReadUByte()  // redfinger

	log.Info("[auth] 角色登录 user=%s auth_key=%d seed=%d", user, authKey, seed)

	// 检查服务器是否就绪
	if !core.Instance().IsReady() {
		log.Error("[auth] 服务器未就绪")
		ctx.Write(&auth.MsgKickOff{Msg: "服务器正忙，请稍后再试！"})
		return nil
	}

	var accounts model.Accounts
	user = user[6:]
	result := db.AuthGORM().Where("token = ?", user).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 非法登录，账号不存在 user=%s", user)
		ctx.Write(&auth.MsgKickOff{Msg: "非法登录，账号不存在！"})
		return nil
	}

	ctx.BindUserId(int32(accounts.ID))

	var characters []model.Characters
	db.GORM().Where("account_id = ?", accounts.ID).Find(&characters)

	accountOnline := int8(0)
	voList := make([]auth.VoExistedChar, 0, len(characters))
	var firstChara *model.Characters
	for i, chara := range characters {
		if i == 0 {
			firstChara = &chara
		}
		vo := auth.VoExistedChar{
			ExistedCharInfo: auth.VoExistedCharInfo{
				LeftTimeToDelete:     0,
				CharOnlineState:      chara.Online,
				TradingGoodsGid:      "",
				Portrait:             0,
				TradingState:         0,
				TradingAppointeeName: "",
				TradingLeftTime:      0,
				Level:                chara.Level,
				Polar:                chara.Polar,
				Icon:                 0,
				Name:                 chara.Name,
				Gid:                  chara.Gid,
				TradingOrgPrice:      0,
				TradingBuyoutPrice:   0,
				TradingCgPriceCt:     0,
				TradingPrice:         0,
				TradingSellBuyType:   0,
			},
			LastLoginTime: 0,
			LoginMac:      "",
		}
		voList = append(voList, vo)

		// 检查是否有在线角色
		if chara.Online == 1 {
			accountOnline = 1
		}
	}

	ctx.Write(&auth.MsgExistedCharList{
		SeverState:     0,
		Chars:          voList,
		OpenServerTime: 0,
		AccountOnline:  accountOnline,
		LineName:       "",
	})

	// 如果有角色，使用第一个角色的信息
	if firstChara != nil {
		ctx.Write(&auth.MsgShowReconnectPara{
			Name: firstChara.Name,
			Para: "",
			Gid:  firstChara.Gid,
			Dist: "127.0.0.1:8800",
		})
	} else {
		ctx.Write(&auth.MsgShowReconnectPara{
			Name: "",
			Para: "",
			Gid:  "",
			Dist: "127.0.0.1:8800",
		})
	}

	return nil
}

// CmdLoadExistedCharHandler 处理 CMD_LOAD_EXISTED_CHAR (cmd=4192)
// 加载角色数据并进入游戏
func CmdLoadExistedCharHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	charName, _ := reader.ReadString()

	log.Info("[auth] 登录游戏 char_name=%s", charName)

	accountID := ctx.GetUserId()
	if accountID == 0 {
		ctx.Write(&auth.MsgAuth{Msg: "账号认证已过期,无法登录。"})
		ctx.Disconnect()
		return nil
	}

	var charaModel model.Characters
	result := db.GORM().Where("account_id = ? AND name = ?", accountID, charName).First(&charaModel)
	if result.Error != nil {
		log.Error("[auth] 角色不存在 account_id=%d char_name=%s", accountID, charName)
		ctx.Write(&auth.MsgKickOff{Msg: "非法登录，无效角色名"})
		ctx.Disconnect()
		return nil
	}

	if charaModel.Deleted {
		ctx.Write(&auth.MsgKickOff{Msg: "此角色已被禁闭"})
		ctx.Disconnect()
		return nil
	}

	if charaModel.Block == 1 {
		ctx.Write(&auth.MsgKickOff{Msg: "此角色已被封闭"})
		ctx.Disconnect()
		return nil
	}

	ctx.BindUserId(int32(charaModel.ID))
	ctx.BindGid(charaModel.Gid)

	chara := game.NewChara(charName, charaModel.Sex, charaModel.Polar, charaModel.Gid)
	chara.ID = int32(charaModel.ID)
	chara.Level = charaModel.Level
	chara.MapId = charaModel.MapId
	chara.MapName = charaModel.MapName
	chara.X = charaModel.X
	chara.Y = charaModel.Y
	chara.Line = 1
	chara.GoldCoin = charaModel.GoldCoin
	chara.Nice = 0 // 默认好心值
	chara.Dir = 0  // 默认方向
	chara.Waiguan = 0

	game.CharaManagerInstance().AddChara(chara)

	// 发送角色基础信息消息
	loadCharaBases(ctx, chara)

	// 进入地图
	gameMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
	if gameMap != nil {
		map_handler.EnterMap(ctx, chara, gameMap, chara.X, chara.Y)
	}

	// 更新数据库在线状态
	db.GORM().Model(&model.Characters{}).Where("id = ?", charaModel.ID).Update("online", 1)

	log.Info("[auth] 玩家登录成功 id=%d gid=%s name=%s", chara.ID, chara.Gid, chara.Name)

	return nil
}

// loadCharaBases 参考 Java 端 LoginUtil.loadCharaBases 方法
// 发送角色基础信息相关的消息
func loadCharaBases(ctx context.MyCmdContext, chara *game.Chara) {
	// 1. MSG_CHAR_ALREADY_LOGIN - 通知客户端角色已登录
	ctx.Write(&system.MsgCharAlreadyLogin{Name: chara.Name})

	// 2. MSG_LOGIN_DONE - 登录完成（角色连接完成）
	ctx.Write(&auth.MsgShowReconnectPara{
		Name: chara.Name,
		Para: "",
		Gid:  chara.Gid,
		Dist: "普通通道", // 对应 Java GameConfig.getLineName()
	})

	// 3. MSG_CS_SERVER_TYPE - 更新服务器类型
	// 普通服务器类型为 0
	ctx.Write(&system.MsgCsServerType{ServerType: 0})

	// 4. MSG_REPLY_SERVER_TIME - 服务器时间回复
	ctx.Write(system.NewMsgReplyServerTime())

	// 5. MSG_UPDATE - 角色基础属性更新
	updateMsg := system.NewMsgUpdate(chara.ID)
	updateMsg.Props.Name = chara.Name
	updateMsg.Props.Level = chara.Level
	updateMsg.Props.Polar = chara.Polar
	updateMsg.Props.Icon = chara.Waiguan
	updateMsg.Props.Str = chara.Str
	updateMsg.Props.Dex = chara.Dex
	updateMsg.Props.Con = chara.Con
	updateMsg.Props.Wiz = chara.Wiz
	updateMsg.Props.Life = chara.Shengming
	updateMsg.Props.MaxLife = chara.MaxShengming
	updateMsg.Props.Mana = chara.Mofa
	updateMsg.Props.MaxMana = chara.MaxMofa
	updateMsg.Props.Metal = chara.Metal
	updateMsg.Props.Wood = chara.Wood
	updateMsg.Props.Water = chara.Water
	updateMsg.Props.Fire = chara.Fire
	updateMsg.Props.Earth = chara.Earth
	updateMsg.Props.Cash = chara.Cash
	updateMsg.Props.Balance = chara.Balance
	updateMsg.Props.GoldCoin = int64(chara.GoldCoin)
	updateMsg.Props.Nice = chara.Nice
	updateMsg.Props.Tao = chara.Tao
	updateMsg.Props.Online = 1
	ctx.Write(updateMsg)

	// 6. MSG_SET_SETTING - 系统设置
	ctx.Write(system.NewMsgSetSetting())

	// 7. MSG_GENERAL_NOTIFY - 御宝仙术通知
	// Java: notify=60001(NOTIFY_TEST_YUBXS_END_TIME), para=DateTimeUtil.now() + 9999999
	now := int32(time.Now().Unix())
	ctx.Write(&system.MsgGeneralNotify{
		Notify: constant.NotifyTestYbxsEndTime,
		Para:   fmt.Sprintf("%d", now+9999999),
	})

	// 8. MSG_UPDATE - 第二次属性更新
	ctx.Write(updateMsg)

	// 9. MSG_GENERAL_NOTIFY - 邮件加载完成通知
	// Java: notify=10012 (NOTIFY_MAIL_ALL_LOADED), para="1"
	ctx.Write(&system.MsgGeneralNotify{
		Notify: constant.NotifyMailAllLoaded,
		Para:   "1",
	})

	// === loadCharaOther 部分 ===

	// 10. MSG_SET_PUSH_SETTINGS - 推送设置
	// Java: value="10011011111"
	ctx.Write(&system.MsgSetPushSettings{Value: "10011011111"})

	// 11. MSG_GENERAL_NOTIFY - iOS审核通知
	// Java: notify=50017 (NOTIFY_IOS_REVIEW), para="0"
	ctx.Write(&system.MsgGeneralNotify{
		Notify: constant.NotifyIOSReview,
		Para:   "0",
	})

	// 12. MSG_GENERAL_NOTIFY - 驱魔香通知
	// Java: notify=20010, para=String.valueOf(chara.qumoxiang)
	ctx.Write(&system.MsgGeneralNotify{
		Notify: constant.NotifyQmoxiangStatus,
		Para:   fmt.Sprintf("%d", chara.Qumoxiang),
	})

	// 13. MSG_FUZZY_IDENTITY - 身份验证
	// Java 固定值: isBindName=1, isBindPhone=1, bindName="王老板", bindId="110101192003127055", bindPhone="130****6666"
	ctx.Write(&system.MsgFuzzyIdentity{
		IsBindName:  1,
		IsBindPhone: 1,
		BindName:    "王老板",
		BindId:      "110101192003127055",
		BindPhone:   "130****6666",
	})

	// 14. MSG_EXECUTE_LUA_CODE - 执行Lua代码(特殊消息)
	ctx.Write(&system.MsgExecuteLuaCode{
		Cookie: 0,
		Code:   "GameMgr.checkServer = function()\nreturn 0\nend",
		Flag:   1,
	})

	ctx.Write(&system.MsgExecuteLuaCode{
		Cookie: 0,
		Code:   "BarrageTalkMgr:creatBarrageLayer()",
		Flag:   1,
	})

	// 15. MSG_GENERAL_NOTIFY - 初始化完成通知(必须在MSG_EXECUTE_LUA_CODE之后)
	// Java: notify=39 (NOTIFY_SEND_INIT_DATA_DONE), para=""
	ctx.Write(&system.MsgGeneralNotify{
		Notify: constant.NotifySendInitDataDone,
		Para:   "",
	})
}

func init() {
	handler.Register(9040, "CmdAccount", CmdAccountHandler)
	handler.Register(12290, "CmdLogin", CmdLoginHandler)
	handler.Register(4192, "CmdLoadExistedChar", CmdLoadExistedCharHandler)
	handler.Register(13140, "CmdGetServerList", CmdGetServerListHandler)
	handler.Register(45144, "CmdRequestLineInfo", CmdRequestLineInfoHandler)
}

// CmdGetServerListHandler 处理 CMD_L_GET_SERVER_LIST (cmd=13140)
// 客户端请求服务器列表
func CmdGetServerListHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	account, _ := reader.ReadString()
	_, _ = reader.ReadInt()    // auth_key
	_, _ = reader.ReadString() // dist

	log.Info("[auth] 客户端请求服务器列表 account=%s", account)

	var accounts model.Accounts
	account = account[6:]
	result := db.AuthGORM().Where("token = ?", account).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 账号不存在 account=%s", account)
		ctx.Write(&auth.MsgAuth{Msg: "账号认证已过期,无法登录。"})
		ctx.Disconnect()
		return nil
	}

	// 发送服务器列表消息
	ctx.Write(&auth.MsgServerList{})

	return nil
}

// CmdRequestLineInfoHandler 处理 CMD_L_REQUEST_LINE_INFO (cmd=45144)
// 客户端请求线路信息
func CmdRequestLineInfoHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	account, _ := reader.ReadString()

	log.Info("[auth] 客户端请求线路信息 account=%s", account)
	var accounts model.Accounts
	account = account[6:]
	result := db.AuthGORM().Where("token = ?", account).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 登录信息已失效 account=%s", account)
		ctx.Write(&auth.MsgAuth{Msg: "登录信息已失效，请重新登录。"})
		ctx.Disconnect()
		return nil
	}

	sendAdmittedLineInfoAndStart(ctx)

	return nil
}

// sendAdmittedLineInfoAndStart 参考 Java 端 LoginLineQueue.sendAdmittedLineInfoAndStart 方法
// 放行后，按原协议顺序下发"线路信息/开始登录/客户端连接确认"
func sendAdmittedLineInfoAndStart(ctx context.MyCmdContext) {
	ctx.Write(system.NewMsgReplyServerTime())
	ctx.Write(&auth.MsgWaitInLine{
		LineName:             "普通通道",
		ExpectTime:           180000,
		ReconnectTime:        180000,
		WaitCode:             1,
		Count:                1,
		KeepAlive:            1,
		NeedWait:             0,
		InsiderLv:            1,
		GoldCoin:             0,
		Status:               1,
		StartServerTime:      int32(time.Now().Unix()),
		LeftGiveLotteryTimes: 1000,
	})
	ctx.Write(&auth.MsgStartLogin{
		Type:   "normal",
		Cookie: "47Q60635Q22",
	})
	ctx.Write(&auth.MsgCheckUserData{
		Result: 1,
		Cookie: "47Q60635Q22",
	})
}
