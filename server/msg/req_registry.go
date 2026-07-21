package msg

import "sync"

// 本文件维护按 cmd 索引的入站请求构造工厂。
//
// 用途：
//   - handler.RegisterTyped 在派发时自动构造 Req 对象并填充字段
//   - 调试工具、客户端模拟器可通过 cmd 反查 Req 类型构造空实例
//
// 注册推荐放在各业务 msg 包的 init() 中（或随 RegisterTyped 自动注册）。

// ReqFactory 按 cmd 注册的请求构造函数，返回一个空的 InMessage 实例。
type ReqFactory func() InMessage

var (
	reqMu       sync.RWMutex
	reqRegistry = make(map[uint16]ReqFactory, 256)
)

// RegisterReq 注册一个 cmd 对应的请求工厂。
// 重复注册（同 cmd）会覆盖。
func RegisterReq(cmd uint16, f ReqFactory) {
	if f == nil {
		panic("msg: RegisterReq factory is nil for cmd " + cmdHex(cmd))
	}
	reqMu.Lock()
	defer reqMu.Unlock()
	reqRegistry[cmd] = f
}

// NewReq 根据 cmd 查找工厂并构造一个 InMessage 实例。
// 未注册时返回 nil, false。
func NewReq(cmd uint16) (InMessage, bool) {
	reqMu.RLock()
	defer reqMu.RUnlock()
	f, ok := reqRegistry[cmd]
	if !ok {
		return nil, false
	}
	return f(), true
}

// IsReqRegistered 判断指定 cmd 是否已注册请求工厂
func IsReqRegistered(cmd uint16) bool {
	reqMu.RLock()
	defer reqMu.RUnlock()
	_, ok := reqRegistry[cmd]
	return ok
}

func cmdHex(c uint16) string {
	const hex = "0123456789ABCDEF"
	return "0x" + string([]byte{hex[(c>>12)&0xF], hex[(c>>8)&0xF], hex[(c>>4)&0xF], hex[c&0xF]})
}
