package msg

// UserQueryReq 查询用户请求，cmd=0x0200
// 入站 body: WriteLong(id)
type UserQueryReq struct {
	ID int64
}

func (m *UserQueryReq) Cmd() uint16 { return CmdUserQueryReq }

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
