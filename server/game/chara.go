package game

import (
	"sync"
	"time"
)

// Chara 表示玩家运行时数据。
// 参考 Java wd-server-fl core/domain/Chara.java。
//
// 注意：这是内存态对象，不是数据库实体。数据库实体是 model.Characters。
// 登录时从 model.Characters 加载到 Chara，离线时写回。
type Chara struct {
	// 基础标识
	ID    int32  // 数据库主键 characters.id
	Gid   string // 全局唯一 id（客户端使用）
	Name  string
	Sex   int32
	Level int32
	Polar int32 // 门派/阵营

	// 位置信息
	MapId   int32
	MapName string
	Line    int32 // 线路号（1-based）
	X, Y    int32
	Dir     int32 // 朝向

	// 移动状态
	MoveType int32 // 0=正常行走, 1=飞行等
	FlyType  int32
	MoveIds  []int32

	// 队伍信息
	TeamId       int32
	IsTeamLeader bool

	// 房屋信息
	HouseType int32
	HouseId   int32

	// 显示相关
	IsHide      int32
	Opacity     int32
	Camp        int32
	Title       string
	ShieldOther bool // 是否屏蔽周围玩家

	// 宠物/跟宠
	FollowPet          int32
	FlowerChild        int32
	FlowerChildVisible int32
	GenchongIcon       int32

	// 坐骑/外观
	CscwQiaozhuang int32

	// 时间戳
	LastUpdate time.Time

	// ======== 扩展字段 ========

	// 飞宠物
	FlyPetID   int32
	FlyChildID int32

	// 修法npc名字
	XiufaNpcName string

	// 签名
	Zdd_Notice   string // 证道殿签名
	Yxh_Notice   string // 英雄榜签名
	LeaderNotice string // 掌门签名

	// 首饰精华
	JewelryEssence int32

	// 物品列表
	HomeStore       []*Goods
	OtherGoods      []*Goods // 当前穿戴 pos:1-40
	Backpack        []*Goods // 背包列表
	Cangku          []*Goods // 仓库
	Shizhuang       []*Goods // 时装
	CustomShizhuang []*Goods // 自定义时装
	TeamIconStore   []*Goods // 队标仓库
	Texiao          []*Goods // 特效仓库
	Genchong        []*Goods // 跟从仓库
	CardStore       []*Goods // 变身卡仓库
	TaiYinCangku    []*Goods // 太阴之气仓库
	WuhunCangku     []*Goods // 武魂仓库

	// 宠物相关
	Pets       []*Petbeibao // 宠物列表
	Childs     []*Petbeibao // 娃娃列表
	ListShouhu []*ShouHu    // 守护列表

	// 装备属性
	ZbAttribute *ZbAttribute

	// 刷道阶数
	TaoStage int32

	// 结拜称号
	BrotherAppellation string

	// 领取首冲礼包
	FetchShouchongGift int32

	// 称号墙
	AppellationWall []string

	// 自动行走
	AutoWalk bool

	// 时装标签 1自定义时装,0时装
	FasionLabel int32

	// 状态
	Status int32

	// 充值积分
	ChargeScore int32

	// 好心值
	Nice int32

	// 技能列表
	JiNengList []*JiNeng

	// 外观
	Waiguan int32

	// 当前任务
	CurrentTask    string // 主线
	ZurenweileTask string // 助人为乐
	FlyTask        string // 宠物飞升
	YaomodaoTask   string // 妖魔道

	// 免费改名次数
	FreeRename int32

	// 炼器信息
	XiLianInfoMap map[string]int32

	// 坐骑助阵
	MountPetCheerName     string
	MountPetCheerAddPolar int32
	MountPetCheerIcon     int32

	// 武魂之尘
	WhzcNum int32

	// 魂窍等级
	WhqLevels [5]int32

	// 武器外观
	WeaponIcon int32

	// 金币使用方式
	UseMoneyType int32

	// 代金券
	Voucher int32

	// 充值抽奖次数
	ShadowSelf int32

	// 银元宝
	SilverCoin    int32
	SilverCoinExt int64 // 溢出的银元宝

	// 金元宝
	GoldCoin    int32
	GoldCoinExt int64 // 溢出的金元宝

	// 金币
	Cash int32

	// 阴气之尘
	YqzcNum int32

	// 钱庄存款
	Balance int32

	// 聚财箱存款
	JcxBalance int32

	// 集市存款
	JishouCoin int64

	// 锁定经验
	LockExp int32

	// 经验
	Jinyan int32

	// 潜能
	Pot int32

	// 注册时间
	AddTime time.Time

	// 登录上线时间
	Uptime     int64
	UpdateTime int64

	// 签到天数
	SignDays int32

	// 参战守护数量
	Canzhanshouhunumber int32

	// 坐骑图标
	PetIcon     int32
	MountIcon   int32
	SpecialIcon int32

	// VIP
	VipType        int32
	VipTime        int32
	VipTimeShengYu int32

	// 套装特效
	SuitIcon        int32
	SuitLightEffect int32

	// 五行竞猜金钱
	WuxingBalance int32

	// 道行(单位天)
	Tao int32

	// 道行点(每1440点才是1天)
	TaoPoint int32

	// 月道行(天)
	MonthTao   int32
	LastMonTao int32 // 上月道行

	// 重置水位
	LastDailyReset  int32
	LastWeeklyReset int32

	// 称号map
	Chenghao        map[string]string
	ChenghaoTimeMap map[string]int64

	// 驱魔香
	Qumoxiang int32

	// 双倍点数
	DoublePoints int32

	// 神木鼎点数
	ShenmuPoints int32

	// 宠风散
	ChongFengSanPoints int32
	ChongfengsanState  int32

	// 双倍状态
	ShuangbeiState int32

	// 神木鼎状态
	ShenmodingState int32

	// 紫气鸿蒙
	ZiqihongmengPoints int32
	ZiqihongmengState  int32

	// 急急如律令
	JijirulvlingPoints int32
	JijirulvlingState  int32

	// 如意刷道令
	RuyishuadaoState  int32
	RuyiAmtState      int32
	RuyishuadaoPoints int32

	// 新手礼包
	Xinshoulibao [8]int32

	// 剧本索引
	NextJuBen     int32
	EndJuBen      int32
	CurrentJuBens []string
	JubenAllTeam  bool

	// 战斗状态
	IsFight bool

	// 挑战掌门boss名称
	ZhandouInfo string

	// 当前确认物品
	CurrentConfirmItem string

	// 自动加点配置
	UserAutoAddPointList      []map[string]int32
	UserAutoAddPolarPointList []map[string]int32
	UserAutoAddPoint2         map[string]interface{}
	UserAutoAddPolarPoint2    map[string]interface{}
	ChildAutoAddPoint         map[string]int32

	// 设置
	Settings map[string]int32

	// 设置拒绝陌生人等级
	SettingrefuseStrangerLevel int32
	SettingautoReplyMsg        string
	SettingRefuseBeAddLevel    int32
	ToVerifyFriendGid          string

	// 珍宝账户金额
	SellCash int32

	// 飞升相关
	IsFeisheng           int32
	UpgradeType          int32 // 0=未飞升,1=元婴,2=血婴,3=仙道,4=魔道
	UpgradeState         int32 // 元婴状态1启用0真身
	UpgradeLevel         int32
	UpgradeExp           int32
	UpgradeMaxPolarExtra int32
	UpgradeAddType       int32
	UpgradeIsOpen        int32
	UpgradeImmortal      int32 // 仙道点
	UpgradeMagic         int32 // 魔道点
	UpgradeTotal         int32 // 剩余仙魔点

	// 帮派相关
	PartyName   string
	PartyJob    string
	UpPartyName string // 上次帮派名称
	Contrib     int32  // 帮贡

	// 变身卡
	CardSize int32

	// 家族
	Family string

	// 红名
	IsNameRed int32

	// PK值
	ForcePk int32

	// 当前装备页
	EquipPage   int32
	HunqiaoPage int32

	// 自定义图标
	CustomIcon  string
	EffectIcons map[string]int32

	// 完成100任务
	IsFinish100Task int32

	// 聊天
	UseChatFloor string
	UseChatHead  string

	// 结婚
	MarriageMarryId   int32
	MarriageCoupleGid string
	MarriageName      string
	MarriageIcon      int32
	MarriageTime      int64

	// 附灵
	ZhenlingType  int32
	ZhenlingLevel int32
	ZhenlingStage int32
	ZhenlingExp   int32
	ZhenlingPhy   int32
	ZhenlingMag   int32
	ZhenlingSpeed int32
	ZhenlingDef   int32

	// 四象附灵
	QinglongZhenlingLevel int32
	BaihuhenlingLevel     int32
	ZhuqueZhenlingLevel   int32
	XuanwuZhenlingLevel   int32

	// 自动战斗
	AutoFight                int32
	AutofightSelect          int32 // 魔法不足,1自动使用药品,2普攻
	AutofightSkillaction     int32
	AutofightMultiIndex      int32
	AutofightSkillno         int32
	AutofightSkillGroupIndex int32
	FixGroupSkill            int32

	// 阵营战
	YuZhu    int32 // 玉珠
	ZhanGong int32 // 战功

	// 角色基础属性
	MaxShengming int32 // 最大生命
	MaxMofa      int32 // 最大魔法
	PhyPower     int32 // 物理攻击
	MagPower     int32 // 魔法攻击
	Def          int32 // 防御
	Speed        int32 // 速度

	// 属性点
	Str         int32
	Dex         int32
	Con         int32
	Wiz         int32
	AttribPoint int32
	PolarPoint  int32

	// 五行属性
	Metal int32
	Wood  int32
	Water int32
	Fire  int32
	Earth int32

	// 经验
	Exp            int64
	ExpToNextLevel int32

	// 忠诚度
	BackupLoyalty int32

	// 当前生命/魔法
	Shengming int32
	Mofa      int32

	// 额外属性
	ExtraMana int32
	ExtraLife int32

	// 充值总额
	ChargeTotal int32

	// 战斗相关
	FightCap      int32
	FightCapTotal int32

	// 活跃度相关
	Baibangmang    int32
	Fabaorenwu     int32
	Xiuxingcishu   int32
	XiuxingNpcname string
	XiulianNpcname string
	Xuanshangcishu int32

	// 刷道相关
	Shuadao           int32
	Shidaodaguaijifen int32

	// 证道殿相关
	Zhengdaodiancishu int32

	// 英雄榜相关
	Heropubcishu int32

	// 封神相关
	Gongchengcishu int32
	Zhanshencishu  int32

	// 海盗相关
	Haidaocishu int32

	// 上古相关
	Shanggucishu int32

	// 万年相关
	Wanniancishu int32

	// 修炼相关
	Xiufacishu int32

	// 地图守卫
	Mapguardcishu int32

	// 天/地劫
	TiantaixingNum int32

	// 刷道经验
	ShidaoExp     int32
	ShidaoTao     int32
	ShidaoMartial int32

	// 跨服相关
	Tongttcishu     int32
	TongttRestcishu int32

	// 掌门挑战
	Zhangmenshijiantime string
	Zhangmentiaozhan    int32

	// 八仙
	BaxianTimes int32

	// 副本
	FbNum int32

	// 神魂
	ShenHunDataSate int32
	ShenHunPhyPower int32
	ShenHunDef      int32
	ShenHunmaxLife  int32
	ShenHunMagPower int32
	ShenHunSpeed    int32

	// 洛书
	LuoshuMagpower int32
	LuoshuPhypower int32
	LuoshuDefense  int32
	LuoshuSpeed    int32
	LuoshuExp      int32

	// 刷道最高记录
	ShuadaoHigestRecord string

	// 擂台积分
	LeitaiScore int32

	// 阵法
	Chubao int32

	// 属性计划
	AttribPlan    int32
	AttribPlanMap map[int32]*AttribPlan

	// 组队飞行
	ShareMountIcon     int32
	ShareMountLeaderId int32
	GatherMountIcons   []int32
	GatherNames        []string

	// 变身卡信息
	ChangeCardInfo *ChangeCardInfo

	// 是否可以签到
	IsCanSgin int32

	mu sync.RWMutex
}

// AttribPlan 表示属性计划
type AttribPlan struct {
	Con          int32
	Wiz          int32
	Str          int32
	Dex          int32
	AttribPoint  int32
	Metal        int32
	Wood         int32
	Water        int32
	Fire         int32
	Earth        int32
	PolarPoint   int32
	UpgradeTotal int32
}

// ChangeCardInfo 表示变身卡信息
type ChangeCardInfo struct {
	CardId   int32
	Duration int32
}

// NewChara 创建一个新的玩家运行时对象（用于角色创建）
func NewChara(name string, sex, polar int32, gid string) *Chara {
	c := &Chara{
		Name:                name,
		Sex:                 sex,
		Polar:               polar,
		Gid:                 gid,
		Level:               1,
		MapId:               1000,
		MapName:             "揽仙镇",
		Line:                1,
		X:                   22,
		Y:                   108,
		Dir:                 6,
		MoveType:            0,
		IsHide:              0,
		ShieldOther:         false,
		TaoStage:            1,
		BrotherAppellation:  "",
		FetchShouchongGift:  0,
		AppellationWall:     []string{},
		AutoWalk:            false,
		FasionLabel:         0,
		Status:              0,
		ChargeScore:         0,
		Nice:                0,
		CurrentTask:         "主线—浮生若梦_s0",
		FreeRename:          0,
		WeaponIcon:          0,
		UseMoneyType:        0,
		Voucher:             100,
		ShadowSelf:          0,
		SilverCoin:          0,
		GoldCoin:            0,
		Cash:                0,
		LockExp:             0,
		Jinyan:              0,
		Pot:                 0,
		SignDays:            0,
		Canzhanshouhunumber: 0,
		PetIcon:             0,
		MountIcon:           0,
		SpecialIcon:         0,
		GenchongIcon:        0,
		VipType:             0,
		Tao:                 0,
		TaoPoint:            0,
		MonthTao:            0,
		Qumoxiang:           0,
		DoublePoints:        0,
		ShenmuPoints:        0,
		ChongFengSanPoints:  0,
		ChongfengsanState:   0,
		ShuangbeiState:      0,
		ShenmodingState:     0,
		ZiqihongmengPoints:  0,
		ZiqihongmengState:   0,
		JijirulvlingPoints:  0,
		JijirulvlingState:   0,
		RuyishuadaoState:    0,
		RuyishuadaoPoints:   0,
		Xinshoulibao:        [8]int32{},
		NextJuBen:           0,
		JubenAllTeam:        false,
		IsFight:             false,
		ZhandouInfo:         "",
		CurrentConfirmItem:  "",
		Settings: map[string]int32{
			"verify_be_added":   0,
			"friend_msg_bubble": 0,
		},
		SellCash:             0,
		IsFeisheng:           0,
		UpgradeType:          0,
		UpgradeState:         0,
		UpgradeLevel:         0,
		PartyName:            "",
		PartyJob:             "",
		UpPartyName:          "",
		Contrib:              0,
		CardSize:             15,
		Family:               "",
		IsNameRed:            0,
		ForcePk:              0,
		EquipPage:            1,
		HunqiaoPage:          1,
		CustomIcon:           "",
		IsFinish100Task:      0,
		UseChatFloor:         "",
		UseChatHead:          "",
		MarriageName:         "",
		MarriageTime:         0,
		ZhenlingType:         0,
		AutoFight:            0,
		AutofightSelect:      1,
		AutofightSkillaction: 2,
		AutofightSkillno:     2,
		YuZhu:                0,
		ZhanGong:             0,

		// 基础属性初始化
		MaxShengming:   105,
		MaxMofa:        84,
		PhyPower:       45,
		MagPower:       45,
		Def:            45,
		Speed:          50,
		Str:            1,
		Dex:            1,
		Con:            1,
		Wiz:            1,
		AttribPoint:    0,
		PolarPoint:     0,
		Metal:          0,
		Wood:           0,
		Water:          0,
		Fire:           0,
		Earth:          0,
		Exp:            0,
		ExpToNextLevel: 517,
		BackupLoyalty:  300,
		Shengming:      159,
		Mofa:           84,
		ExtraMana:      1000000,
		ExtraLife:      1000000,
		FightCap:       0,
		FightCapTotal:  0,

		// 活跃度
		Baibangmang:    0,
		Fabaorenwu:     0,
		Xiuxingcishu:   1,
		Xuanshangcishu: 0,

		// 刷道
		Shuadao: 1,

		// 神魂
		ShenHunDataSate: 1,

		// 阵法
		Chubao: 1,

		// 签到
		IsCanSgin: 1,
	}

	// 初始化列表
	c.HomeStore = []*Goods{}
	c.OtherGoods = []*Goods{}
	c.Backpack = []*Goods{}
	c.Cangku = []*Goods{}
	c.Shizhuang = []*Goods{}
	c.CustomShizhuang = []*Goods{}
	c.TeamIconStore = []*Goods{}
	c.Texiao = []*Goods{}
	c.Genchong = []*Goods{}
	c.CardStore = []*Goods{}
	c.TaiYinCangku = []*Goods{}
	c.WuhunCangku = []*Goods{}
	c.Pets = []*Petbeibao{}
	c.Childs = []*Petbeibao{}
	c.ListShouhu = []*ShouHu{}
	c.JiNengList = []*JiNeng{}
	c.XiLianInfoMap = map[string]int32{}
	c.Chenghao = map[string]string{}
	c.ChenghaoTimeMap = map[string]int64{}
	c.EffectIcons = map[string]int32{}
	c.UserAutoAddPointList = []map[string]int32{}
	c.UserAutoAddPolarPointList = []map[string]int32{}
	c.ChildAutoAddPoint = map[string]int32{}
	c.AttribPlanMap = map[int32]*AttribPlan{}
	c.GatherMountIcons = []int32{}
	c.GatherNames = []string{}

	// 装备属性
	c.ZbAttribute = NewZbAttribute()

	// 根据门派和性别设置外观
	c.WaiguanByPolarAndSex()

	return c
}

// ======== 核心方法 ========

// MaxMana 返回最大魔法值
func (c *Chara) MaxMana() int32 {
	return c.MaxMofa
}

// MaxLife 返回最大生命值
func (c *Chara) MaxLife() int32 {
	return c.MaxShengming
}

// MagPowerTotal 返回魔法攻击（含装备加成）
func (c *Chara) MagPowerTotal() int32 {
	if c.ZbAttribute == nil {
		return c.MagPower
	}
	return c.MagPower + c.ZbAttribute.MagPower
}

// PhyPowerTotal 返回物理攻击（含装备加成）
func (c *Chara) PhyPowerTotal() int32 {
	if c.ZbAttribute == nil {
		return c.PhyPower
	}
	return c.PhyPower + c.ZbAttribute.PhyPower
}

// DefTotal 返回防御（含装备加成）
func (c *Chara) DefTotal() int32 {
	if c.ZbAttribute == nil {
		return c.Def
	}
	return c.Def + c.ZbAttribute.Def
}

// SpeedTotal 返回速度（含装备加成）
func (c *Chara) SpeedTotal() int32 {
	if c.ZbAttribute == nil {
		return c.Speed
	}
	return c.Speed + c.ZbAttribute.Speed
}

// HasCoin 检查是否有足够的元宝（银元宝+金元宝）
func (c *Chara) HasCoin(num int64) bool {
	return int64(c.SilverCoin)+int64(c.GoldCoin) >= num
}

// SubSilverCoin 扣除元宝（先扣银元宝，不足用金元宝补）
func (c *Chara) SubSilverCoin(num int64) bool {
	if !c.HasCoin(num) {
		return false
	}
	if int64(c.SilverCoin) < num {
		c.GoldCoin -= int32(num - int64(c.SilverCoin))
		c.SilverCoin = 0
	} else {
		c.SilverCoin -= int32(num)
	}
	return true
}

// AddCash 添加金币
func (c *Chara) AddCash(num int64) {
	if c.Cash < 0 {
		c.Cash = 0
	}
	if int64(c.Cash)+num > 2000000000 {
		c.Cash = 2000000000
	} else {
		c.Cash += int32(num)
		if c.Cash < 0 {
			c.Cash = 0
		}
	}
}

// AddPot 添加潜能
func (c *Chara) AddPot(num int64) bool {
	if int64(c.Pot)+num > 2000000000 {
		c.Pot = 2000000000
	} else {
		c.Pot += int32(num)
	}
	return true
}

// AddExp 添加经验
func (c *Chara) AddExp(num int64) {
	if int64(c.Exp)+num > 2000000000 {
		c.Exp = 2000000000
	} else {
		c.Exp += num
		if c.Exp < 0 {
			c.Exp = 0
		}
	}
}

// GetSetting 获取设置
func (c *Chara) GetSetting(key string) int32 {
	if c.Settings == nil {
		c.Settings = map[string]int32{}
	}
	if val, ok := c.Settings[key]; ok {
		return val
	}
	return 0
}

// PutSetting 设置配置
func (c *Chara) PutSetting(key string, val int32) {
	if c.Settings == nil {
		c.Settings = map[string]int32{}
	}
	c.Settings[key] = val
}

// WaiguanByPolarAndSex 根据门派和性别设置外观
func (c *Chara) WaiguanByPolarAndSex() {
	switch {
	case c.Polar == 11 && c.Sex == 1:
		c.Waiguan = 20033
	case c.Polar == 1 && c.Sex == 1:
		c.Waiguan = 6001
	case c.Polar == 2 && c.Sex == 1:
		c.Waiguan = 7002
	case c.Polar == 3 && c.Sex == 1:
		c.Waiguan = 7003
	case c.Polar == 4 && c.Sex == 1:
		c.Waiguan = 6004
	case c.Polar == 5 && c.Sex == 1:
		c.Waiguan = 6005
	case c.Polar == 1 && c.Sex == 2:
		c.Waiguan = 7001
	case c.Polar == 2 && c.Sex == 2:
		c.Waiguan = 6002
	case c.Polar == 3 && c.Sex == 2:
		c.Waiguan = 6003
	case c.Polar == 4 && c.Sex == 2:
		c.Waiguan = 7004
	case c.Polar == 5 && c.Sex == 2:
		c.Waiguan = 7005
	}
}
