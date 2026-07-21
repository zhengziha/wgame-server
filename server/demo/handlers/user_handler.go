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

func init() {
	handler.RegisterTyped(
		demomsg.CmdUserQueryReq,
		"UserQuery",
		func() *demomsg.UserQueryReq { return &demomsg.UserQueryReq{} },
		UserQueryHandler,
	)
}

// UserQueryHandler 直接拿到已反序列化好的 *UserQueryReq，
// 不再需要手写 reader.ReadLong() 等读取代码。
func UserQueryHandler(ctx myctx.MyCmdContext, frame *codec.Frame, req *demomsg.UserQueryReq) error {
	id := req.ID

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
