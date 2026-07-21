package game

// Goods 表示物品（装备、道具等）。
// 参考 Java wd-server-fl core/domain/goods/Goods.java。
type Goods struct {
	Id          int32
	No          int32
	Name        string
	Type        int32
	SubType     int32
	Level       int32
	Quality     int32
	Polar       int32
	Pos         int32 // 穿戴位置 1-40
	Num         int32
	IsBind      int32
	ZbAttribute *ZbAttribute
}

// Petbeibao 表示宠物背包项。
// 参考 Java wd-server-fl core/domain/Petbeibao.java。
type Petbeibao struct {
	Id           int32
	No           int32
	Pet_status   int32 // 0=休息, 1=参战, 2=掠阵, 3=坐骑
	Name         string
	PetTypeId    int32
}

// ShouHu 表示守护。
// 参考 Java wd-server-fl core/domain/ShouHu.java。
type ShouHu struct {
	Id    int32
	No    int32
	Name  string
	Level int32
}

// JiNeng 表示技能。
// 参考 Java wd-server-fl core/domain/JiNeng.java。
type JiNeng struct {
	Skill_no    int32
	Skill_name  string
	Skill_level int32
}

// ZbAttribute 表示装备属性。
// 参考 Java wd-server-fl core/domain/ZbAttribute.java。
type ZbAttribute struct {
	PhyPower int32 // 物理攻击
	MagPower int32 // 魔法攻击
	Def      int32 // 防御
	Speed    int32 // 速度
}

// NewZbAttribute 创建新的装备属性对象
func NewZbAttribute() *ZbAttribute {
	return &ZbAttribute{}
}
