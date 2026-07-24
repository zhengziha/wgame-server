package constant

// ClientButtonIdConst 对应 Java 端 com.fengshen.server.data.constant.ClientButtonIdConst
// 客户端按钮/通知ID常量

const (
	// 角色相关
	CharRename = 1 // 重命名
	DropTask   = 2 // 放弃任务

	// 排行相关
	GetRankInfo = 3 // 获取排行信息

	// 装备相关
	DeleteStoneAttrib     = 4 // 删除宝石属性
	DeleteGodbookSkill    = 5 // 删除天书技能
	EquipIdentify         = 7 // 装备鉴定
	EquipIdentifyGem      = 20031
	EquipIdentifyOk       = 20022
	EquipReformOk         = 49 // 装备改造成功
	EquipRefineOk         = 50 // 装备精炼成功
	EquipStrengthenOk     = 51 // 装备强化成功
	EquipResonanceOk      = 53 // 装备共鸣成功
	EquipUpgradeInheritOk = 54 // 装备升级继承成功
	EquipUpgradeOk        = 48 // 装备升级成功
	EquipEvolveOk         = 20027
	EquipDegenerationOk   = 20035
	HigherJewelryRecastOk = 20034
	UpgradeJewelryOk      = 10000
	JewelryRecastOk       = 20034

	// 守护相关
	CallGuard              = 6
	GuardUseSkillD         = 8
	GuardSaveGrow          = 9
	GuardGrowOk            = 47
	GuardNextFightScore    = 38
	GuardBasicAttrib       = 30038
	NextGuardInfo          = 30039
	RequestGuardId         = 30021
	RequestGuardExperience = 30022

	// 购买/兑换相关
	WhetherBuyItem      = 10
	WhetherExchangeCash = 11
	WhetherBuyGold      = 12

	// 好友相关
	RecommendFriend = 13
	VerifyFriend    = 14
	FetchMinfo      = 100

	// 角色信息
	GetCharInfo = 15

	// 状态/通知
	NotifyClientStatus   = 16
	NotifyAutoDisconnect = 20011
	CombatStatusInfo     = 20007

	// 双倍经验
	NotifyFetchDoublePoints  = 17
	NotifyFrozenDoublePoints = 18
	NotifyBuyDoublePoints    = 19
	NotifyEnableDoublePoints = 52

	// 战斗/战斗辅助
	NotifyStartAutoPractice = 20
	NotifyStartAutoFight    = 37
	NotifyAutoFightSkill    = 10004
	NotifyAutoFightLessMana = 10005

	// 竞技场
	NotifyOpenArena            = 21
	NotifyArenaTopBonusList    = 22
	NotifyFetchArenaRankBonus  = 23
	NotifyFetchArenaTimeBonus  = 24
	NotifyOpenArenaShop        = 25
	NotifyArenaChallenge       = 27
	NotifyArenaRefreshOpponent = 28
	NotifyArenaBuyTimes        = 29
	NotifyArenaRefreshShop     = 30
	NotifyArenaBuyItem         = 31

	// 活力
	NotifyGetLivenessInfo    = 32
	NotifyFetchLivenessBonus = 33

	// 宠物/宠物相关
	NotifyShowRankPet = 34
	NoticeBuyElitePet = 50001
	SubmitPet         = 30020

	// 初始化/系统通知
	NotifySendInitDataDone = 39 // 初始化数据发送完成

	// 帮派/门派相关
	NotifyLevelUpParty  = 41
	NotifyFinishAlchemy = 46
	NotifyOpenParty     = 99
	PartyWarScore       = 50002
	PartyWarInfo        = 50003
	PartySalary         = 20013
	PartyContributor    = 20014
	PartyShouwei        = 20015
	PartyHangbaRuqin    = 20016
	JoinPartyWar        = 30042
	StartBangpaiShouwei = 50004
	StartHanbaRuqin     = 50005

	// 商店/交易
	NotifyOpenStore      = 10002
	NotifyCloseStore     = 10003
	NotifyStallBatchNum  = 10011
	NotifyMailAllLoaded  = 10012 // 邮件全部加载完成
	NoticeGetItemSuccess = 20006

	// 技能/装备
	SkillStoneOk = 12000

	// 游戏逻辑
	NotifyOpenDlg  = 97
	NotifyCloseDlg = 98

	// 推荐属性/相性
	GetRecommendAttrib = 26
	GetRecommendPolar  = 44

	// 训练/练习
	GetExercise = 20000

	// 商城/活动
	NotifyQmoxiangStatus = 20010 // 驱魔香状态
	NotifyCloseExorcism  = 20009
	NotifyOpenExorcism   = 20008

	// 背包/仓库
	NoticeFetchBonus = 20005
	NoticeGetItem    = 20006

	// 任务/操作
	NoticeStopAutoWalk    = 20003
	NoticeUpdateMainIcon  = 20002
	NoticeOverInstruction = 20004

	// 排行榜
	RankMeInfo   = 30017
	RankGetGuard = 30018
	RankGetEquip = 30019

	// 队伍
	NotifyRequestMatchSize = 30013
	NotifyTeamAskAgree     = 30030
	NotifyTeamAskRefuse    = 30031
	NotifyQueryTeamInfo    = 20012
	NotifyQueryTeamExInfo  = 10008
	NotifyGetTeamData      = 30044
	ZoneHasNoTeamQuit      = 30008
	ZoneHasNoTeamConfirm   = 30009

	// 匹配
	NotifyCancelMatchLeader = 40024
	NotifyCancelMatchMember = 40025
	NotifyStartMatchMember  = 40026
	NotifyMatchTeamList     = 40023

	// 宠物/抽卡
	NotifyFetchSouchongGift  = 50009
	NotifyDrawLottery        = 50011
	NotifyCancelLottery      = 50013
	NotifyFetchLottery       = 50012
	NotifyFetchDone          = 50015
	NotifyRequestLotteryInfo = 50014

	// VIP/充值
	NotifyRequestRebateInfo  = 50010
	NotifyBuyInsider         = 50006
	NotifyDrawInsiderCoin    = 50007
	NotifyRequestInsiderInfo = 50008
	RechargeCoin             = 30015

	// 试炼/通天塔
	TTTGetBonus       = 40000
	TTTDoRevive       = 40001
	TTTJisuFeiSheng   = 40002
	TTTKuaisuFeiSheng = 40003
	TTTJumpAssure     = 30025
	TTTJumpCancel     = 30026
	TTTResetTask      = 40004
	TTTGoNextLayer    = 40006
	TTTLeaveTower     = 40007
	TTTDoReviveOk     = 40001
	TongttGetTask     = 30040

	// 八仙
	BaxianReset = 11001
	BaxianEnter = 11002

	// 宝石/改造
	JewelryRecast = 20034

	// 其他装备操作
	EquipIdentifyGemOk = 20031
	EquipEvolveOk2     = 20027

	// 洗练/重铸
	ShuadaoOpenInterface   = 30002
	ShuadaoSetOffline      = 30003
	ShuadaoBuyOfflineTime  = 30004
	ShuadaoDoBonus         = 30005
	ShuadaoSetJiji         = 30029
	ShuadaoSetChongfengan  = 30046
	BuyChongfengan         = 30047
	ShuadaoSetZiqihongmeng = 30048
	BuyZiqihongmeng        = 30049
	CloseOfflineShuadao    = 30016

	// 帮派/门派其他
	SetCombatGuard  = 30010
	RemoveAllInvite = 30011
	RemoveAllJoin   = 30012

	// 锁定经验
	SetLockExp = 30024

	// 商城/摆摊
	NoticeQueryCardInfo = 20001
	MountMergeResult    = 61000
	LookPlayerEquip     = 40005
	OpenNewbieGift      = 40014
	FetchNewbieGift     = 40013
	OpenDailySign       = 40009
	DoDailySign         = 40010
	OpenShenmiDali      = 40011
	OpenWelfare         = 40008
	OpenMyStall         = 40015
	StallRemoveGoods    = 40016
	StallRestartGoods   = 40017
	OpenStallList       = 40018
	StallSearchItem     = 40019
	StallOpenRecord     = 40020
	StallItemPrice      = 45
	StallQueryPrice     = 40021
	StallTakeCash       = 40022

	// 门派/角色操作
	RandomName         = 30007
	RequestBuybackCard = 50019
	BuyBack            = 50020
	DeleteChar         = 30032
	ResponseSecret     = 30033
	CancelDeleteChar   = 30034
	ConfirmResult      = 30037

	// 交易/市场
	MarketCard      = 30027
	MarketCheckGood = 30028

	// 宠物/宠物心法
	BuyJiji = 30045

	// 其他
	QueryShidaoInfo       = 20017
	BuyJiJ                = 30045
	CombatGetCurrentRound = 30041
	SubmitEquip           = 20024
	CharChangeSex         = 20025
	SubmitNanhws          = 20026
	FinishGather          = 20023
	EnableShenmuPoints    = 10009
	BuyShenmuPoints       = 10010
	SetUseMoneyType       = 30023
	BaoxiangReadySearch   = 10001
	CloseParty            = 10006
	FastAddExtra          = 10007
	AutoFightSkill        = 10004
	AutoFightLessMana     = 10005

	// 充值/礼包
	FetchRechargeGift     = 20018
	OpenRechargeGift      = 20019
	FetchLoginGift        = 20020
	OpenLoginGift         = 20021
	QueryPartySalary      = 20013
	QueryPartyContributor = 20014

	// 其他功能
	HigherJewelryRecastOK = 20034
	StallBatchNum         = 10011
	QueryShidaoInfo2      = 20017
	FriendClearXinmo      = 61004
	JoinParty2            = 99
	AssignXMD             = 50021
	TTTDLeaveTower        = 50022
	HideNPC               = 61003
	BuyHouseResult        = 61002

	// 御宝仙术相关
	NotifyTestYbxsEndTime = 60001 // 御宝仙术活动结束时间

	// iOS审核相关
	NotifyIOSReview = 50017 // iOS审核
)
