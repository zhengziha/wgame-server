package handlers

import (
	"context"
	"time"

	"wgame-server/comm/log"
	"wgame-server/server/codec"
	myctx "wgame-server/server/context"
	"wgame-server/server/dao"
	demomsg "wgame-server/server/demo/msg"
	"wgame-server/server/network/handler"
)

// CmdUserQueryReq 与 demomsg 中定义保持一致
const cmdUserQueryReq uint16 = demomsg.CmdUserQueryReq

func init() {
	handler.Register(cmdUserQueryReq, "UserQuery", UserQueryHandler)
}

// UserQueryHandler 读取一个 int64 用户 id，返回 UserQueryResp。
// 演示了从自定义协议层调用 DAO（SQLite + Redis）的完整路径。
//
// 入站 body: WriteLong(id)
func UserQueryHandler(ctx myctx.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	id, err := reader.ReadLong()
	if err != nil {
		return err
	}

	// DAO 调用：自带 5 秒超时，避免单请求卡住主线程
	cctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d := dao.NewUserDAO()
	u, err := d.GetByID(cctx, id)
	if err != nil {
		log.Error("[demo] user query err id=%d: %v", id, err)
		// 出错也回一个 Found=false，避免客户端等超时
		ctx.Write(&demomsg.UserQueryResp{Found: false})
		return err
	}

	resp := &demomsg.UserQueryResp{Found: false}
	if u != nil {
		resp.Found = true
		resp.ID = u.ID
		resp.Account = u.AccountName
		resp.Nickname = u.Nickname
		resp.Level = u.Level
	}
	log.Info("[demo] user query id=%d found=%v", id, resp.Found)
	ctx.Write(resp)
	return nil
}
