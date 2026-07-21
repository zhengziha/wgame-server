package msg

import "wgame-server/server/codec"

// UserQueryResp 查询用户返回，cmd=0x0201
type UserQueryResp struct {
	Found    bool
	ID       int64
	Account  string
	Nickname string
	Level    int
}

const CmdUserQueryReq uint16 = 0x0200
const CmdUserQueryResp uint16 = 0x0201

func (m *UserQueryResp) Cmd() uint16 { return CmdUserQueryResp }

func (m *UserQueryResp) WriteBody(w *codec.GameWriter) {
	w.WriteBoolean(m.Found)
	w.WriteLong(m.ID)
	w.WriteString(m.Account)
	w.WriteString(m.Nickname)
	w.WriteInt(int32(m.Level))
}
