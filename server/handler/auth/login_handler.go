package auth

import (
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/core"
	"wgame-server/server/db"
	"wgame-server/server/game"
	map_handler "wgame-server/server/handler/map"
	"wgame-server/server/model"
	"wgame-server/server/msg/auth"
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
	result := db.AuthGORM().Where("name = ?", account).First(&accounts)
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
			Account: account,
			Data:    "",
		})
		ctx.Write(&auth.MsgStartLogin{
			Type:   "normal",
			Cookie: "47Q60635Q22", // 固定值，与 Java 一致
		})
		return nil
	}

	ctx.Write(&auth.MsgAgentResult{
		Result:       1,
		ID:           int32(accounts.ID),
		Privilege:    accounts.Privilege,
		IP:           "127.0.0.1",
		Port:         8800,
		ServerName:   dist,
		ServerStatus: 1,
		Msg:          "允许该账号登录",
	})

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
	result := db.AuthGORM().Where("name = ?", user).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 非法登录，账号不存在 user=%s", user)
		ctx.Write(&auth.MsgKickOff{Msg: "非法登录，账号不存在！"})
		return nil
	}

	ctx.BindUserId(int32(accounts.ID))

	var characters []model.Characters
	db.GORM().Where("account_id = ?", accounts.ID).Find(&characters)

	accountOnline := int32(0)
	voList := make([]*auth.VoExistedChar, 0, len(characters))
	for _, chara := range characters {
		vo := &auth.VoExistedChar{
			CharID:          int32(chara.ID),
			Name:            chara.Name,
			Level:           chara.Level,
			Polar:           chara.Polar,
			Sex:             chara.Sex,
			OnlineState:     chara.Online,
			FashionIcon:     0,
			UpgradeLevel:    0,
			PetIcon:         0,
			MountIcon:       0,
			SpecialIcon:     0,
			GenchongIcon:    0,
			UpgradeType:     0,
			Nice:            0,
			WeeklyLoginDays: 0,
			IsFeisheng:      0,
			Tao:             chara.MonthTao,
			Gid:             chara.Gid,
			MapID:           chara.MapId,
			MapName:         chara.MapName,
			Line:            1,
			X:               chara.X,
			Y:               chara.Y,
			PartyName:       "",
			Family:          "",
			Title:           "",
		}
		voList = append(voList, vo)

		// 检查是否有在线角色
		if chara.Online == 1 {
			accountOnline = 1
		}
	}

	ctx.Write(&auth.MsgExistedCharList{
		AccountOnline: accountOnline,
		VoList:        voList,
	})

	ctx.Write(&auth.MsgShowReconnectPara{
		IP:      "127.0.0.1",
		Port:    8800,
		AuthKey: authKey,
		Seed:    seed,
	})

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

	gameMap := core.Instance().GetGameMap(chara.Line, chara.MapId)
	if gameMap != nil {
		map_handler.EnterMap(ctx, chara, gameMap, chara.X, chara.Y)
	}

	// 更新数据库在线状态
	db.GORM().Model(&model.Characters{}).Where("id = ?", charaModel.ID).Update("online", 1)

	log.Info("[auth] 玩家登录成功 id=%d gid=%s name=%s", chara.ID, chara.Gid, chara.Name)

	return nil
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
	result := db.AuthGORM().Where("name = ?", account).First(&accounts)
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
	result := db.AuthGORM().Where("name = ?", account).First(&accounts)
	if result.Error != nil {
		log.Error("[auth] 登录信息已失效 account=%s", account)
		ctx.Write(&auth.MsgAuth{Msg: "登录信息已失效，请重新登录。"})
		ctx.Disconnect()
		return nil
	}

	// 获取角色的元宝数量（取第一个角色的元宝）
	var characters []model.Characters
	db.GORM().Where("account_id = ?", accounts.ID).First(&characters)
	goldCoin := int32(0)
	if len(characters) > 0 {
		goldCoin = characters[0].GoldCoin
	}

	// 发送线路信息
	ctx.Write(&auth.MsgWaitInLine{
		LineName:        "line1", // 默认线路名
		ExpectTime:      0,       // 无需等待
		ReconnectTime:   0,       // 无需重连
		WaitCode:        0,       // 排名
		Count:           1,       // 线路数量
		KeepAlive:       0,       // 保持连接
		NeedWait:        1,       // 显示线路和排名
		InsiderLv:       0,       // 会员等级
		GoldCoin:        goldCoin,
		Status:          0, // 正常
		StartServerTime: 0, // 开服时间（暂不设置）
	})

	return nil
}
