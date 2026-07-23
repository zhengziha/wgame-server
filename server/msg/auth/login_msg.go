package auth

import (
	"wgame-server/server/msg"
)

// MsgAgentResult 对应 Java MSG_L_AGENT_RESULT (cmd=13143)
// 账号认证成功后返回给客户端
// Java写入顺序：auth_key(int), result(int), privilege(short), ip(String), port(short), seed(int), id(short), serverName(String), serverStatus(byte), msg(String)
type MsgAgentResult struct {
	AuthKey      int32  `codec:"int"`    // 认证key，固定为0
	Result       int32  `codec:"int"`    // 结果 1=成功
	Privilege    int16  `codec:"short"`  // 权限等级
	IP           string `codec:"string"` // 服务器IP
	Port         int16  `codec:"short"`  // 服务器端口
	Seed         int32  `codec:"int"`    // 种子，固定为0
	ID           int16  `codec:"short"`  // 账号ID
	ServerName   string `codec:"string"` // 服务器名称
	ServerStatus int8   `codec:"byte"`   // 服务器状态
	Msg          string `codec:"string"` // 消息
}

func (m *MsgAgentResult) Cmd() uint16 {
	return 13143
}

// MsgAuth 对应 Java MSG_L_AUTH (cmd=21329)
// 账号认证失败时返回给客户端
// Java写入顺序：type(String), result(int), auth_key(int), msg(String2)
type MsgAuth struct {
	Type    string `codec:"string"`  // 认证类型
	Result  int32  `codec:"int"`     // 结果码
	AuthKey int32  `codec:"int"`     // 认证key
	Msg     string `codec:"string2"` // 错误消息
}

func (m *MsgAuth) Cmd() uint16 {
	return 21329
}

// MsgExistedCharList 对应 Java MSG_EXISTED_CHAR_LIST (cmd=61537)
// 返回账号下的角色列表
// Java写入顺序：severState(short), listSize(short), chars(list), openServerTime(int), account_online(byte), lineName(String)
type MsgExistedCharList struct {
	SeverState     int16           `codec:"short"`      // 服务器状态
	CharCount      int16           `codec:"short"`      // 角色数量
	Chars          []VoExistedChar `codec:"list:short"` // 角色列表
	OpenServerTime int32           `codec:"int"`        // 开服时间
	AccountOnline  int8            `codec:"byte"`       // 账号是否在线
	LineName       string          `codec:"string"`     // 线路名称
}

func (m *MsgExistedCharList) Cmd() uint16 {
	return 61537
}

// VoExistedChar 对应 Java Vo_61537_0
// 单个角色信息
// Java中使用BuildFieldsNew按顺序写入各个字段
// 字段顺序（对照Java MSG_EXISTED_CHAR_LIST.writeO）：
//  1. left_time_to_delete (int, key=263)
//  2. char_online_state (int, key=435)
//  3. trading_goods_gid (string, key=428)
//  4. portrait (int, key=86)
//  5. trading_state (int, key=429)
//  6. trading_appointee_name (string, key=437)
//  7. trading_left_time (int, key=430)
//  8. level (int, key=31)
//  9. polar (int, key=44)
//  10. icon (int, key=40)
//  11. name (string, key=1)
//  12. gid (string, key=305)
//  13. trading_org_price (int, key=432)
//  14. trading_buyout_price (int, key=438)
//  15. trading_cg_price_ct (int, key=434)
//  16. trading_price (int, key=431)
//  17. trading_sell_buy_type (int, key=436)
//  18. last_login_time (int) - 普通字段
//  19. login_mac (string) - 普通字段
type VoExistedChar struct {
	Fixed17              int16  `codec:"short"`                 // 固定值17，表示后面有17个BuildFieldsNew字段
	LeftTimeToDelete     int32  `codec:"buildfield:int:263"`    // 剩余删除时间
	CharOnlineState      int32  `codec:"buildfield:int:435"`    // 角色在线状态
	TradingGoodsGid      string `codec:"buildfield:string:428"` // 交易商品GID
	Portrait             int32  `codec:"buildfield:int:86"`     // 头像
	TradingState         int32  `codec:"buildfield:int:429"`    // 交易状态
	TradingAppointeeName string `codec:"buildfield:string:437"` // 交易指定玩家名
	TradingLeftTime      int32  `codec:"buildfield:int:430"`    // 交易剩余时间
	Level                int32  `codec:"buildfield:int:31"`     // 等级
	Polar                int32  `codec:"buildfield:int:44"`     // 门派
	Icon                 int32  `codec:"buildfield:int:40"`     // 图标
	Name                 string `codec:"buildfield:string:1"`   // 角色名
	Gid                  string `codec:"buildfield:string:305"` // 全局唯一ID
	TradingOrgPrice      int32  `codec:"buildfield:int:432"`    // 交易原价
	TradingBuyoutPrice   int32  `codec:"buildfield:int:438"`    // 交易一口价
	TradingCgPriceCt     int32  `codec:"buildfield:int:434"`    // 交易成功价格
	TradingPrice         int32  `codec:"buildfield:int:431"`    // 交易价格
	TradingSellBuyType   int32  `codec:"buildfield:int:436"`    // 交易买卖类型
	LastLoginTime        int32  `codec:"int"`                   // 最后登录时间
	LoginMac             string `codec:"string"`                // 登录MAC地址
}

// MsgKickOff 对应 Java MSG_KICK_OFF (cmd=53405)
// 踢下线消息
// Java写入顺序：msg(String)
type MsgKickOff struct {
	Msg string `codec:"string"` // 踢下线原因
}

func (m *MsgKickOff) Cmd() uint16 {
	return 53405
}

// MsgShowReconnectPara 对应 Java MSG_LOGIN_DONE (cmd=4099)
// 重连参数/登录完成消息
// Java写入顺序：name(String), para(String), gid(String), dist(String)
type MsgShowReconnectPara struct {
	Name string `codec:"string"` // 角色名
	Para string `codec:"string"` // 参数
	Gid  string `codec:"string"` // 全局唯一ID
	Dist string `codec:"string"` // 线路名称
}

func (m *MsgShowReconnectPara) Cmd() uint16 {
	return 4099
}

// MsgStartLogin 对应 Java MSG_L_START_LOGIN (cmd=45555)
// 登录开始消息（登录预览后发送）
// Java写入顺序：type(String), cookie(String)
type MsgStartLogin struct {
	Type   string `codec:"string"` // 类型 "normal"
	Cookie string `codec:"string"` // Cookie值
}

func (m *MsgStartLogin) Cmd() uint16 {
	return 45555
}

// MsgLoginPreviewPlayer 对应 Java MSG_L_LOGIN_PREVIEW_PLAYER (cmd=24113)
// 登录预览玩家列表
// Java写入顺序：server_time(int), time_zone(byte), account(String), gid(String), time(int), cookie(String), serverName(String), ip(String), port(short)
type MsgLoginPreviewPlayer struct {
	ServerTime int32  `codec:"int"`    // 服务器时间
	TimeZone   int8   `codec:"byte"`   // 时区
	Account    string `codec:"string"` // 账号
	Gid        string `codec:"string"` // 全局唯一ID
	Time       int32  `codec:"int"`    // 时间
	Cookie     string `codec:"string"` // Cookie值
	ServerName string `codec:"string"` // 服务器名称
	IP         string `codec:"string"` // IP地址
	Port       int16  `codec:"short"`  // 端口
}

func (m *MsgLoginPreviewPlayer) Cmd() uint16 {
	return 24113
}

// MsgServerList 对应 Java MSG_L_SERVER_LIST (cmd=17237)
// 服务器列表（空实现）
type MsgServerList struct {
}

func (m *MsgServerList) Cmd() uint16 {
	return 17237
}

// MsgWaitInLine 对应 Java MSG_L_WAIT_IN_LINE (cmd=45143)
// 线路信息/排队信息
// Java写入顺序：line_name(String), expect_time(int), reconnet_time(int), waitCode(int), count(int), keep_alive(byte), need_wait(byte), indsider_lv(byte), gold_coin(int), status(byte), start_server_time(int), left_give_lottery_times(short)
type MsgWaitInLine struct {
	LineName             string `codec:"string"` // 分配的线路名称
	ExpectTime           int32  `codec:"int"`    // 等待时间
	ReconnectTime        int32  `codec:"int"`    // 重新获取数据的时间
	WaitCode             int32  `codec:"int"`    // 排名
	Count                int32  `codec:"int"`    // 线路数量
	KeepAlive            int8   `codec:"byte"`   // 保持连接
	NeedWait             int8   `codec:"byte"`   // 是否需要等待
	InsiderLv            int8   `codec:"byte"`   // 会员等级
	GoldCoin             int32  `codec:"int"`    // 账号元宝数量
	Status               int8   `codec:"byte"`   // 服务器状态
	StartServerTime      int32  `codec:"int"`    // 开服时间
	LeftGiveLotteryTimes int16  `codec:"short"`  // 剩余赠送抽奖次数
}

func (m *MsgWaitInLine) Cmd() uint16 {
	return 45143
}

// 确保实现 msg.OutMessage 接口
var _ msg.OutMessage = (*MsgAgentResult)(nil)
var _ msg.OutMessage = (*MsgAuth)(nil)
var _ msg.OutMessage = (*MsgExistedCharList)(nil)
var _ msg.OutMessage = (*MsgKickOff)(nil)
var _ msg.OutMessage = (*MsgShowReconnectPara)(nil)
var _ msg.OutMessage = (*MsgStartLogin)(nil)
var _ msg.OutMessage = (*MsgLoginPreviewPlayer)(nil)
var _ msg.OutMessage = (*MsgServerList)(nil)
var _ msg.OutMessage = (*MsgWaitInLine)(nil)
