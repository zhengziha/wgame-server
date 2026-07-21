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
		Name: charName,
		Gid:  newUUID,
	})

	var charList []model.Characters
	db.GORM().Where("account_id = ?", accountID).Find(&charList)

	voList := make([]*auth.VoExistedChar, 0, len(charList))
	for _, chara := range charList {
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
	}

	ctx.Write(&auth.MsgExistedCharList{
		AccountOnline: 0,
		VoList:        voList,
	})

	return nil
}

func init() {
	handler.Register(8284, "CmdCreateNewChar", CmdCreateNewCharHandler)
}
