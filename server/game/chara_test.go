package game

import (
	"testing"
)

// TestCharaAddCash 测试 AddCash 方法的边界条件
func TestCharaAddCash(t *testing.T) {
	tests := []struct {
		name     string
		initCash int32
		addNum   int64
		wantCash int32
	}{
		{"normal add", 100, 50, 150},
		{"overflow", 1999999999, 1000, 2000000000},
		{"exact max", 2000000000, 0, 2000000000},
		{"negative add", 100, -50, 50},
		{"negative to zero", 50, -100, 0},
		{"large negative", 100, -1000000000, 0},
		{"zero add", 100, 0, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Cash: tt.initCash}
			c.AddCash(tt.addNum)
			if c.Cash != tt.wantCash {
				t.Errorf("AddCash(%d) to %d: got %d, want %d", tt.addNum, tt.initCash, c.Cash, tt.wantCash)
			}
		})
	}
}

// TestCharaAddPot 测试 AddPot 方法的边界条件
func TestCharaAddPot(t *testing.T) {
	tests := []struct {
		name     string
		initPot  int32
		addNum   int64
		wantPot  int32
		wantBool bool
	}{
		{"normal add", 100, 50, 150, true},
		{"overflow", 1999999999, 1000, 2000000000, true},
		{"exact max", 2000000000, 0, 2000000000, true},
		{"zero add", 100, 0, 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Pot: tt.initPot}
			result := c.AddPot(tt.addNum)
			if c.Pot != tt.wantPot {
				t.Errorf("AddPot(%d) to %d: got %d, want %d", tt.addNum, tt.initPot, c.Pot, tt.wantPot)
			}
			if result != tt.wantBool {
				t.Errorf("AddPot return: got %v, want %v", result, tt.wantBool)
			}
		})
	}
}

// TestCharaAddExp 测试 AddExp 方法的边界条件
func TestCharaAddExp(t *testing.T) {
	tests := []struct {
		name     string
		initExp  int64
		addNum   int64
		wantExp  int64
	}{
		{"normal add", 100, 50, 150},
		{"overflow", 1999999999, 1000, 2000000000},
		{"exact max", 2000000000, 0, 2000000000},
		{"negative add", 100, -50, 50},
		{"negative to zero", 50, -100, 0},
		{"zero add", 100, 0, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Exp: tt.initExp}
			c.AddExp(tt.addNum)
			if c.Exp != tt.wantExp {
				t.Errorf("AddExp(%d) to %d: got %d, want %d", tt.addNum, tt.initExp, c.Exp, tt.wantExp)
			}
		})
	}
}

// TestCharaMagPowerTotal 测试魔法攻击计算（含装备加成）
func TestCharaMagPowerTotal(t *testing.T) {
	tests := []struct {
		name        string
		magPower    int32
		zbMagPower  int32
		zbNil       bool
		wantTotal   int32
	}{
		{"no equipment", 100, 0, true, 100},
		{"with equipment", 100, 50, false, 150},
		{"zero equipment bonus", 100, 0, false, 100},
		{"negative equipment bonus", 100, -10, false, 90},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{MagPower: tt.magPower}
			if !tt.zbNil {
				c.ZbAttribute = &ZbAttribute{MagPower: tt.zbMagPower}
			}
			if got := c.MagPowerTotal(); got != tt.wantTotal {
				t.Errorf("MagPowerTotal(): got %d, want %d", got, tt.wantTotal)
			}
		})
	}
}

// TestCharaPhyPowerTotal 测试物理攻击计算（含装备加成）
func TestCharaPhyPowerTotal(t *testing.T) {
	tests := []struct {
		name        string
		phyPower    int32
		zbPhyPower  int32
		zbNil       bool
		wantTotal   int32
	}{
		{"no equipment", 100, 0, true, 100},
		{"with equipment", 100, 50, false, 150},
		{"zero equipment bonus", 100, 0, false, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{PhyPower: tt.phyPower}
			if !tt.zbNil {
				c.ZbAttribute = &ZbAttribute{PhyPower: tt.zbPhyPower}
			}
			if got := c.PhyPowerTotal(); got != tt.wantTotal {
				t.Errorf("PhyPowerTotal(): got %d, want %d", got, tt.wantTotal)
			}
		})
	}
}

// TestCharaDefTotal 测试防御计算（含装备加成）
func TestCharaDefTotal(t *testing.T) {
	tests := []struct {
		name      string
		def       int32
		zbDef     int32
		zbNil     bool
		wantTotal int32
	}{
		{"no equipment", 100, 0, true, 100},
		{"with equipment", 100, 50, false, 150},
		{"zero equipment bonus", 100, 0, false, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Def: tt.def}
			if !tt.zbNil {
				c.ZbAttribute = &ZbAttribute{Def: tt.zbDef}
			}
			if got := c.DefTotal(); got != tt.wantTotal {
				t.Errorf("DefTotal(): got %d, want %d", got, tt.wantTotal)
			}
		})
	}
}

// TestCharaSpeedTotal 测试速度计算（含装备加成）
func TestCharaSpeedTotal(t *testing.T) {
	tests := []struct {
		name      string
		speed     int32
		zbSpeed   int32
		zbNil     bool
		wantTotal int32
	}{
		{"no equipment", 100, 0, true, 100},
		{"with equipment", 100, 50, false, 150},
		{"zero equipment bonus", 100, 0, false, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Speed: tt.speed}
			if !tt.zbNil {
				c.ZbAttribute = &ZbAttribute{Speed: tt.zbSpeed}
			}
			if got := c.SpeedTotal(); got != tt.wantTotal {
				t.Errorf("SpeedTotal(): got %d, want %d", got, tt.wantTotal)
			}
		})
	}
}

// TestCharaHasCoin 测试元宝检查
func TestCharaHasCoin(t *testing.T) {
	tests := []struct {
		name       string
		silverCoin int32
		goldCoin   int32
		checkNum   int64
		want       bool
	}{
		{"enough silver", 100, 0, 50, true},
		{"enough gold", 0, 100, 50, true},
		{"enough combined", 50, 50, 100, true},
		{"exact amount", 50, 50, 100, true},
		{"not enough", 50, 50, 101, false},
		{"zero check", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{SilverCoin: tt.silverCoin, GoldCoin: tt.goldCoin}
			if got := c.HasCoin(tt.checkNum); got != tt.want {
				t.Errorf("HasCoin(%d): got %v, want %v", tt.checkNum, got, tt.want)
			}
		})
	}
}

// TestCharaSubSilverCoin 测试元宝扣除
func TestCharaSubSilverCoin(t *testing.T) {
	tests := []struct {
		name           string
		silverCoin     int32
		goldCoin       int32
		subNum         int64
		wantSilver     int32
		wantGold       int32
		wantSuccess    bool
	}{
		{"sub from silver only", 100, 0, 50, 50, 0, true},
		{"sub all silver", 100, 0, 100, 0, 0, true},
		{"sub from gold when silver not enough", 50, 100, 75, 0, 75, true},
		{"sub from both", 50, 50, 100, 0, 0, true},
		{"not enough total", 50, 50, 101, 50, 50, false},
		{"sub zero", 100, 100, 0, 100, 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{SilverCoin: tt.silverCoin, GoldCoin: tt.goldCoin}
			result := c.SubSilverCoin(tt.subNum)
			if c.SilverCoin != tt.wantSilver {
				t.Errorf("SilverCoin: got %d, want %d", c.SilverCoin, tt.wantSilver)
			}
			if c.GoldCoin != tt.wantGold {
				t.Errorf("GoldCoin: got %d, want %d", c.GoldCoin, tt.wantGold)
			}
			if result != tt.wantSuccess {
				t.Errorf("result: got %v, want %v", result, tt.wantSuccess)
			}
		})
	}
}

// TestCharaNewChara 测试新角色创建
func TestCharaNewChara(t *testing.T) {
	c := NewChara("test", 1, 1, "gid123")
	if c.Name != "test" {
		t.Errorf("Name: got %s, want test", c.Name)
	}
	if c.Sex != 1 {
		t.Errorf("Sex: got %d, want 1", c.Sex)
	}
	if c.Polar != 1 {
		t.Errorf("Polar: got %d, want 1", c.Polar)
	}
	if c.Gid != "gid123" {
		t.Errorf("Gid: got %s, want gid123", c.Gid)
	}
	if c.Level != 1 {
		t.Errorf("Level: got %d, want 1", c.Level)
	}
	if c.ZbAttribute == nil {
		t.Error("ZbAttribute should not be nil")
	}
}

// TestCharaWaiguanByPolarAndSex 测试外观设置
func TestCharaWaiguanByPolarAndSex(t *testing.T) {
	tests := []struct {
		name    string
		polar   int32
		sex     int32
		wantWg  int32
	}{
		{"male 11", 11, 1, 20033},
		{"male 1", 1, 1, 6001},
		{"male 2", 2, 1, 7002},
		{"male 3", 3, 1, 7003},
		{"male 4", 4, 1, 6004},
		{"male 5", 5, 1, 6005},
		{"female 1", 1, 2, 7001},
		{"female 2", 2, 2, 6002},
		{"female 3", 3, 2, 6003},
		{"female 4", 4, 2, 7004},
		{"female 5", 5, 2, 7005},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chara{Polar: tt.polar, Sex: tt.sex}
			c.WaiguanByPolarAndSex()
			if c.Waiguan != tt.wantWg {
				t.Errorf("WaiguanByPolarAndSex(%d, %d): got %d, want %d", tt.polar, tt.sex, c.Waiguan, tt.wantWg)
			}
		})
	}
}
