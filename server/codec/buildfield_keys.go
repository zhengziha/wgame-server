package codec

// BuildField key 常量定义
// 迁移自 Java 端 BuildFieldsNew.java
// 采用集中定义的方式，避免在 tag 中使用"魔法数字"
//
// 使用方式：
//   1. 在 codec tag 中使用常量名: `codec:"bf:LeftTimeToDelete"`
//   2. 或直接使用常量值: `codec:"bf:263"`（不推荐，仅兼容）

const (
	// === 基础属性 ===
	BFKeyName              int16 = 1  // 角色名
	BFKeyStr               int16 = 2  // 字符串
	BFKeyPhyPower          int16 = 3  // 物理攻击
	BFKeyAccurate          int16 = 4  // 准确
	BFKeyCon               int16 = 5  // 体质
	BFKeyLife              int16 = 6  // 生命
	BFKeyMaxLife           int16 = 7  // 最大生命
	BFKeyDef               int16 = 8  // 防御
	BFKeyWiz               int16 = 9  // 智力
	BFKeyMagPower          int16 = 10 // 法术攻击
	BFKeyMana              int16 = 11 // 内力
	BFKeyMaxMana           int16 = 12 // 最大内力
	BFKeyDex               int16 = 13 // 敏捷
	BFKeySpeed             int16 = 14 // 速度
	BFKeyParry             int16 = 15 // 格挡
	BFKeyAttribPoint       int16 = 16 // 属性点
	BFKeyPolarPoint        int16 = 17 // 五行点
	BFKeyStamina           int16 = 18 // 体力
	BFKeyMaxStamina        int16 = 19 // 最大体力
	BFKeyTao               int16 = 20 // 道行
	BFKeyFriend            int16 = 21 // 友好度
	BFKeyTotalPk           int16 = 23 // 总PK数
	BFKeyDegree            int16 = 24 // 等级
	BFKeyExp               int16 = 25 // 经验
	BFKeyPot               int16 = 26 // 元宝
	BFKeyCash              int16 = 27 // 现金
	BFKeyBalance           int16 = 28 // 余额
	BFKeyGender            int16 = 29 // 性别
	BFKeyMaster            int16 = 30 // 师傅
	BFKeyLevel             int16 = 31 // 等级
	BFKeySkill             int16 = 32 // 技能
	BFKeyPartyName         int16 = 33 // 帮派名
	BFKeyPartyContrib      int16 = 34 // 帮派贡献
	BFKeyNick              int16 = 35 // 昵称
	BFKeyTitle             int16 = 36 // 称号
	BFKeyNice              int16 = 37 // 好感度
	BFKeyReputation        int16 = 38 // 声望
	BFKeyCouple            int16 = 39 // 配偶
	BFKeyIcon              int16 = 40 // 图标
	BFKeyType              int16 = 41 // 类型
	BFKeyDurability        int16 = 42 // 耐久
	BFKeyMaxDurability     int16 = 43 // 最大耐久
	BFKeyPolar             int16 = 44 // 门派
	BFKeyMetal             int16 = 45 // 金
	BFKeyWood              int16 = 46 // 木
	BFKeyWater             int16 = 47 // 水
	BFKeyFire              int16 = 48 // 火
	BFKeyEarth             int16 = 49 // 土
	BFKeyResistMetal       int16 = 50 // 抗金
	BFKeyResistWood        int16 = 51 // 抗木
	BFKeyResistWater       int16 = 52 // 抗水
	BFKeyResistFire        int16 = 53 // 抗火
	BFKeyResistEarth       int16 = 54 // 抗土
	BFKeyExpToNextLevel    int16 = 55 // 升级所需经验
	BFKeyResistPoison      int16 = 56 // 抗毒
	BFKeyResistFrozen      int16 = 57 // 抗冰冻
	BFKeyResistSleep       int16 = 58 // 抗睡眠
	BFKeyResistForgotten   int16 = 59 // 抗遗忘
	BFKeyResistConfusion   int16 = 60 // 抗混乱
	BFKeyLongevity         int16 = 61 // 寿命
	BFKeyMartial           int16 = 62 // 武学
	BFKeyIntimacy          int16 = 63 // 亲密度
	BFKeyShape             int16 = 64 // 形态
	BFKeyResistPoint       int16 = 65 // 抗点
	BFKeyLoyalty           int16 = 66 // 忠诚度
	BFKeyDoubleHit         int16 = 67 // 双击
	BFKeyStunt             int16 = 68 // 特技
	BFKeyCounterAttack     int16 = 69 // 反击
	BFKeyLifeRecover       int16 = 70 // 生命恢复
	BFKeyManaRecover       int16 = 71 // 内力恢复
	BFKeyLifeRecoverRate   int16 = 72 // 生命恢复率
	BFKeyManaRecoverRate   int16 = 73 // 内力恢复率
	BFKeyItemType          int16 = 74 // 物品类型
	BFKeyTotalScore        int16 = 75 // 总分
	BFKeyCounterAttackRate int16 = 77 // 反击率
	BFKeyDoubleHitRate     int16 = 78 // 双击率
	BFKeyStuntRate         int16 = 79 // 特技率
	BFKeyDamageSel         int16 = 80 // 伤害选择
	BFKeyFamily            int16 = 81 // 家族
	BFKeyReqStr            int16 = 82 // 需求力量
	BFKeyTotalDied         int16 = 83 // 总死亡次数
	BFKeyItemUnique        int16 = 84 // 物品唯一
	BFKeyDamageSelRate     int16 = 85 // 伤害选择率
	BFKeyPortrait          int16 = 86 // 头像
	BFKeyPassiveMode       int16 = 87 // 被动模式
	BFKeyStaticMode        int16 = 88 // 静态模式
	BFKeySource            int16 = 89 // 来源
	BFKeySignature         int16 = 90 // 签名

	// === 性格/角色详情 ===
	BFKeyCharacterHarmony      int16 = 91  // 和谐
	BFKeyCharacterKindness     int16 = 92  // 善良
	BFKeyCharacterCarefulness  int16 = 93  // 谨慎
	BFKeyCharacterCourage      int16 = 94  // 勇气
	BFKeyCharacterDesc         int16 = 95  // 角色描述
	BFKeyPartyDesc             int16 = 96  // 帮派描述
	BFKeyPartyJob              int16 = 97  // 帮派职位
	BFKeyFamilyTitle           int16 = 98  // 家族称号
	BFKeyPartyTitle            int16 = 99  // 帮派称号
	BFKeyApprenticeTitle       int16 = 100 // 徒弟称号
	BFKeyReqCon                int16 = 101 // 需求体质
	BFKeyReqWiz                int16 = 102 // 需求智力
	BFKeyReqDex                int16 = 103 // 需求敏捷
	BFKeyPetLifeShape          int16 = 104 // 宠物生命形态
	BFKeyPetManaShape          int16 = 105 // 宠物体力形态
	BFKeyPetSpeedShape         int16 = 106 // 宠物速度形态
	BFKeyPetPhyShape           int16 = 107 // 宠物物理形态
	BFKeyPetMagShape           int16 = 108 // 宠物法术形态
	BFKeyRank                  int16 = 109 // 排名
	BFKeyPenetrate             int16 = 110 // 穿透
	BFKeySpecialIcon           int16 = 116 // 特殊图标
	BFKeyPenetrateRate         int16 = 117 // 穿透率
	BFKeyNimbus                int16 = 118 // 灵气
	BFKeySilverCoin            int16 = 119 // 银币
	BFKeyGoldCoin              int16 = 120 // 金币
	BFKeyExtraLife             int16 = 121 // 额外生命
	BFKeyExtraMana             int16 = 122 // 额外内力
	BFKeyHaveCoinPwd           int16 = 123 // 是否有金币密码
	BFKeyMaxCash               int16 = 124 // 最大现金
	BFKeyMaxBalance            int16 = 125 // 最大余额
	BFKeyIgnoreResistMetal     int16 = 126 // 忽略抗金
	BFKeyIgnoreResistWood      int16 = 127 // 忽略抗木
	BFKeyIgnoreResistWater     int16 = 128 // 忽略抗水
	BFKeyIgnoreResistFire      int16 = 129 // 忽略抗火
	BFKeyIgnoreResistEarth     int16 = 130 // 忽略抗土
	BFKeyIgnoreResistForgotten int16 = 131 // 忽略抗遗忘
	BFKeyIgnoreResistPoison    int16 = 132 // 忽略抗毒
	BFKeyIgnoreResistFrozen    int16 = 133 // 忽略抗冰冻
	BFKeyIgnoreResistSleep     int16 = 134 // 忽略抗睡眠
	BFKeyIgnoreResistConfusion int16 = 135 // 忽略抗混乱
	BFKeySuperExcluseMetal     int16 = 136 // 超级排除金
	BFKeySuperExcluseWood      int16 = 137 // 超级排除木
	BFKeySuperExcluseWater     int16 = 138 // 超级排除水
	BFKeySuperExcluseFire      int16 = 139 // 超级排除火
	BFKeySuperExcluseEarth     int16 = 140 // 超级排除土
	BFKeyBSkillLowCost         int16 = 141 // B技能低消耗
	BFKeyCSkillLowCost         int16 = 142 // C技能低消耗
	BFKeyDSkillLowCost         int16 = 143 // D技能低消耗
	BFKeySuperPoison           int16 = 144 // 超级毒
	BFKeySuperSleep            int16 = 145 // 超级睡眠
	BFKeySuperForgotten        int16 = 146 // 超级遗忘
	BFKeySuperConfusion        int16 = 147 // 超级混乱
	BFKeySuperFrozen           int16 = 148 // 超级冰冻
	BFKeyEnhancedMetal         int16 = 149 // 增强金
	BFKeyEnhancedWood          int16 = 150 // 增强木
	BFKeyEnhancedWater         int16 = 151 // 增强水
	BFKeyEnhancedFire          int16 = 152 // 增强火
	BFKeyEnhancedEarth         int16 = 153 // 增强土
	BFKeyMagDodge              int16 = 154 // 法术闪避

	// === 反击率（特定技能） ===
	BFKeyJinguangZhaxianCounterAttRate int16 = 155 // 金光乍现反击率
	BFKeyZhaiyeFeihuaCounterAttRate    int16 = 156 // 摘叶飞花反击率
	BFKeyDishuiChuanshiCounterAttRate  int16 = 157 // 滴水穿石反击率
	BFKeyJuhuoFentianCounterAttRate    int16 = 158 // 举火焚天反击率
	BFKeyLuotuFeiyanCounterAttRate     int16 = 159 // 落土飞岩反击率

	// === 五行法术攻击 ===
	BFKeyMetalMagPower int16 = 160 // 金系法术攻击
	BFKeyWoodMagPower  int16 = 161 // 木系法术攻击
	BFKeyWaterMagPower int16 = 162 // 水系法术攻击
	BFKeyFireMagPower  int16 = 163 // 火系法术攻击
	BFKeyEarthMagPower int16 = 164 // 土系法术攻击

	// === 卡片系统 ===
	BFKeyCardLevel             int16 = 165 // 卡片等级
	BFKeyMaxCardAmount         int16 = 166 // 最大卡片数量
	BFKeyCardAmount            int16 = 167 // 卡片数量
	BFKeyIgnoreAllResistPolar  int16 = 168 // 忽略所有抗五行
	BFKeyIgnoreAllResistExcept int16 = 169 // 忽略所有抗除外
	BFKeyReleaseForgotten      int16 = 170 // 解除遗忘
	BFKeyReleasePoison         int16 = 171 // 解除毒
	BFKeyReleaseFrozen         int16 = 172 // 解除冰冻
	BFKeyReleaseSleep          int16 = 173 // 解除睡眠
	BFKeyReleaseConfusion      int16 = 174 // 解除混乱
	BFKeyTaoEx                 int16 = 175 // 道行额外
	BFKeyOwnerName             int16 = 176 // 拥有者名
	BFKeyBackupLoyalty         int16 = 178 // 备份忠诚度
	BFKeyUseSkillD             int16 = 179 // 使用技能D
	BFKeyMaxDegree             int16 = 180 // 最大等级
	BFKeyCostNum               int16 = 181 // 消耗数量
	BFKeyConMax                int16 = 182 // 体质最大
	BFKeyStrMax                int16 = 183 // 力量最大
	BFKeyWizMax                int16 = 184 // 智力最大
	BFKeyDexMax                int16 = 185 // 敏捷最大
	BFKeyConAdd                int16 = 186 // 体质加成
	BFKeyStrAdd                int16 = 187 // 力量加成
	BFKeyWizAdd                int16 = 188 // 智力加成
	BFKeyDexAdd                int16 = 189 // 敏捷加成
	BFKeyPracticeTimes         int16 = 190 // 修炼次数
	BFKeyDoublePoints          int16 = 191 // 双倍点数
	BFKeyEnableDoublePoints    int16 = 192 // 启用双倍点数
	BFKeyCanBuyDpTimes         int16 = 193 // 可购买双倍次数
	BFKeyOnline                int16 = 194 // 在线状态
	BFKeyArenaRank             int16 = 195 // 竞技场排名
	BFKeyPartyName2            int16 = 196 // 帮派名（重复）
	BFKeyUnidentified          int16 = 197 // 未鉴定
	BFKeyDegree32              int16 = 198 // 等级32位
	BFKeyStoreExp              int16 = 199 // 存储经验
	BFKeyEquipIdentify         int16 = 200 // 装备鉴定
	BFKeyDesc                  int16 = 201 // 描述
	BFKeyEquipType             int16 = 202 // 装备类型
	BFKeyAmount                int16 = 203 // 数量
	BFKeyOwnerId               int16 = 204 // 拥有者ID
	BFKeyReqLevel              int16 = 205 // 需求等级
	BFKeyAttrib                int16 = 206 // 属性
	BFKeyValue                 int16 = 207 // 值
	BFKeyRebuildLevel          int16 = 208 // 重建等级
	BFKeyColor                 int16 = 209 // 颜色
	BFKeyQuality               int16 = 210 // 品质
	BFKeyUseTimes              int16 = 211 // 使用次数
	BFKeyMaxUseTimes           int16 = 212 // 最大使用次数
	BFKeyCarpetRadius          int16 = 213 // 地毯半径
	BFKeyEquipPage             int16 = 214 // 装备页
	BFKeyMaxReqLevel           int16 = 215 // 最大需求等级
	BFKeyCreateTime            int16 = 216 // 创建时间

	// === CT数据 ===
	BFKeyCtDataScore     int16 = 217 // CT分数
	BFKeyCtDataTopRank   int16 = 218 // CT排名
	BFKeyRealDesc        int16 = 219 // 真实描述
	BFKeyAllAttrib       int16 = 220 // 所有属性
	BFKeyAllPolar        int16 = 221 // 所有五行
	BFKeyAllResistPolar  int16 = 222 // 所有抗五行
	BFKeyAllResistExcept int16 = 223 // 所有抗除外
	BFKeyAllSkill        int16 = 224 // 所有技能
	BFKeyMstuntRate      int16 = 225 // M特技率

	// === 技能升级 ===
	BFKeySkillDevelopExp      int16 = 230 // 技能升级经验
	BFKeySkillDevelopIntimacy int16 = 231 // 技能升级亲密度
	BFKeyLifeEffect           int16 = 232 // 生命效果
	BFKeyManaEffect           int16 = 233 // 内力效果
	BFKeyAttackEffect         int16 = 234 // 攻击效果
	BFKeySpeedEffect          int16 = 235 // 速度效果
	BFKeyPhyEffect            int16 = 236 // 物理效果
	BFKeyMagEffect            int16 = 237 // 法术效果
	BFKeyPhyAbsorb            int16 = 238 // 物理吸收
	BFKeyMagAbsorb            int16 = 239 // 法术吸收
	BFKeyAddMaxLife           int16 = 240 // 增加最大生命
	BFKeyAddMaxMana           int16 = 241 // 增加最大内力
	BFKeyAddUserLevel         int16 = 242 // 增加用户等级
	BFKeyAddRandomSkill       int16 = 243 // 增加随机技能
	BFKeyDoubleTime           int16 = 244 // 双倍时间
	BFKeyAddPetLevel          int16 = 245 // 增加宠物等级
	BFKeyPetAheadSkill        int16 = 246 // 宠物前置技能
	BFKeyPetLongevity         int16 = 247 // 宠物寿命
	BFKeyUpgradeType          int16 = 248 // 升级类型
	BFKeyAddPetExp            int16 = 249 // 增加宠物经验

	// === 房屋系统 ===
	BFKeyHouseId               int16 = 250 // 房屋ID
	BFKeyHouseHouseClass       int16 = 251 // 房屋等级
	BFKeyPlantLevel            int16 = 252 // 植物等级
	BFKeyPlantExp              int16 = 253 // 植物经验
	BFKeyToBeDeleted           int16 = 260 // 待删除
	BFKeyLocked                int16 = 261 // 锁定
	BFKeyPetUpgraded           int16 = 262 // 宠物已升级
	BFKeyLeftTimeToDelete      int16 = 263 // 剩余删除时间
	BFKeyExtraDesc             int16 = 264 // 额外描述
	BFKeyPhyRebuildLevel       int16 = 265 // 物理重建等级
	BFKeyMagRebuildLevel       int16 = 266 // 法术重建等级
	BFKeyRawName               int16 = 267 // 原始名字
	BFKeySuitPolar             int16 = 268 // 套装五行
	BFKeySuitEnabled           int16 = 269 // 套装启用
	BFKeyGift                  int16 = 270 // 礼物
	BFKeyRecognizeRecognized   int16 = 271 // 识别已识别
	BFKeyPartyStagePartyName   int16 = 272 // 阶段帮派名
	BFKeyPartyStagePassedCount int16 = 273 // 阶段通过次数
	BFKeyPartyStageCostTime    int16 = 274 // 阶段消耗时间
	BFKeyPartyStageMemberName  int16 = 275 // 阶段成员名
	BFKeyProp2Color            int16 = 276 // 属性2颜色

	// === 摔跤系统 ===
	BFKeyWrestleScore            int16 = 277 // 摔跤分数
	BFKeyWrestleScore2           int16 = 278 // 摔跤分数2
	BFKeyDefEffect               int16 = 279 // 防御效果
	BFKeyCombatGuard             int16 = 280 // 战斗守护
	BFKeyCombined                int16 = 282 // 组合
	BFKeyOpenNimbus              int16 = 283 // 启灵
	BFKeyLimitUseTimeIneffective int16 = 284 // 限制使用时间无效
	BFKeyWuhunqiaoLevel          int16 = 285 // 武魂窍等级
	BFKeyWeaponIcon              int16 = 290 // 武器图标
	BFKeySuitIcon                int16 = 291 // 套装图标
	BFKeyOrgIcon                 int16 = 292 // 原始图标
	BFKeyTaoEffect               int16 = 293 // 道行效果
	BFKeyMountType               int16 = 294 // 坐骑类型
	BFKeySuitLightEffect         int16 = 295 // 套装光效
	BFKeySuitPolarPreview        int16 = 296 // 套装五行预览
	BFKeySignature2              int16 = 300 // 签名(重复)
	BFKeyInsiderTime             int16 = 301 // 内部时间
	BFKeyUserState               int16 = 302 // 用户状态
	BFKeyAutoReply               int16 = 303 // 自动回复
	BFKeyFriendImage             int16 = 304 // 好友形象
	BFKeyGid                     int16 = 305 // 全局唯一ID
	BFKeyIidStr                  int16 = 306 // IID字符串
	BFKeyAutoFight               int16 = 307 // 自动战斗
	BFKeyFreeRename              int16 = 308 // 免费改名
	BFKeyVoucher                 int16 = 309 // 代金券
	BFKeyUseMoneyType            int16 = 310 // 使用货币类型
	BFKeyLockExp                 int16 = 311 // 锁定经验
	BFKeyShuadaoJijiRulvling     int16 = 312 // 蜀道竞技
	BFKeyFetchNice               int16 = 313 // 获取好感
	BFKeyRecharge                int16 = 314 // 充值

	// === 额外效果 ===
	BFKeyExtraLifeEffect  int16 = 315 // 额外生命效果
	BFKeyExtraManaEffect  int16 = 316 // 额外内力效果
	BFKeyExtraMagEffect   int16 = 317 // 额外法术效果
	BFKeyExtraPhyEffect   int16 = 318 // 额外物理效果
	BFKeyExtraSpeedEffect int16 = 319 // 额外速度效果

	// === 变化时间 ===
	BFKeyMorphLifeTimes     int16 = 320 // 变化生命次数
	BFKeyMorphManaTimes     int16 = 321 // 变化内力次数
	BFKeyMorphSpeedTimes    int16 = 322 // 变化速度次数
	BFKeyMorphPhyTimes      int16 = 323 // 变化物理次数
	BFKeyMorphMagTimes      int16 = 324 // 变化法术次数
	BFKeyMorphLifeStat      int16 = 325 // 变化生命状态
	BFKeyMorphManaStat      int16 = 326 // 变化内力状态
	BFKeyMorphSpeedStat     int16 = 327 // 变化速度状态
	BFKeyMorphPhyStat       int16 = 328 // 变化物理状态
	BFKeyMorphMagStat       int16 = 329 // 变化法术状态
	BFKeyFreeUnlockExpTimes int16 = 330 // 免费解锁经验次数
	BFKeyWeekAct            int16 = 333 // 周活动
	BFKeyComebackFlag       int16 = 334 // 回归标志
	BFKeyPlacedAmount       int16 = 335 // 放置数量
	BFKeyAchieve            int16 = 336 // 成就
	BFKeyAchieveName        int16 = 337 // 成就名
	BFKeyAchieveTime        int16 = 338 // 成就时间

	// === 升级系统 ===
	BFKeyUpgradeState          int16 = 340 // 升级状态
	BFKeyUpgradeType2          int16 = 341 // 升级类型(重复)
	BFKeyUpgradeLevel          int16 = 342 // 升级等级
	BFKeyUpgradeExp            int16 = 343 // 升级经验
	BFKeyUpgradeExpToNextLevel int16 = 344 // 升级到下一级经验
	BFKeyUpgradeMaxPolarExtra  int16 = 345 // 升级最大五行额外
	BFKeyUpgradeLevel2         int16 = 346 // 升级等级2
	BFKeyHasUpgraded           int16 = 347 // 是否已升级
	BFKeyLimitUseTime          int16 = 348 // 限制使用时间
	BFKeyFasionType            int16 = 349 // 时装类型
	BFKeyFoodNum               int16 = 350 // 食物数量
	BFKeyMaxFoodNum            int16 = 351 // 最大食物数量
	BFKeyHouseId2              int16 = 352 // 房屋ID(重复)
	BFKeyComfort               int16 = 353 // 舒适度
	BFKeyCoupleName            int16 = 354 // 配偶名
	BFKeyHouseType             int16 = 355 // 房屋类型
	BFKeySubType               int16 = 356 // 子类型
	BFKeyCoupleGid             int16 = 357 // 配偶GID

	// === 经验仓库 ===
	BFKeyExpWareDataUnlockTime      int16 = 358 // 解锁时间
	BFKeyExpWareDataLockTime        int16 = 359 // 锁定时间
	BFKeyExpWareDataExpWare         int16 = 360 // 经验仓库
	BFKeyExpWareDataFetchTimes      int16 = 361 // 领取次数
	BFKeyExpWareDataTodayFetchTimes int16 = 362 // 今日领取次数
	BFKeyStage                      int16 = 363 // 阶段
	BFKeyEnergy                     int16 = 364 // 能量

	// === 属性分配 ===
	BFKeyAttribAssignStr int16 = 365 // 分配力量
	BFKeyAttribAssignWiz int16 = 366 // 分配智力
	BFKeyAttribAssignCon int16 = 367 // 分配体质
	BFKeyAttribAssignDex int16 = 368 // 分配敏捷
	BFKeyPhyShape        int16 = 369 // 物理形态
	BFKeyMagShape        int16 = 370 // 法术形态
	BFKeySpeedShape      int16 = 371 // 速度形态
	BFKeyManaShape       int16 = 372 // 内力形态
	BFKeyLifeShape       int16 = 373 // 生命形态
	BFKeyUpgradeGid      int16 = 374 // 升级GID
	BFKeyEnhancedPhy     int16 = 405 // 增强物理
	BFKeyIgnoreMagDodge  int16 = 406 // 忽略法术闪避
	BFKeyEnhancedMag     int16 = 407 // 增强法术
	BFKeyPopular         int16 = 410 // 人气

	// === 交易系统 ===
	BFKeyTradingGoodsGid      int16 = 428 // 交易商品GID
	BFKeyTradingState         int16 = 429 // 交易状态
	BFKeyTradingLeftTime      int16 = 430 // 交易剩余时间
	BFKeyTradingPrice         int16 = 431 // 交易价格
	BFKeyTradingOrgPrice      int16 = 432 // 交易原价
	BFKeyTradingCgPriceTi     int16 = 433 // 交易成功价格时间
	BFKeyTradingCgPriceCt     int16 = 434 // 交易成功价格
	BFKeyCharOnlineState      int16 = 435 // 角色在线状态
	BFKeyTradingSellBuyType   int16 = 436 // 交易买卖类型
	BFKeyTradingAppointeeName int16 = 437 // 交易指定玩家名
	BFKeyTradingBuyoutPrice   int16 = 438 // 交易一口价

	// === 丹道系统 ===
	BFKeyDanDataState          int16 = 439 // 丹道状态
	BFKeyDanDataStage          int16 = 440 // 丹道阶段
	BFKeyDanDataExp            int16 = 441 // 丹道经验
	BFKeyDanDataExpToNextLevel int16 = 442 // 丹道升级经验
	BFKeyDanDataAttribPoint    int16 = 443 // 丹道属性点
	BFKeyDanDataPolarPoint     int16 = 444 // 丹道五行点
	BFKeyNotCheckBw            int16 = 445 // 不检查BW
	BFKeyHasBreakLvLimit       int16 = 446 // 已突破等级限制
	BFKeyDanDataTodayExp       int16 = 447 // 丹道今日经验
	BFKeySoulState             int16 = 448 // 魂魄状态
	BFKeyHornName              int16 = 449 // 号角名
	BFKeyJewelryEssence        int16 = 450 // 饰品精华
	BFKeyTransformNum          int16 = 451 // 变形数量
	BFKeyTransformCoolTi       int16 = 452 // 变形冷却时间
	BFKeyMarriageStartTime     int16 = 453 // 结婚开始时间
	BFKeyBookId                int16 = 454 // 书籍ID
	BFKeyFasionCustomDisable   int16 = 455 // 时装自定义禁用
	BFKeyFasionEffectDisable   int16 = 456 // 时装效果禁用
	BFKeyMarriageBookId        int16 = 457 // 结婚书籍ID
	BFKeyStrengthenJewelryNum  int16 = 458 // 强化饰品数量
	BFKeyStrengthenLevel       int16 = 459 // 强化等级
	BFKeyStrengthenExp         int16 = 460 // 强化经验
	BFKeyStrengthenDegree      int16 = 461 // 强化等级度
	BFKeyMonTao                int16 = 462 // 本日道行
	BFKeyMonTaoEx              int16 = 463 // 本日道行额外
	BFKeyLastMonTao            int16 = 464 // 昨日道行
	BFKeyLastMonTaoEx          int16 = 465 // 昨日道行额外
	BFKeyMonMartial            int16 = 466 // 本日武学
	BFKeyLastMonMartial        int16 = 467 // 昨日武学
	BFKeyMonTaoRank            int16 = 468 // 本日道行排名

	// === 神魂系统 ===
	BFKeyShenhunDataState          int16 = 469 // 神魂状态
	BFKeyShenhunDataLayer          int16 = 470 // 神魂层
	BFKeyShenhunDataExp            int16 = 471 // 神魂经验
	BFKeyShenhunDataExpToNextLevel int16 = 472 // 神魂升级经验
	BFKeyYqzcNum                   int16 = 473 // 运气值
	BFKeyCSkillDodge               int16 = 474 // C技能闪避
	BFKeyIgnoreCSkillDodge         int16 = 475 // 忽略C技能闪避
	BFKeyStealBuffRate             int16 = 476 // 偷取buff率
	BFKeyIgnoreStealBuffRate       int16 = 477 // 忽略偷取buff率
	BFKeyReliveRate                int16 = 478 // 重生率
	BFKeyIgnoreReliveRate          int16 = 479 // 忽略重生率
	BFKeySuckBloodRate             int16 = 480 // 吸血率
	BFKeyIgnoreSuckBloodRate       int16 = 481 // 忽略吸血率
	BFKeyCrippleRate               int16 = 482 // 残废率
	BFKeyHunqiaoPage               int16 = 487 // 魂器页
	BFKeyFightCap                  int16 = 493 // 战斗力
	BFKeyFightCapWithoutIntimacy   int16 = 494 // 战斗力(无亲密度)
	BFKeyWuhunDataLevel            int16 = 495 // 武魂等级
	BFKeyWuhunDataExp              int16 = 496 // 武魂经验
	BFKeyWuhunDataExpToNextLevel   int16 = 497 // 武魂升级经验
	BFKeyWhzcNum                   int16 = 498 // 武魂值
	BFKeyWuhunqiaoPage             int16 = 499 // 武魂窍页
	BFKeyFightCapTotal             int16 = 500 // 总战斗力
	BFKeyCwCompeteFscore           int16 = 501 // 竞争分数
	BFKeyQmMoney                   int16 = 502 // 求魔币
	BFKeyGhostdomChallengeLevel    int16 = 503 // 鬼域挑战等级

	// === 重建属性 ===
	BFKeyPhyRebuildRate    int16 = 719 // 物理重建率
	BFKeyMagRebuildRate    int16 = 720 // 法术重建率
	BFKeyLifeAddTemp       int16 = 721 // 生命临时加成
	BFKeyManaAddTemp       int16 = 722 // 内力临时加成
	BFKeyPhyPowerAddTemp   int16 = 723 // 物理攻击临时加成
	BFKeyMagPowerAddTemp   int16 = 724 // 法术攻击临时加成
	BFKeySpeedAddTemp      int16 = 725 // 速度临时加成
	BFKeyDefAddTemp        int16 = 726 // 防御临时加成
	BFKeyPhyRebuildAdd     int16 = 800 // 物理重建加成
	BFKeyMagRebuildAdd     int16 = 801 // 法术重建加成
	BFKeyPetLifeShapeTemp  int16 = 802 // 宠物生命形态临时
	BFKeyPetManaShapeTemp  int16 = 803 // 宠物体力形态临时
	BFKeyPetSpeedShapeTemp int16 = 804 // 宠物速度形态临时
	BFKeyPetPhyShapeTemp   int16 = 805 // 宠物物理形态临时
	BFKeyPetMagShapeTemp   int16 = 806 // 宠物法术形态临时
	BFKeyEvolveDegree      int16 = 807 // 进化等级
	BFKeyPower             int16 = 808 // 力量
	BFKeySkillLowCost      int16 = 809 // 技能低消耗
	BFKeyRebuildDegree     int16 = 810 // 重建等级度
	BFKeyTime              int16 = 811 // 时间
	BFKeyEffectFoot        int16 = 812 // 脚步效果
	BFKeyFightScore        int16 = 813 // 战斗分数
	BFKeyRecvGid           int16 = 814 // 接收GID
	BFKeySuitLevel         int16 = 815 // 套装等级
	BFKeySuitDegree        int16 = 816 // 套装等级度

	// === 状态效果 ===
	BFKeyStatusPoison           int16 = 817 // 毒状态
	BFKeyStatusSleep            int16 = 818 // 睡眠状态
	BFKeyStatusForgotten        int16 = 819 // 遗忘状态
	BFKeyStatusFrozen           int16 = 820 // 冰冻状态
	BFKeyStatusConfusion        int16 = 821 // 混乱状态
	BFKeyStatusJointAttack      int16 = 822 // 合击状态
	BFKeyStatusRevive           int16 = 823 // 复活状态
	BFKeyStatusStunt            int16 = 824 // 特技状态
	BFKeyStatusDoubleHit        int16 = 825 // 双击状态
	BFKeyStatusDamageSel        int16 = 826 // 伤害选择状态
	BFKeyStatusCounterAttack    int16 = 827 // 反击状态
	BFKeyStatusProtected        int16 = 828 // 保护状态
	BFKeyStatusSpeed            int16 = 829 // 速度状态
	BFKeyStatusPhyPower         int16 = 830 // 物理攻击状态
	BFKeyStatusDefense          int16 = 831 // 防御状态
	BFKeyStatusMaxLife          int16 = 832 // 最大生命状态
	BFKeyStatusDodge            int16 = 833 // 闪避状态
	BFKeyStatusDef              int16 = 834 // 防御状态2
	BFKeyStatusRecoverLife      int16 = 835 // 生命恢复状态
	BFKeyStatusMetal            int16 = 836 // 金状态
	BFKeyStatusWood             int16 = 837 // 木状态
	BFKeyStatusWater            int16 = 838 // 水状态
	BFKeyStatusFire             int16 = 839 // 火状态
	BFKeyStatusEarth            int16 = 840 // 土状态
	BFKeyStatusLeechPhyDamage   int16 = 841 // 吸取物理伤害状态
	BFKeyStatusLeechMagDamage   int16 = 842 // 吸取法术伤害状态
	BFKeyStatusPassiveAttack    int16 = 843 // 被动攻击状态
	BFKeyStatusDeadlyKiss       int16 = 844 // 致命之吻状态
	BFKeyStatusLoyalty          int16 = 845 // 忠诚度状态
	BFKeyStatusImmunePhyDamage  int16 = 846 // 免疫物理伤害状态
	BFKeyStatusImmuneMagDamage  int16 = 847 // 免疫法术伤害状态
	BFKeyStatusPolarChanged     int16 = 848 // 五行变化状态
	BFKeyStatusFanzhuanQiankun  int16 = 849 // 翻转乾坤状态
	BFKeyStatusManaShield       int16 = 850 // 内力护盾状态
	BFKeyStatusPassiveMagAttack int16 = 851 // 被动法术攻击状态
	BFKeyStatusAddLifeByMana    int16 = 852 // 内力加血状态
	BFKeyCombatGuardIndex       int16 = 853 // 战斗守护索引
	BFKeySalary                 int16 = 854 // 俸禄
	BFKeyParty                  int16 = 855 // 帮派
	BFKeyPopulation             int16 = 856 // 人口
	BFKeyPartyWarWin            int16 = 857 // 帮派战胜利
	BFKeyLevelUpTime            int16 = 858 // 升级时间
	BFKeyHigestXiangy           int16 = 859 // 最高逍遥
	BFKeyHigestChub             int16 = 860 // 最高出没
	BFKeyHigestFum              int16 = 861 // 最高父母
	BFKeyHigestTongtt           int16 = 862 // 最高通天
	BFKeyHigestYasby            int16 = 863 // 最高亚洲
	BFKeyTonttLayer             int16 = 864 // 通天层
	BFKeyMoney                  int16 = 865 // 金钱
	BFKeyTaoRank                int16 = 866 // 道行排名
	BFKeyLotteryTimes           int16 = 868 // 抽奖次数
	BFKeyShadowSelf             int16 = 869 // 影子分身
	BFKeyEnchant                int16 = 870 // 附魔
	BFKeyEnchantNimbus          int16 = 871 // 附魔灵气
	BFKeyMaxEnchantNimbus       int16 = 872 // 最大附魔灵气
	BFKeyCardType               int16 = 873 // 卡片类型
	BFKeyEffectTime             int16 = 874 // 效果时间
	BFKeyShenmuPoints           int16 = 875 // 神木点
	BFKeyEnableShenmuPoints     int16 = 876 // 启用神木点
	BFKeyGiftKey                int16 = 877 // 礼物键
	BFKeyExpiredTime            int16 = 878 // 过期时间
	BFKeyEvolve                 int16 = 879 // 进化
	BFKeyInsiderLevel           int16 = 880 // 内部等级
	BFKeyEvolveLevel            int16 = 881 // 进化等级
	BFKeyMailingItemTimes       int16 = 882 // 邮寄物品次数
	BFKeyMountAttribEndTime     int16 = 883 // 坐骑属性结束时间
	BFKeyMountAttribMoveSpeed   int16 = 884 // 坐骑属性移动速度
	BFKeyCapacityLevel          int16 = 885 // 容量等级
	BFKeyHideMount              int16 = 886 // 隐藏坐骑
	BFKeyDeadline               int16 = 887 // 截止时间
	BFKeyMergeRate              int16 = 888 // 合并率
	BFKeyEquipPerfectPercent    int16 = 889 // 装备完美百分比
	BFKeyDunwuTimes             int16 = 890 // 遁物次数
	BFKeyDunwuRate              int16 = 891 // 遁物率
	BFKeyPetAnger               int16 = 892 // 宠物怒气

	// === 状态效果(续) ===
	BFKeyStatusHuanbingZhiji        int16 = 893  // 寒冰直击状态
	BFKeyStatusAitongYujue          int16 = 894  // 哀痛欲绝状态
	BFKeyStatusShushouJiuqin        int16 = 895  // 束手就擒状态
	BFKeyStatusWenfengSangdan       int16 = 896  // 文风丧胆状态
	BFKeyStatusYangjingXurui        int16 = 897  // 养精蓄锐状态
	BFKeyStatusXuwu                 int16 = 898  // 虚无状态
	BFKeyItemPolar                  int16 = 899  // 物品五行
	BFKeyShuadaoZiqihongmeng        int16 = 900  // 蜀道紫气鸿蒙
	BFKeyExtraSkill                 int16 = 901  // 额外技能
	BFKeyExtraSkillLevel            int16 = 902  // 额外技能等级
	BFKeyDiandqkFrozenRound         int16 = 903  // 电池冻结回合
	BFKeyStatusChaofeng             int16 = 904  // 嘲讽状态
	BFKeyStatusDiandaoQiankun       int16 = 905  // 颠倒乾坤状态
	BFKeyStatusJingangquan          int16 = 906  // 金刚圈状态
	BFKeyStatusQinmiWujian          int16 = 907  // 亲密无间状态
	BFKeyStatusTianyan              int16 = 908  // 天眼状态
	BFKeyStatusWujiBifan            int16 = 909  // 无极必反状态
	BFKeyStatusShowOpponentLife     int16 = 910  // 显示对手生命状态
	BFKeyStatusAddLifeByManaAdd     int16 = 911  // 内力加血加成状态
	BFKeyStatusRecoverLifeAdd       int16 = 912  // 生命恢复加成状态
	BFKeyStatusDefAdd               int16 = 913  // 防御加成状态
	BFKeyStatusPhyPowerAdd          int16 = 914  // 物理攻击加成状态
	BFKeyStatusMagPowerAdd          int16 = 915  // 法术攻击加成状态
	BFKeyStatusSpeedAdd             int16 = 916  // 速度加成状态
	BFKeyBrotherAppellation         int16 = 917  // 兄弟称号
	BFKeyShuadaoRuyiPoint           int16 = 918  // 蜀道如意点
	BFKeyChushiEx                   int16 = 919  // 初始额外
	BFKeyPhyPowerWithoutIntimacy    int16 = 920  // 物理攻击(无亲密度)
	BFKeyMagPowerWithoutIntimacy    int16 = 921  // 法术攻击(无亲密度)
	BFKeyDefWithoutIntimacy         int16 = 922  // 防御(无亲密度)
	BFKeyOriginIntimacy             int16 = 923  // 原始亲密度
	BFKeyDouchongRank               int16 = 924  // 斗宠排名
	BFKeyUpgradeImmortal            int16 = 925  // 升级仙人
	BFKeyUpgradeMagic               int16 = 926  // 升级魔法
	BFKeyUpgradeTotal               int16 = 927  // 升级总计
	BFKeyArtifactUpgradedEnabled    int16 = 928  // 神器升级启用
	BFKeyStatusQishaYin             int16 = 929  // 七杀阴状态
	BFKeyStatusQishaYang            int16 = 930  // 七杀阳状态
	BFKeyEclosion                   int16 = 931  // 羽化
	BFKeyEclosionNimbus             int16 = 932  // 羽化灵气
	BFKeyMaxEclosionNimbus          int16 = 933  // 最大羽化灵气
	BFKeyStatusAllResistExceptAdd   int16 = 934  // 全抗除外加成状态
	BFKeyEclosionStage              int16 = 935  // 羽化阶段
	BFKeyStatusYanchuanShenjiao     int16 = 936  // 言传身教状态
	BFKeyBossAnger                  int16 = 937  // BOSS怒气
	BFKeyHigestScore                int16 = 938  // 最高分
	BFKeyPartyId                    int16 = 939  // 帮派ID
	BFKeyStatusDaofaWubian          int16 = 940  // 道法无边状态
	BFKeyObtainTime                 int16 = 941  // 获取时间
	BFKeyEffectWaist                int16 = 942  // 腰部效果
	BFKeyEffectHead                 int16 = 943  // 头部效果
	BFKeyStatusWeiya                int16 = 944  // 威压状态
	BFKeyStatusWeiYaCount           int16 = 945  // 威压次数状态
	BFKeyGraphicInstructionMark     int16 = 946  // 图形指令标记
	BFKeyStatusDilieboFlag          int16 = 947  // 敌对波标志状态
	BFKeyPartIndex                  int16 = 948  // 部分索引
	BFKeyPartColorIndex             int16 = 949  // 部分颜色索引
	BFKeyNpcChat                    int16 = 950  // NPC聊天
	BFKeyFollowPetType              int16 = 951  // 跟随宠物类型
	BFKeyDyeIcon                    int16 = 952  // 染色图标
	BFKeyFasionId                   int16 = 953  // 时装ID
	BFKeyFasionVisible              int16 = 954  // 时装可见
	BFKeyActCamp                    int16 = 955  // 活动阵营
	BFKeyLingchenPoint              int16 = 956  // 凌晨点
	BFKeyNpcState                   int16 = 957  // NPC状态
	BFKeyChannelSource              int16 = 958  // 频道来源
	BFKeyFlagChild                  int16 = 959  // 标记子项
	BFKeyGhostGas                   int16 = 960  // 鬼气
	BFKeyPeiyuanDataLevel           int16 = 961  // 培元等级
	BFKeyPeiyuanDataStage           int16 = 962  // 培元阶段
	BFKeyNingshenState              int16 = 963  // 凝神状态
	BFKeyPeiyuanDataState           int16 = 964  // 培元状态
	BFKeyGongshengDataTargetPetIid  int16 = 975  // 共生目标宠物IID
	BFKeySkillLevel                 int16 = 965  // 技能等级
	BFKeyUpgradeDegree              int16 = 966  // 升级等级度
	BFKeyStatusFengmangAdd          int16 = 967  // 锋芒加成状态
	BFKeyStatusGuibuAdd             int16 = 968  // 归步加成状态
	BFKeyStatusBomuAdd              int16 = 969  // 博弈加成状态
	BFKeyStatusFengmang             int16 = 970  // 锋芒状态
	BFKeyStatusHundeng              int16 = 971  // 混沌状态
	BFKeyStatusGuibu                int16 = 972  // 归步状态
	BFKeyStatusBomu                 int16 = 973  // 博弈状态
	BFKeyStatusFuhu                 int16 = 974  // 伏击状态
	BFKeyYoutjEffectFlag            int16 = 977  // 有机效果标志
	BFKeyYoutjSunPoint              int16 = 978  // 有机太阳点
	BFKeyYoutjMoonPoint             int16 = 979  // 有机月亮点
	BFKeyHigestXiangy2              int16 = 989  // 最高逍遥2
	BFKeyHigestFum2                 int16 = 990  // 最高父母2
	BFKeyHigestFeixdx2              int16 = 991  // 最高飞仙2
	BFKeyStatusTiandaoHuti          int16 = 992  // 天道糊涂状态
	BFKeyUlevel                     int16 = 993  // U等级
	BFKeyMaxStoreExp                int16 = 994  // 最大存储经验
	BFKeyDigongMaxPassLevel         int16 = 995  // 地宫最大通关等级
	BFKeyStatusLihunZhihuo          int16 = 996  // 离魂致火状态
	BFKeyStatusXuanwuYin            int16 = 997  // 玄武阴状态
	BFKeyItemMaxLiveness            int16 = 998  // 物品最大活力
	BFKeyItemLiveness               int16 = 999  // 物品活力
	BFKeyItemMaxCostGoldCoin        int16 = 1000 // 物品最大花费金币
	BFKeyItemCostGoldCoin           int16 = 1001 // 物品花费金币
	BFKeyStatusYishouWeigong        int16 = 1002 // 易守难攻状态
	BFKeyStatusYigongWeishou        int16 = 1003 // 易攻难守状态
	BFKeyZhenlingType               int16 = 1005 // 真灵类型
	BFKeyZhenlingLevel              int16 = 1006 // 真灵等级
	BFKeyStatusXielingzhiyan        int16 = 1007 // 血灵之眼状态
	BFKeyStatusTongshengXueshi      int16 = 1008 // 同声血誓状态
	BFKeyStatusGongsiZuzhou         int16 = 1009 // 公司诅咒状态
	BFKeyStatusShaluZhixin          int16 = 1010 // 杀戮之心状态
	BFKeyStatusMeihuo               int16 = 1011 // 魅惑状态
	BFKeyStatusJiuchan              int16 = 1012 // 纠缠状态
	BFKeyStatusShashu               int16 = 1013 // 杀戮状态
	BFKeyStatusFatianXiangdi        int16 = 1014 // 翻天象地状态
	BFKeyStatusZhanxian             int16 = 1015 // 斩仙状态
	BFKeyStatusZhanxianAdd          int16 = 1016 // 斩仙加成状态
	BFKeyOverseaAccount             int16 = 1017 // 海外账号
	BFKeyJcxBalance                 int16 = 1018 // JCX余额
	BFKeyMaxJcxBalance              int16 = 1019 // 最大JCX余额
	BFKeyCanUsePetAnger             int16 = 1032 // 可使用宠物怒气
	BFKeyNotAllowMove               int16 = 1054 // 不允许移动
	BFKeyHudunNum                   int16 = 1055 // 护盾数量
	BFKeyLiehunMoyinCount           int16 = 1057 // 裂魂魔音次数
	BFKeyPhantomIcon                int16 = 1058 // 幻影图标
	BFKeyStatusShihun               int16 = 1059 // 失魂状态
	BFKeyCombatTitle                int16 = 1060 // 战斗称号
	BFKeyShowLifeBarBoss            int16 = 1061 // 显示BOSS生命条
	BFKeyWuxingMozhen               int16 = 1062 // 五行魔真
	BFKeyStatusChunyangZhenhuo      int16 = 1063 // 纯阳真火状态
	BFKeyStatusZhuoshao             int16 = 1064 // 灼烧状态
	BFKeyStatusZhuoshaoAdd          int16 = 1065 // 灼烧加成状态
	BFKeyOpacity                    int16 = 1066 // 不透明度
	BFKeyCwYayun                    int16 = 1067 // CW押韵
	BFKeyStatusShuanghanQinxi       int16 = 1068 // 双寒侵袭状态
	BFKeyStatusXinwuPangwu          int16 = 1069 // 心无旁骛状态
	BFKeyStatusGubenPeiyuan         int16 = 1070 // 固本培元状态
	BFKeyStatusGuwuShiqi            int16 = 1071 // 鼓舞士气状态
	BFKeyStatusShuanghanQinxiAdd    int16 = 1072 // 双寒侵袭加成状态
	BFKeyStatusXinwuPangwuNum       int16 = 1073 // 心无旁骛次数状态
	BFKeyStatusGubenPeiyuanAdd      int16 = 1074 // 固本培元加成状态
	BFKeyStatusGuwuShiqiAdd         int16 = 1075 // 鼓舞士气加成状态
	BFKeyStatusWandaoChengkong      int16 = 1076 // 弯道成空状态
	BFKeyWufaWushuang               int16 = 1077 // 无法无双
	BFKeyJianruPanshi               int16 = 1078 // 坚韧磐石
	BFKeyXinghuoFeichi              int16 = 1079 // 星火飞驰
	BFKeyNingqiHuadun               int16 = 1080 // 凝气化盾
	BFKeyCanFly                     int16 = 1081 // 可飞行
	BFKeyStrengthenDegree80         int16 = 1082 // 强化等级度80
	BFKeyStatusRecoverMana          int16 = 1083 // 内力恢复状态
	BFKeyStatusRecoverManaAdd       int16 = 1084 // 内力恢复加成状态
	BFKeyCamp                       int16 = 1085 // 阵营
	BFKeyZhengdaoTaskIndex          int16 = 1086 // 正道路线索引
	BFKeyStatusZd05                 int16 = 1087 // 状态ZD05
	BFKeyStatusYinyangjing          int16 = 1088 // 阴阳镜状态
	BFKeyStatusYinyangjingCount     int16 = 1089 // 阴阳镜次数状态
	BFKeyStatusXiejiaJinhu          int16 = 1090 // 卸甲金弧状态
	BFKeyStatusXiejiaJinhuSkill     int16 = 1091 // 卸甲金弧技能状态
	BFKeyYubxsLevel                 int16 = 1092 // 预备等级
	BFKeyShidaoDahui                int16 = 2000 // 十大大会
	BFKeyHigestFeixdx               int16 = 2001 // 最高飞仙
	BFKeyOpenState                  int16 = 2002 // 开启状态
	BFKeyStatusDiandaoCuoluanAdd    int16 = 2007 // 颠倒错乱加成状态
	BFKeyStatusShuanghanZhihuAdd    int16 = 2008 // 双寒至呼加成状态
	BFKeyStatus                     int16 = 2012 // 状态
	BFKeyCscwQiaozhuang             int16 = 2038 // CSW乔装
	BFKeyMarriageMarryId            int16 = 3000 // 结婚ID
	BFKeyAlias                      int16 = 3001 // 别名
	BFKeyShuadaoChongfengSan        int16 = 3002 // 蜀道冲锋散
	BFKeyGroupName                  int16 = 3003 // 群组名
	BFKeyGroupId                    int16 = 3004 // 群组ID
	BFKeyLeaderGid                  int16 = 3005 // 领袖GID
	BFKeyMemberGid                  int16 = 3006 // 成员GID
	BFKeySetting                    int16 = 3007 // 设置
	BFKeyAnnouncement               int16 = 3008 // 公告
	BFKeySettingRefuseStrangerLevel int16 = 3009 // 设置拒绝陌生人等级
	BFKeySettingAutoReplyMsg        int16 = 3010 // 设置自动回复消息
	BFKeySettingRefuseBeAddLevel    int16 = 3011 // 设置拒绝被加等级
	BFKeyServerName                 int16 = 3012 // 服务器名
	BFKeyBullyKillNum               int16 = 3013 // 恶霸击杀数
	BFKeyPoliceKillNum              int16 = 3014 // 警察击杀数
	BFKeyShowSandglass              int16 = 3015 // 显示沙漏

	// === GM属性 ===
	BFKeyGmAttribsMaxLife      int16 = 3016 // GM最大生命
	BFKeyGmAttribsMaxMana      int16 = 3017 // GM最大内力
	BFKeyGmAttribsPhyPower     int16 = 3018 // GM物理攻击
	BFKeyGmAttribsMagPower     int16 = 3019 // GM法术攻击
	BFKeyGmAttribsDef          int16 = 3020 // GM防御
	BFKeyGmAttribsSpeed        int16 = 3021 // GM速度
	BFKeyMarriageCoupleGid     int16 = 3022 // 结婚配偶GID
	BFKeyChatHead              int16 = 3023 // 聊天头
	BFKeyChatFloor             int16 = 3024 // 聊天楼层
	BFKeyDistName              int16 = 3025 // 距离名
	BFKeyChannelFilterType     int16 = 3026 // 频道过滤类型
	BFKeyRoomName              int16 = 3027 // 房间名
	BFKeyTeamAssessScore       int16 = 3028 // 团队评估分数
	BFKeyChannelActType        int16 = 3029 // 频道活动类型
	BFKeyBanRule               int16 = 3030 // 禁言规则
	BFKeyAiteParty             int16 = 3031 // 艾特帮派
	BFKeyMessageId             int16 = 3032 // 消息ID
	BFKeySpecialMonsterChannel int16 = 3033 // 特殊怪物频道
)

// buildFieldKeyMap 常量名到数字的映射表
// 运行时通过 tag 中的常量名查找对应的数字 key
var buildFieldKeyMap = map[string]int16{
	// === 基础属性 ===
	"Name":              BFKeyName,
	"Str":               BFKeyStr,
	"PhyPower":          BFKeyPhyPower,
	"Accurate":          BFKeyAccurate,
	"Con":               BFKeyCon,
	"Life":              BFKeyLife,
	"MaxLife":           BFKeyMaxLife,
	"Def":               BFKeyDef,
	"Wiz":               BFKeyWiz,
	"MagPower":          BFKeyMagPower,
	"Mana":              BFKeyMana,
	"MaxMana":           BFKeyMaxMana,
	"Dex":               BFKeyDex,
	"Speed":             BFKeySpeed,
	"Parry":             BFKeyParry,
	"AttribPoint":       BFKeyAttribPoint,
	"PolarPoint":        BFKeyPolarPoint,
	"Stamina":           BFKeyStamina,
	"MaxStamina":        BFKeyMaxStamina,
	"Tao":               BFKeyTao,
	"Friend":            BFKeyFriend,
	"TotalPk":           BFKeyTotalPk,
	"Degree":            BFKeyDegree,
	"Exp":               BFKeyExp,
	"Pot":               BFKeyPot,
	"Cash":              BFKeyCash,
	"Balance":           BFKeyBalance,
	"Gender":            BFKeyGender,
	"Master":            BFKeyMaster,
	"Level":             BFKeyLevel,
	"Skill":             BFKeySkill,
	"PartyName":         BFKeyPartyName,
	"PartyContrib":      BFKeyPartyContrib,
	"Nick":              BFKeyNick,
	"Title":             BFKeyTitle,
	"Nice":              BFKeyNice,
	"Reputation":        BFKeyReputation,
	"Couple":            BFKeyCouple,
	"Icon":              BFKeyIcon,
	"Type":              BFKeyType,
	"Durability":        BFKeyDurability,
	"MaxDurability":     BFKeyMaxDurability,
	"Polar":             BFKeyPolar,
	"Metal":             BFKeyMetal,
	"Wood":              BFKeyWood,
	"Water":             BFKeyWater,
	"Fire":              BFKeyFire,
	"Earth":             BFKeyEarth,
	"ResistMetal":       BFKeyResistMetal,
	"ResistWood":        BFKeyResistWood,
	"ResistWater":       BFKeyResistWater,
	"ResistFire":        BFKeyResistFire,
	"ResistEarth":       BFKeyResistEarth,
	"ExpToNextLevel":    BFKeyExpToNextLevel,
	"ResistPoison":      BFKeyResistPoison,
	"ResistFrozen":      BFKeyResistFrozen,
	"ResistSleep":       BFKeyResistSleep,
	"ResistForgotten":   BFKeyResistForgotten,
	"ResistConfusion":   BFKeyResistConfusion,
	"Longevity":         BFKeyLongevity,
	"Martial":           BFKeyMartial,
	"Intimacy":          BFKeyIntimacy,
	"Shape":             BFKeyShape,
	"ResistPoint":       BFKeyResistPoint,
	"Loyalty":           BFKeyLoyalty,
	"DoubleHit":         BFKeyDoubleHit,
	"Stunt":             BFKeyStunt,
	"CounterAttack":     BFKeyCounterAttack,
	"LifeRecover":       BFKeyLifeRecover,
	"ManaRecover":       BFKeyManaRecover,
	"LifeRecoverRate":   BFKeyLifeRecoverRate,
	"ManaRecoverRate":   BFKeyManaRecoverRate,
	"ItemType":          BFKeyItemType,
	"TotalScore":        BFKeyTotalScore,
	"CounterAttackRate": BFKeyCounterAttackRate,
	"DoubleHitRate":     BFKeyDoubleHitRate,
	"StuntRate":         BFKeyStuntRate,
	"DamageSel":         BFKeyDamageSel,
	"Family":            BFKeyFamily,
	"ReqStr":            BFKeyReqStr,
	"TotalDied":         BFKeyTotalDied,
	"ItemUnique":        BFKeyItemUnique,
	"DamageSelRate":     BFKeyDamageSelRate,
	"Portrait":          BFKeyPortrait,
	"PassiveMode":       BFKeyPassiveMode,
	"StaticMode":        BFKeyStaticMode,
	"Source":            BFKeySource,
	"Signature":         BFKeySignature,

	// === 性格/角色详情 ===
	"CharacterHarmony":      BFKeyCharacterHarmony,
	"CharacterKindness":     BFKeyCharacterKindness,
	"CharacterCarefulness":  BFKeyCharacterCarefulness,
	"CharacterCourage":      BFKeyCharacterCourage,
	"CharacterDesc":         BFKeyCharacterDesc,
	"PartyDesc":             BFKeyPartyDesc,
	"PartyJob":              BFKeyPartyJob,
	"FamilyTitle":           BFKeyFamilyTitle,
	"PartyTitle":            BFKeyPartyTitle,
	"ApprenticeTitle":       BFKeyApprenticeTitle,
	"ReqCon":                BFKeyReqCon,
	"ReqWiz":                BFKeyReqWiz,
	"ReqDex":                BFKeyReqDex,
	"PetLifeShape":          BFKeyPetLifeShape,
	"PetManaShape":          BFKeyPetManaShape,
	"PetSpeedShape":         BFKeyPetSpeedShape,
	"PetPhyShape":           BFKeyPetPhyShape,
	"PetMagShape":           BFKeyPetMagShape,
	"Rank":                  BFKeyRank,
	"Penetrate":             BFKeyPenetrate,
	"SpecialIcon":           BFKeySpecialIcon,
	"PenetrateRate":         BFKeyPenetrateRate,
	"Nimbus":                BFKeyNimbus,
	"SilverCoin":            BFKeySilverCoin,
	"GoldCoin":              BFKeyGoldCoin,
	"ExtraLife":             BFKeyExtraLife,
	"ExtraMana":             BFKeyExtraMana,
	"HaveCoinPwd":           BFKeyHaveCoinPwd,
	"MaxCash":               BFKeyMaxCash,
	"MaxBalance":            BFKeyMaxBalance,
	"IgnoreResistMetal":     BFKeyIgnoreResistMetal,
	"IgnoreResistWood":      BFKeyIgnoreResistWood,
	"IgnoreResistWater":     BFKeyIgnoreResistWater,
	"IgnoreResistFire":      BFKeyIgnoreResistFire,
	"IgnoreResistEarth":     BFKeyIgnoreResistEarth,
	"IgnoreResistForgotten": BFKeyIgnoreResistForgotten,
	"IgnoreResistPoison":    BFKeyIgnoreResistPoison,
	"IgnoreResistFrozen":    BFKeyIgnoreResistFrozen,
	"IgnoreResistSleep":     BFKeyIgnoreResistSleep,
	"IgnoreResistConfusion": BFKeyIgnoreResistConfusion,
	"SuperExcluseMetal":     BFKeySuperExcluseMetal,
	"SuperExcluseWood":      BFKeySuperExcluseWood,
	"SuperExcluseWater":     BFKeySuperExcluseWater,
	"SuperExcluseFire":      BFKeySuperExcluseFire,
	"SuperExcluseEarth":     BFKeySuperExcluseEarth,
	"BSkillLowCost":         BFKeyBSkillLowCost,
	"CSkillLowCost":         BFKeyCSkillLowCost,
	"DSkillLowCost":         BFKeyDSkillLowCost,
	"SuperPoison":           BFKeySuperPoison,
	"SuperSleep":            BFKeySuperSleep,
	"SuperForgotten":        BFKeySuperForgotten,
	"SuperConfusion":        BFKeySuperConfusion,
	"SuperFrozen":           BFKeySuperFrozen,
	"EnhancedMetal":         BFKeyEnhancedMetal,
	"EnhancedWood":          BFKeyEnhancedWood,
	"EnhancedWater":         BFKeyEnhancedWater,
	"EnhancedFire":          BFKeyEnhancedFire,
	"EnhancedEarth":         BFKeyEnhancedEarth,
	"MagDodge":              BFKeyMagDodge,

	// === 反击率（特定技能） ===
	"JinguangZhaxianCounterAttRate": BFKeyJinguangZhaxianCounterAttRate,
	"ZhaiyeFeihuaCounterAttRate":    BFKeyZhaiyeFeihuaCounterAttRate,
	"DishuiChuanshiCounterAttRate":  BFKeyDishuiChuanshiCounterAttRate,
	"JuhuoFentianCounterAttRate":    BFKeyJuhuoFentianCounterAttRate,
	"LuotuFeiyanCounterAttRate":     BFKeyLuotuFeiyanCounterAttRate,

	// === 五行法术攻击 ===
	"MetalMagPower": BFKeyMetalMagPower,
	"WoodMagPower":  BFKeyWoodMagPower,
	"WaterMagPower": BFKeyWaterMagPower,
	"FireMagPower":  BFKeyFireMagPower,
	"EarthMagPower": BFKeyEarthMagPower,

	// === 卡片系统 ===
	"CardLevel":             BFKeyCardLevel,
	"MaxCardAmount":         BFKeyMaxCardAmount,
	"CardAmount":            BFKeyCardAmount,
	"IgnoreAllResistPolar":  BFKeyIgnoreAllResistPolar,
	"IgnoreAllResistExcept": BFKeyIgnoreAllResistExcept,
	"ReleaseForgotten":      BFKeyReleaseForgotten,
	"ReleasePoison":         BFKeyReleasePoison,
	"ReleaseFrozen":         BFKeyReleaseFrozen,
	"ReleaseSleep":          BFKeyReleaseSleep,
	"ReleaseConfusion":      BFKeyReleaseConfusion,
	"TaoEx":                 BFKeyTaoEx,
	"OwnerName":             BFKeyOwnerName,
	"BackupLoyalty":         BFKeyBackupLoyalty,
	"UseSkillD":             BFKeyUseSkillD,
	"MaxDegree":             BFKeyMaxDegree,
	"CostNum":               BFKeyCostNum,
	"ConMax":                BFKeyConMax,
	"StrMax":                BFKeyStrMax,
	"WizMax":                BFKeyWizMax,
	"DexMax":                BFKeyDexMax,
	"ConAdd":                BFKeyConAdd,
	"StrAdd":                BFKeyStrAdd,
	"WizAdd":                BFKeyWizAdd,
	"DexAdd":                BFKeyDexAdd,
	"PracticeTimes":         BFKeyPracticeTimes,
	"DoublePoints":          BFKeyDoublePoints,
	"EnableDoublePoints":    BFKeyEnableDoublePoints,
	"CanBuyDpTimes":         BFKeyCanBuyDpTimes,
	"Online":                BFKeyOnline,
	"ArenaRank":             BFKeyArenaRank,
	"PartyName2":            BFKeyPartyName2,
	"Unidentified":          BFKeyUnidentified,
	"Degree32":              BFKeyDegree32,
	"StoreExp":              BFKeyStoreExp,
	"EquipIdentify":         BFKeyEquipIdentify,
	"Desc":                  BFKeyDesc,
	"EquipType":             BFKeyEquipType,
	"Amount":                BFKeyAmount,
	"OwnerId":               BFKeyOwnerId,
	"ReqLevel":              BFKeyReqLevel,
	"Attrib":                BFKeyAttrib,
	"Value":                 BFKeyValue,
	"RebuildLevel":          BFKeyRebuildLevel,
	"Color":                 BFKeyColor,
	"Quality":               BFKeyQuality,
	"UseTimes":              BFKeyUseTimes,
	"MaxUseTimes":           BFKeyMaxUseTimes,
	"CarpetRadius":          BFKeyCarpetRadius,
	"EquipPage":             BFKeyEquipPage,
	"MaxReqLevel":           BFKeyMaxReqLevel,
	"CreateTime":            BFKeyCreateTime,

	// === CT数据 ===
	"CtDataScore":     BFKeyCtDataScore,
	"CtDataTopRank":   BFKeyCtDataTopRank,
	"RealDesc":        BFKeyRealDesc,
	"AllAttrib":       BFKeyAllAttrib,
	"AllPolar":        BFKeyAllPolar,
	"AllResistPolar":  BFKeyAllResistPolar,
	"AllResistExcept": BFKeyAllResistExcept,
	"AllSkill":        BFKeyAllSkill,
	"MstuntRate":      BFKeyMstuntRate,

	// === 技能升级 ===
	"SkillDevelopExp":      BFKeySkillDevelopExp,
	"SkillDevelopIntimacy": BFKeySkillDevelopIntimacy,
	"LifeEffect":           BFKeyLifeEffect,
	"ManaEffect":           BFKeyManaEffect,
	"AttackEffect":         BFKeyAttackEffect,
	"SpeedEffect":          BFKeySpeedEffect,
	"PhyEffect":            BFKeyPhyEffect,
	"MagEffect":            BFKeyMagEffect,
	"PhyAbsorb":            BFKeyPhyAbsorb,
	"MagAbsorb":            BFKeyMagAbsorb,
	"AddMaxLife":           BFKeyAddMaxLife,
	"AddMaxMana":           BFKeyAddMaxMana,
	"AddUserLevel":         BFKeyAddUserLevel,
	"AddRandomSkill":       BFKeyAddRandomSkill,
	"DoubleTime":           BFKeyDoubleTime,
	"AddPetLevel":          BFKeyAddPetLevel,
	"PetAheadSkill":        BFKeyPetAheadSkill,
	"PetLongevity":         BFKeyPetLongevity,
	"UpgradeType":          BFKeyUpgradeType,
	"AddPetExp":            BFKeyAddPetExp,

	// === 房屋系统 ===
	"HouseId":               BFKeyHouseId,
	"HouseHouseClass":       BFKeyHouseHouseClass,
	"PlantLevel":            BFKeyPlantLevel,
	"PlantExp":              BFKeyPlantExp,
	"ToBeDeleted":           BFKeyToBeDeleted,
	"Locked":                BFKeyLocked,
	"PetUpgraded":           BFKeyPetUpgraded,
	"LeftTimeToDelete":      BFKeyLeftTimeToDelete,
	"ExtraDesc":             BFKeyExtraDesc,
	"PhyRebuildLevel":       BFKeyPhyRebuildLevel,
	"MagRebuildLevel":       BFKeyMagRebuildLevel,
	"RawName":               BFKeyRawName,
	"SuitPolar":             BFKeySuitPolar,
	"SuitEnabled":           BFKeySuitEnabled,
	"Gift":                  BFKeyGift,
	"RecognizeRecognized":   BFKeyRecognizeRecognized,
	"PartyStagePartyName":   BFKeyPartyStagePartyName,
	"PartyStagePassedCount": BFKeyPartyStagePassedCount,
	"PartyStageCostTime":    BFKeyPartyStageCostTime,
	"PartyStageMemberName":  BFKeyPartyStageMemberName,
	"Prop2Color":            BFKeyProp2Color,

	// === 摔跤系统 ===
	"WrestleScore":            BFKeyWrestleScore,
	"WrestleScore2":           BFKeyWrestleScore2,
	"DefEffect":               BFKeyDefEffect,
	"CombatGuard":             BFKeyCombatGuard,
	"Combined":                BFKeyCombined,
	"OpenNimbus":              BFKeyOpenNimbus,
	"LimitUseTimeIneffective": BFKeyLimitUseTimeIneffective,
	"WuhunqiaoLevel":          BFKeyWuhunqiaoLevel,
	"WeaponIcon":              BFKeyWeaponIcon,
	"SuitIcon":                BFKeySuitIcon,
	"OrgIcon":                 BFKeyOrgIcon,
	"TaoEffect":               BFKeyTaoEffect,
	"MountType":               BFKeyMountType,
	"SuitLightEffect":         BFKeySuitLightEffect,
	"SuitPolarPreview":        BFKeySuitPolarPreview,
	"Signature2":              BFKeySignature2,
	"InsiderTime":             BFKeyInsiderTime,
	"UserState":               BFKeyUserState,
	"AutoReply":               BFKeyAutoReply,
	"FriendImage":             BFKeyFriendImage,
	"Gid":                     BFKeyGid,
	"IidStr":                  BFKeyIidStr,
	"AutoFight":               BFKeyAutoFight,
	"FreeRename":              BFKeyFreeRename,
	"Voucher":                 BFKeyVoucher,
	"UseMoneyType":            BFKeyUseMoneyType,
	"LockExp":                 BFKeyLockExp,
	"ShuadaoJijiRulvling":     BFKeyShuadaoJijiRulvling,
	"FetchNice":               BFKeyFetchNice,
	"Recharge":                BFKeyRecharge,

	// === 额外效果 ===
	"ExtraLifeEffect":  BFKeyExtraLifeEffect,
	"ExtraManaEffect":  BFKeyExtraManaEffect,
	"ExtraMagEffect":   BFKeyExtraMagEffect,
	"ExtraPhyEffect":   BFKeyExtraPhyEffect,
	"ExtraSpeedEffect": BFKeyExtraSpeedEffect,

	// === 变化时间 ===
	"MorphLifeTimes":     BFKeyMorphLifeTimes,
	"MorphManaTimes":     BFKeyMorphManaTimes,
	"MorphSpeedTimes":    BFKeyMorphSpeedTimes,
	"MorphPhyTimes":      BFKeyMorphPhyTimes,
	"MorphMagTimes":      BFKeyMorphMagTimes,
	"MorphLifeStat":      BFKeyMorphLifeStat,
	"MorphManaStat":      BFKeyMorphManaStat,
	"MorphSpeedStat":     BFKeyMorphSpeedStat,
	"MorphPhyStat":       BFKeyMorphPhyStat,
	"MorphMagStat":       BFKeyMorphMagStat,
	"FreeUnlockExpTimes": BFKeyFreeUnlockExpTimes,
	"WeekAct":            BFKeyWeekAct,
	"ComebackFlag":       BFKeyComebackFlag,
	"PlacedAmount":       BFKeyPlacedAmount,
	"Achieve":            BFKeyAchieve,
	"AchieveName":        BFKeyAchieveName,
	"AchieveTime":        BFKeyAchieveTime,

	// === 升级系统 ===
	"UpgradeState":          BFKeyUpgradeState,
	"UpgradeType2":          BFKeyUpgradeType2,
	"UpgradeLevel":          BFKeyUpgradeLevel,
	"UpgradeExp":            BFKeyUpgradeExp,
	"UpgradeExpToNextLevel": BFKeyUpgradeExpToNextLevel,
	"UpgradeMaxPolarExtra":  BFKeyUpgradeMaxPolarExtra,
	"UpgradeLevel2":         BFKeyUpgradeLevel2,
	"HasUpgraded":           BFKeyHasUpgraded,
	"LimitUseTime":          BFKeyLimitUseTime,
	"FasionType":            BFKeyFasionType,
	"FoodNum":               BFKeyFoodNum,
	"MaxFoodNum":            BFKeyMaxFoodNum,
	"HouseId2":              BFKeyHouseId2,
	"Comfort":               BFKeyComfort,
	"CoupleName":            BFKeyCoupleName,
	"HouseType":             BFKeyHouseType,
	"SubType":               BFKeySubType,
	"CoupleGid":             BFKeyCoupleGid,

	// === 经验仓库 ===
	"ExpWareDataUnlockTime":      BFKeyExpWareDataUnlockTime,
	"ExpWareDataLockTime":        BFKeyExpWareDataLockTime,
	"ExpWareDataExpWare":         BFKeyExpWareDataExpWare,
	"ExpWareDataFetchTimes":      BFKeyExpWareDataFetchTimes,
	"ExpWareDataTodayFetchTimes": BFKeyExpWareDataTodayFetchTimes,
	"Stage":                      BFKeyStage,
	"Energy":                     BFKeyEnergy,

	// === 属性分配 ===
	"AttribAssignStr": BFKeyAttribAssignStr,
	"AttribAssignWiz": BFKeyAttribAssignWiz,
	"AttribAssignCon": BFKeyAttribAssignCon,
	"AttribAssignDex": BFKeyAttribAssignDex,
	"PhyShape":        BFKeyPhyShape,
	"MagShape":        BFKeyMagShape,
	"SpeedShape":      BFKeySpeedShape,
	"ManaShape":       BFKeyManaShape,
	"LifeShape":       BFKeyLifeShape,
	"UpgradeGid":      BFKeyUpgradeGid,
	"EnhancedPhy":     BFKeyEnhancedPhy,
	"IgnoreMagDodge":  BFKeyIgnoreMagDodge,
	"EnhancedMag":     BFKeyEnhancedMag,
	"Popular":         BFKeyPopular,

	// === 交易系统 ===
	"TradingGoodsGid":      BFKeyTradingGoodsGid,
	"TradingState":         BFKeyTradingState,
	"TradingLeftTime":      BFKeyTradingLeftTime,
	"TradingPrice":         BFKeyTradingPrice,
	"TradingOrgPrice":      BFKeyTradingOrgPrice,
	"TradingCgPriceTi":     BFKeyTradingCgPriceTi,
	"TradingCgPriceCt":     BFKeyTradingCgPriceCt,
	"CharOnlineState":      BFKeyCharOnlineState,
	"TradingSellBuyType":   BFKeyTradingSellBuyType,
	"TradingAppointeeName": BFKeyTradingAppointeeName,
	"TradingBuyoutPrice":   BFKeyTradingBuyoutPrice,

	// === 丹道系统 ===
	"DanDataState":          BFKeyDanDataState,
	"DanDataStage":          BFKeyDanDataStage,
	"DanDataExp":            BFKeyDanDataExp,
	"DanDataExpToNextLevel": BFKeyDanDataExpToNextLevel,
	"DanDataAttribPoint":    BFKeyDanDataAttribPoint,
	"DanDataPolarPoint":     BFKeyDanDataPolarPoint,
	"NotCheckBw":            BFKeyNotCheckBw,
	"HasBreakLvLimit":       BFKeyHasBreakLvLimit,
	"DanDataTodayExp":       BFKeyDanDataTodayExp,
	"SoulState":             BFKeySoulState,
	"HornName":              BFKeyHornName,
	"JewelryEssence":        BFKeyJewelryEssence,
	"TransformNum":          BFKeyTransformNum,
	"TransformCoolTi":       BFKeyTransformCoolTi,
	"MarriageStartTime":     BFKeyMarriageStartTime,
	"BookId":                BFKeyBookId,
	"FasionCustomDisable":   BFKeyFasionCustomDisable,
	"FasionEffectDisable":   BFKeyFasionEffectDisable,
	"MarriageBookId":        BFKeyMarriageBookId,
	"StrengthenJewelryNum":  BFKeyStrengthenJewelryNum,
	"StrengthenLevel":       BFKeyStrengthenLevel,
	"StrengthenExp":         BFKeyStrengthenExp,
	"StrengthenDegree":      BFKeyStrengthenDegree,
	"MonTao":                BFKeyMonTao,
	"MonTaoEx":              BFKeyMonTaoEx,
	"LastMonTao":            BFKeyLastMonTao,
	"LastMonTaoEx":          BFKeyLastMonTaoEx,
	"MonMartial":            BFKeyMonMartial,
	"LastMonMartial":        BFKeyLastMonMartial,
	"MonTaoRank":            BFKeyMonTaoRank,

	// === 神魂系统 ===
	"ShenhunDataState":          BFKeyShenhunDataState,
	"ShenhunDataLayer":          BFKeyShenhunDataLayer,
	"ShenhunDataExp":            BFKeyShenhunDataExp,
	"ShenhunDataExpToNextLevel": BFKeyShenhunDataExpToNextLevel,
	"YqzcNum":                   BFKeyYqzcNum,
	"CSkillDodge":               BFKeyCSkillDodge,
	"IgnoreCSkillDodge":         BFKeyIgnoreCSkillDodge,
	"StealBuffRate":             BFKeyStealBuffRate,
	"IgnoreStealBuffRate":       BFKeyIgnoreStealBuffRate,
	"ReliveRate":                BFKeyReliveRate,
	"IgnoreReliveRate":          BFKeyIgnoreReliveRate,
	"SuckBloodRate":             BFKeySuckBloodRate,
	"IgnoreSuckBloodRate":       BFKeyIgnoreSuckBloodRate,
	"CrippleRate":               BFKeyCrippleRate,
	"HunqiaoPage":               BFKeyHunqiaoPage,
	"FightCap":                  BFKeyFightCap,
	"FightCapWithoutIntimacy":   BFKeyFightCapWithoutIntimacy,
	"WuhunDataLevel":            BFKeyWuhunDataLevel,
	"WuhunDataExp":              BFKeyWuhunDataExp,
	"WuhunDataExpToNextLevel":   BFKeyWuhunDataExpToNextLevel,
	"WhzcNum":                   BFKeyWhzcNum,
	"WuhunqiaoPage":             BFKeyWuhunqiaoPage,
	"FightCapTotal":             BFKeyFightCapTotal,
	"CwCompeteFscore":           BFKeyCwCompeteFscore,
	"QmMoney":                   BFKeyQmMoney,
	"GhostdomChallengeLevel":    BFKeyGhostdomChallengeLevel,

	// === 重建属性 ===
	"PhyRebuildRate":    BFKeyPhyRebuildRate,
	"MagRebuildRate":    BFKeyMagRebuildRate,
	"LifeAddTemp":       BFKeyLifeAddTemp,
	"ManaAddTemp":       BFKeyManaAddTemp,
	"PhyPowerAddTemp":   BFKeyPhyPowerAddTemp,
	"MagPowerAddTemp":   BFKeyMagPowerAddTemp,
	"SpeedAddTemp":      BFKeySpeedAddTemp,
	"DefAddTemp":        BFKeyDefAddTemp,
	"PhyRebuildAdd":     BFKeyPhyRebuildAdd,
	"MagRebuildAdd":     BFKeyMagRebuildAdd,
	"PetLifeShapeTemp":  BFKeyPetLifeShapeTemp,
	"PetManaShapeTemp":  BFKeyPetManaShapeTemp,
	"PetSpeedShapeTemp": BFKeyPetSpeedShapeTemp,
	"PetPhyShapeTemp":   BFKeyPetPhyShapeTemp,
	"PetMagShapeTemp":   BFKeyPetMagShapeTemp,
	"EvolveDegree":      BFKeyEvolveDegree,
	"Power":             BFKeyPower,
	"SkillLowCost":      BFKeySkillLowCost,
	"RebuildDegree":     BFKeyRebuildDegree,
	"Time":              BFKeyTime,
	"EffectFoot":        BFKeyEffectFoot,
	"FightScore":        BFKeyFightScore,
	"RecvGid":           BFKeyRecvGid,
	"SuitLevel":         BFKeySuitLevel,
	"SuitDegree":        BFKeySuitDegree,

	// === 状态效果 ===
	"StatusPoison":           BFKeyStatusPoison,
	"StatusSleep":            BFKeyStatusSleep,
	"StatusForgotten":        BFKeyStatusForgotten,
	"StatusFrozen":           BFKeyStatusFrozen,
	"StatusConfusion":        BFKeyStatusConfusion,
	"StatusJointAttack":      BFKeyStatusJointAttack,
	"StatusRevive":           BFKeyStatusRevive,
	"StatusStunt":            BFKeyStatusStunt,
	"StatusDoubleHit":        BFKeyStatusDoubleHit,
	"StatusDamageSel":        BFKeyStatusDamageSel,
	"StatusCounterAttack":    BFKeyStatusCounterAttack,
	"StatusProtected":        BFKeyStatusProtected,
	"StatusSpeed":            BFKeyStatusSpeed,
	"StatusPhyPower":         BFKeyStatusPhyPower,
	"StatusDefense":          BFKeyStatusDefense,
	"StatusMaxLife":          BFKeyStatusMaxLife,
	"StatusDodge":            BFKeyStatusDodge,
	"StatusDef":              BFKeyStatusDef,
	"StatusRecoverLife":      BFKeyStatusRecoverLife,
	"StatusMetal":            BFKeyStatusMetal,
	"StatusWood":             BFKeyStatusWood,
	"StatusWater":            BFKeyStatusWater,
	"StatusFire":             BFKeyStatusFire,
	"StatusEarth":            BFKeyStatusEarth,
	"StatusLeechPhyDamage":   BFKeyStatusLeechPhyDamage,
	"StatusLeechMagDamage":   BFKeyStatusLeechMagDamage,
	"StatusPassiveAttack":    BFKeyStatusPassiveAttack,
	"StatusDeadlyKiss":       BFKeyStatusDeadlyKiss,
	"StatusLoyalty":          BFKeyStatusLoyalty,
	"StatusImmunePhyDamage":  BFKeyStatusImmunePhyDamage,
	"StatusImmuneMagDamage":  BFKeyStatusImmuneMagDamage,
	"StatusPolarChanged":     BFKeyStatusPolarChanged,
	"StatusFanzhuanQiankun":  BFKeyStatusFanzhuanQiankun,
	"StatusManaShield":       BFKeyStatusManaShield,
	"StatusPassiveMagAttack": BFKeyStatusPassiveMagAttack,
	"StatusAddLifeByMana":    BFKeyStatusAddLifeByMana,
	"CombatGuardIndex":       BFKeyCombatGuardIndex,
	"Salary":                 BFKeySalary,
	"Party":                  BFKeyParty,
	"Population":             BFKeyPopulation,
	"PartyWarWin":            BFKeyPartyWarWin,
	"LevelUpTime":            BFKeyLevelUpTime,
	"HigestXiangy":           BFKeyHigestXiangy,
	"HigestChub":             BFKeyHigestChub,
	"HigestFum":              BFKeyHigestFum,
	"HigestTongtt":           BFKeyHigestTongtt,
	"HigestYasby":            BFKeyHigestYasby,
	"TonttLayer":             BFKeyTonttLayer,
	"Money":                  BFKeyMoney,
	"TaoRank":                BFKeyTaoRank,
	"LotteryTimes":           BFKeyLotteryTimes,
	"ShadowSelf":             BFKeyShadowSelf,
	"Enchant":                BFKeyEnchant,
	"EnchantNimbus":          BFKeyEnchantNimbus,
	"MaxEnchantNimbus":       BFKeyMaxEnchantNimbus,
	"CardType":               BFKeyCardType,
	"EffectTime":             BFKeyEffectTime,
	"ShenmuPoints":           BFKeyShenmuPoints,
	"EnableShenmuPoints":     BFKeyEnableShenmuPoints,
	"GiftKey":                BFKeyGiftKey,
	"ExpiredTime":            BFKeyExpiredTime,
	"Evolve":                 BFKeyEvolve,
	"InsiderLevel":           BFKeyInsiderLevel,
	"EvolveLevel":            BFKeyEvolveLevel,
	"MailingItemTimes":       BFKeyMailingItemTimes,
	"MountAttribEndTime":     BFKeyMountAttribEndTime,
	"MountAttribMoveSpeed":   BFKeyMountAttribMoveSpeed,
	"CapacityLevel":          BFKeyCapacityLevel,
	"HideMount":              BFKeyHideMount,
	"Deadline":               BFKeyDeadline,
	"MergeRate":              BFKeyMergeRate,
	"EquipPerfectPercent":    BFKeyEquipPerfectPercent,
	"DunwuTimes":             BFKeyDunwuTimes,
	"DunwuRate":              BFKeyDunwuRate,
	"PetAnger":               BFKeyPetAnger,

	// === 状态效果(续) ===
	"StatusHuanbingZhiji":        BFKeyStatusHuanbingZhiji,
	"StatusAitongYujue":          BFKeyStatusAitongYujue,
	"StatusShushouJiuqin":        BFKeyStatusShushouJiuqin,
	"StatusWenfengSangdan":       BFKeyStatusWenfengSangdan,
	"StatusYangjingXurui":        BFKeyStatusYangjingXurui,
	"StatusXuwu":                 BFKeyStatusXuwu,
	"ItemPolar":                  BFKeyItemPolar,
	"ShuadaoZiqihongmeng":        BFKeyShuadaoZiqihongmeng,
	"ExtraSkill":                 BFKeyExtraSkill,
	"ExtraSkillLevel":            BFKeyExtraSkillLevel,
	"DiandqkFrozenRound":         BFKeyDiandqkFrozenRound,
	"StatusChaofeng":             BFKeyStatusChaofeng,
	"StatusDiandaoQiankun":       BFKeyStatusDiandaoQiankun,
	"StatusJingangquan":          BFKeyStatusJingangquan,
	"StatusQinmiWujian":          BFKeyStatusQinmiWujian,
	"StatusTianyan":              BFKeyStatusTianyan,
	"StatusWujiBifan":            BFKeyStatusWujiBifan,
	"StatusShowOpponentLife":     BFKeyStatusShowOpponentLife,
	"StatusAddLifeByManaAdd":     BFKeyStatusAddLifeByManaAdd,
	"StatusRecoverLifeAdd":       BFKeyStatusRecoverLifeAdd,
	"StatusDefAdd":               BFKeyStatusDefAdd,
	"StatusPhyPowerAdd":          BFKeyStatusPhyPowerAdd,
	"StatusMagPowerAdd":          BFKeyStatusMagPowerAdd,
	"StatusSpeedAdd":             BFKeyStatusSpeedAdd,
	"BrotherAppellation":         BFKeyBrotherAppellation,
	"ShuadaoRuyiPoint":           BFKeyShuadaoRuyiPoint,
	"ChushiEx":                   BFKeyChushiEx,
	"PhyPowerWithoutIntimacy":    BFKeyPhyPowerWithoutIntimacy,
	"MagPowerWithoutIntimacy":    BFKeyMagPowerWithoutIntimacy,
	"DefWithoutIntimacy":         BFKeyDefWithoutIntimacy,
	"OriginIntimacy":             BFKeyOriginIntimacy,
	"DouchongRank":               BFKeyDouchongRank,
	"UpgradeImmortal":            BFKeyUpgradeImmortal,
	"UpgradeMagic":               BFKeyUpgradeMagic,
	"UpgradeTotal":               BFKeyUpgradeTotal,
	"ArtifactUpgradedEnabled":    BFKeyArtifactUpgradedEnabled,
	"StatusQishaYin":             BFKeyStatusQishaYin,
	"StatusQishaYang":            BFKeyStatusQishaYang,
	"Eclosion":                   BFKeyEclosion,
	"EclosionNimbus":             BFKeyEclosionNimbus,
	"MaxEclosionNimbus":          BFKeyMaxEclosionNimbus,
	"StatusAllResistExceptAdd":   BFKeyStatusAllResistExceptAdd,
	"EclosionStage":              BFKeyEclosionStage,
	"StatusYanchuanShenjiao":     BFKeyStatusYanchuanShenjiao,
	"BossAnger":                  BFKeyBossAnger,
	"HigestScore":                BFKeyHigestScore,
	"PartyId":                    BFKeyPartyId,
	"StatusDaofaWubian":          BFKeyStatusDaofaWubian,
	"ObtainTime":                 BFKeyObtainTime,
	"EffectWaist":                BFKeyEffectWaist,
	"EffectHead":                 BFKeyEffectHead,
	"StatusWeiya":                BFKeyStatusWeiya,
	"StatusWeiYaCount":           BFKeyStatusWeiYaCount,
	"GraphicInstructionMark":     BFKeyGraphicInstructionMark,
	"StatusDilieboFlag":          BFKeyStatusDilieboFlag,
	"PartIndex":                  BFKeyPartIndex,
	"PartColorIndex":             BFKeyPartColorIndex,
	"NpcChat":                    BFKeyNpcChat,
	"FollowPetType":              BFKeyFollowPetType,
	"DyeIcon":                    BFKeyDyeIcon,
	"FasionId":                   BFKeyFasionId,
	"FasionVisible":              BFKeyFasionVisible,
	"ActCamp":                    BFKeyActCamp,
	"LingchenPoint":              BFKeyLingchenPoint,
	"NpcState":                   BFKeyNpcState,
	"ChannelSource":              BFKeyChannelSource,
	"FlagChild":                  BFKeyFlagChild,
	"GhostGas":                   BFKeyGhostGas,
	"PeiyuanDataLevel":           BFKeyPeiyuanDataLevel,
	"PeiyuanDataStage":           BFKeyPeiyuanDataStage,
	"NingshenState":              BFKeyNingshenState,
	"PeiyuanDataState":           BFKeyPeiyuanDataState,
	"GongshengDataTargetPetIid":  BFKeyGongshengDataTargetPetIid,
	"SkillLevel":                 BFKeySkillLevel,
	"UpgradeDegree":              BFKeyUpgradeDegree,
	"StatusFengmangAdd":          BFKeyStatusFengmangAdd,
	"StatusGuibuAdd":             BFKeyStatusGuibuAdd,
	"StatusBomuAdd":              BFKeyStatusBomuAdd,
	"StatusFengmang":             BFKeyStatusFengmang,
	"StatusHundeng":              BFKeyStatusHundeng,
	"StatusGuibu":                BFKeyStatusGuibu,
	"StatusBomu":                 BFKeyStatusBomu,
	"StatusFuhu":                 BFKeyStatusFuhu,
	"YoutjEffectFlag":            BFKeyYoutjEffectFlag,
	"YoutjSunPoint":              BFKeyYoutjSunPoint,
	"YoutjMoonPoint":             BFKeyYoutjMoonPoint,
	"HigestXiangy2":              BFKeyHigestXiangy2,
	"HigestFum2":                 BFKeyHigestFum2,
	"HigestFeixdx2":              BFKeyHigestFeixdx2,
	"StatusTiandaoHuti":          BFKeyStatusTiandaoHuti,
	"Ulevel":                     BFKeyUlevel,
	"MaxStoreExp":                BFKeyMaxStoreExp,
	"DigongMaxPassLevel":         BFKeyDigongMaxPassLevel,
	"StatusLihunZhihuo":          BFKeyStatusLihunZhihuo,
	"StatusXuanwuYin":            BFKeyStatusXuanwuYin,
	"ItemMaxLiveness":            BFKeyItemMaxLiveness,
	"ItemLiveness":               BFKeyItemLiveness,
	"ItemMaxCostGoldCoin":        BFKeyItemMaxCostGoldCoin,
	"ItemCostGoldCoin":           BFKeyItemCostGoldCoin,
	"StatusYishouWeigong":        BFKeyStatusYishouWeigong,
	"StatusYigongWeishou":        BFKeyStatusYigongWeishou,
	"ZhenlingType":               BFKeyZhenlingType,
	"ZhenlingLevel":              BFKeyZhenlingLevel,
	"StatusXielingzhiyan":        BFKeyStatusXielingzhiyan,
	"StatusTongshengXueshi":      BFKeyStatusTongshengXueshi,
	"StatusGongsiZuzhou":         BFKeyStatusGongsiZuzhou,
	"StatusShaluZhixin":          BFKeyStatusShaluZhixin,
	"StatusMeihuo":               BFKeyStatusMeihuo,
	"StatusJiuchan":              BFKeyStatusJiuchan,
	"StatusShashu":               BFKeyStatusShashu,
	"StatusFatianXiangdi":        BFKeyStatusFatianXiangdi,
	"StatusZhanxian":             BFKeyStatusZhanxian,
	"StatusZhanxianAdd":          BFKeyStatusZhanxianAdd,
	"OverseaAccount":             BFKeyOverseaAccount,
	"JcxBalance":                 BFKeyJcxBalance,
	"MaxJcxBalance":              BFKeyMaxJcxBalance,
	"CanUsePetAnger":             BFKeyCanUsePetAnger,
	"NotAllowMove":               BFKeyNotAllowMove,
	"HudunNum":                   BFKeyHudunNum,
	"LiehunMoyinCount":           BFKeyLiehunMoyinCount,
	"PhantomIcon":                BFKeyPhantomIcon,
	"StatusShihun":               BFKeyStatusShihun,
	"CombatTitle":                BFKeyCombatTitle,
	"ShowLifeBarBoss":            BFKeyShowLifeBarBoss,
	"WuxingMozhen":               BFKeyWuxingMozhen,
	"StatusChunyangZhenhuo":      BFKeyStatusChunyangZhenhuo,
	"StatusZhuoshao":             BFKeyStatusZhuoshao,
	"StatusZhuoshaoAdd":          BFKeyStatusZhuoshaoAdd,
	"Opacity":                    BFKeyOpacity,
	"CwYayun":                    BFKeyCwYayun,
	"StatusShuanghanQinxi":       BFKeyStatusShuanghanQinxi,
	"StatusXinwuPangwu":          BFKeyStatusXinwuPangwu,
	"StatusGubenPeiyuan":         BFKeyStatusGubenPeiyuan,
	"StatusGuwuShiqi":            BFKeyStatusGuwuShiqi,
	"StatusShuanghanQinxiAdd":    BFKeyStatusShuanghanQinxiAdd,
	"StatusXinwuPangwuNum":       BFKeyStatusXinwuPangwuNum,
	"StatusGubenPeiyuanAdd":      BFKeyStatusGubenPeiyuanAdd,
	"StatusGuwuShiqiAdd":         BFKeyStatusGuwuShiqiAdd,
	"StatusWandaoChengkong":      BFKeyStatusWandaoChengkong,
	"WufaWushuang":               BFKeyWufaWushuang,
	"JianruPanshi":               BFKeyJianruPanshi,
	"XinghuoFeichi":              BFKeyXinghuoFeichi,
	"NingqiHuadun":               BFKeyNingqiHuadun,
	"CanFly":                     BFKeyCanFly,
	"StrengthenDegree80":         BFKeyStrengthenDegree80,
	"StatusRecoverMana":          BFKeyStatusRecoverMana,
	"StatusRecoverManaAdd":       BFKeyStatusRecoverManaAdd,
	"Camp":                       BFKeyCamp,
	"ZhengdaoTaskIndex":          BFKeyZhengdaoTaskIndex,
	"StatusZd05":                 BFKeyStatusZd05,
	"StatusYinyangjing":          BFKeyStatusYinyangjing,
	"StatusYinyangjingCount":     BFKeyStatusYinyangjingCount,
	"StatusXiejiaJinhu":          BFKeyStatusXiejiaJinhu,
	"StatusXiejiaJinhuSkill":     BFKeyStatusXiejiaJinhuSkill,
	"YubxsLevel":                 BFKeyYubxsLevel,
	"ShidaoDahui":                BFKeyShidaoDahui,
	"HigestFeixdx":               BFKeyHigestFeixdx,
	"OpenState":                  BFKeyOpenState,
	"StatusDiandaoCuoluanAdd":    BFKeyStatusDiandaoCuoluanAdd,
	"StatusShuanghanZhihuAdd":    BFKeyStatusShuanghanZhihuAdd,
	"Status":                     BFKeyStatus,
	"CscwQiaozhuang":             BFKeyCscwQiaozhuang,
	"MarriageMarryId":            BFKeyMarriageMarryId,
	"Alias":                      BFKeyAlias,
	"ShuadaoChongfengSan":        BFKeyShuadaoChongfengSan,
	"GroupName":                  BFKeyGroupName,
	"GroupId":                    BFKeyGroupId,
	"LeaderGid":                  BFKeyLeaderGid,
	"MemberGid":                  BFKeyMemberGid,
	"Setting":                    BFKeySetting,
	"Announcement":               BFKeyAnnouncement,
	"SettingRefuseStrangerLevel": BFKeySettingRefuseStrangerLevel,
	"SettingAutoReplyMsg":        BFKeySettingAutoReplyMsg,
	"SettingRefuseBeAddLevel":    BFKeySettingRefuseBeAddLevel,
	"ServerName":                 BFKeyServerName,
	"BullyKillNum":               BFKeyBullyKillNum,
	"PoliceKillNum":              BFKeyPoliceKillNum,
	"ShowSandglass":              BFKeyShowSandglass,

	// === GM属性 ===
	"GmAttribsMaxLife":      BFKeyGmAttribsMaxLife,
	"GmAttribsMaxMana":      BFKeyGmAttribsMaxMana,
	"GmAttribsPhyPower":     BFKeyGmAttribsPhyPower,
	"GmAttribsMagPower":     BFKeyGmAttribsMagPower,
	"GmAttribsDef":          BFKeyGmAttribsDef,
	"GmAttribsSpeed":        BFKeyGmAttribsSpeed,
	"MarriageCoupleGid":     BFKeyMarriageCoupleGid,
	"ChatHead":              BFKeyChatHead,
	"ChatFloor":             BFKeyChatFloor,
	"DistName":              BFKeyDistName,
	"ChannelFilterType":     BFKeyChannelFilterType,
	"RoomName":              BFKeyRoomName,
	"TeamAssessScore":       BFKeyTeamAssessScore,
	"ChannelActType":        BFKeyChannelActType,
	"BanRule":               BFKeyBanRule,
	"AiteParty":             BFKeyAiteParty,
	"MessageId":             BFKeyMessageId,
	"SpecialMonsterChannel": BFKeySpecialMonsterChannel,
}

// LookupBuildFieldKey 根据常量名查找数字 key
// 如果找不到，尝试解析为数字（兼容旧格式）
func LookupBuildFieldKey(name string) (int16, bool) {
	// 1. 先在映射表中查找常量名
	if key, ok := buildFieldKeyMap[name]; ok {
		return key, true
	}
	// 2. 兼容旧格式：直接解析数字
	var key int16
	n, err := parseKeyAsInt16(name, &key)
	if err == nil && n == 1 {
		return key, true
	}
	return 0, false
}

// RegisterBuildFieldKey 注册一个 BuildField key 到映射表
// 用于动态注册新的 key（如通过 init() 函数）
func RegisterBuildFieldKey(name string, key int16) {
	buildFieldKeyMap[name] = key
}

// parseKeyAsInt16 尝试将字符串解析为 int16
func parseKeyAsInt16(s string, key *int16) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	var n int64
	var i int
	if s[0] == '-' {
		i = 1
	}
	if i >= len(s) {
		return 0, nil // 只有负号，没有数字
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, nil // 遇到非数字，返回失败
		}
		n = n*10 + int64(c-'0')
	}
	if s[0] == '-' {
		n = -n
	}
	if n < -32768 || n > 32767 {
		return 0, nil
	}
	*key = int16(n)
	return 1, nil
}
