package auth

import (
	"strings"
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/db"
	"wgame-server/server/model"
	"wgame-server/server/msg/auth"
	"wgame-server/server/network/handler"

	"github.com/google/uuid"
)

// CmdCreateNewCharHandler 处理 CMD_CREATE_NEW_CHAR (cmd=8284)
// 创建新角色
func CmdCreateNewCharHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	charName, _ := reader.ReadString()
	gender, _ := reader.ReadShort()
	polar, _ := reader.ReadShort()

	log.Info("[auth] 创建角色 name=%s gender=%d polar=%d", charName, gender, polar)

	accountID := ctx.GetUserId()
	if accountID == 0 {
		ctx.Write(&auth.MsgKickOff{Msg: "请先登录账号"})
		return nil
	}

	var accounts model.Accounts
	result := db.GORM().Where("id = ?", accountID).First(&accounts)
	if result.Error != nil {
		ctx.Write(&auth.MsgKickOff{Msg: "账号不存在"})
		return nil
	}

	charName = strings.TrimSpace(charName)
	if charName == "" {
		ctx.Write(&auth.MsgKickOff{Msg: "昵称不能为空"})
		return nil
	}

	if len(charName) > 12 {
		ctx.Write(&auth.MsgKickOff{Msg: "昵称不能超过12个字符"})
		return nil
	}

	var count int64
	db.GORM().Model(&model.Characters{}).Where("deleted = false AND name = ?", charName).Count(&count)
	if count > 0 {
		ctx.Write(&auth.MsgKickOff{Msg: "该名字已被使用"})
		return nil
	}

	newUUID := uuid.New().String()
	newUUID = strings.ReplaceAll(newUUID, "-", "")

	characters := &model.Characters{
		Name:      charName,
		Polar:     int32(polar),
		Sex:       int32(gender),
		Gid:       newUUID,
		AccountId: accountID,
		Level:     1,
		MapId:     10001,
		MapName:   "无名小镇",
		X:         400,
		Y:         300,
		Block:     0,
		Xiaozi:    0,
		Online:    1,
	}

	result = db.GORM().Create(characters)
	if result.Error != nil {
		log.Error("[auth] 创建角色失败 name=%s err=%v", charName, result.Error)
		ctx.Write(&auth.MsgKickOff{Msg: "创建角色失败"})
		return nil
	}

	log.Info("[auth] 创建角色成功 id=%d name=%s gid=%s", characters.ID, charName, newUUID)

	ctx.Write(&auth.MsgCreateNewChar{
		Gid:  newUUID,
		Name: charName,
	})

	var charList []model.Characters
	db.GORM().Where("account_id = ?", accountID).Find(&charList)

	voList := make([]auth.VoExistedChar, 0, len(charList))
	for _, chara := range charList {
		vo := auth.VoExistedChar{
			Fixed17:              17,
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
			LastLoginTime:        0,
			LoginMac:             "",
		}
		voList = append(voList, vo)
	}

	ctx.Write(&auth.MsgExistedCharList{
		SeverState:     0,
		CharCount:      int16(len(voList)),
		Chars:          voList,
		OpenServerTime: 0,
		AccountOnline:  0,
		LineName:       "",
	})

	return nil
}

func init() {
	handler.Register(8284, "CmdCreateNewChar", CmdCreateNewCharHandler)
}
