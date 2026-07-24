package system

import (
	"testing"
	"wgame-server/server/codec"
)

// TestMsgCharAlreadyLogin 测试 MSG_CHAR_ALREADY_LOGIN 消息编解码
func TestMsgCharAlreadyLogin(t *testing.T) {
	original := &MsgCharAlreadyLogin{Name: "测试角色"}

	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgCharAlreadyLogin{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.Name != original.Name {
		t.Errorf("名称不匹配: 期望 %s, 实际 %s", original.Name, decoded.Name)
	}
}

// TestMsgCsServerType 测试 MSG_CS_SERVER_TYPE 消息编解码
func TestMsgCsServerType(t *testing.T) {
	original := &MsgCsServerType{ServerType: 1}

	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgCsServerType{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.ServerType != original.ServerType {
		t.Errorf("服务器类型不匹配: 期望 %d, 实际 %d", original.ServerType, decoded.ServerType)
	}
}

// TestMsgGeneralNotify 测试 MSG_GENERAL_NOTIFY 消息编解码
func TestMsgGeneralNotify(t *testing.T) {
	original := &MsgGeneralNotify{Notify: 60001, Para: "test"}

	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgGeneralNotify{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.Notify != original.Notify {
		t.Errorf("通知ID不匹配: 期望 %d, 实际 %d", original.Notify, decoded.Notify)
	}
	if decoded.Para != original.Para {
		t.Errorf("参数不匹配: 期望 %s, 实际 %s", original.Para, decoded.Para)
	}
}

// TestMsgSetPushSettings 测试 MSG_SET_PUSH_SETTINGS 消息编解码
func TestMsgSetPushSettings(t *testing.T) {
	original := &MsgSetPushSettings{Value: "10011011111"}

	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgSetPushSettings{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.Value != original.Value {
		t.Errorf("值不匹配: 期望 %s, 实际 %s", original.Value, decoded.Value)
	}
}

// TestMsgFuzzyIdentity 测试 MSG_FUZZY_IDENTITY 消息编解码
func TestMsgFuzzyIdentity(t *testing.T) {
	original := &MsgFuzzyIdentity{
		IsBindName:  1,
		IsBindPhone: 1,
		BindName:    "王老板",
		BindId:      "110101192003127055",
		BindPhone:   "130****6666",
	}

	w := codec.NewGameWriter(128)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgFuzzyIdentity{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.IsBindName != original.IsBindName {
		t.Errorf("IsBindName 不匹配: 期望 %d, 实际 %d", original.IsBindName, decoded.IsBindName)
	}
	if decoded.IsBindPhone != original.IsBindPhone {
		t.Errorf("IsBindPhone 不匹配: 期望 %d, 实际 %d", original.IsBindPhone, decoded.IsBindPhone)
	}
	if decoded.BindName != original.BindName {
		t.Errorf("BindName 不匹配: 期望 %s, 实际 %s", original.BindName, decoded.BindName)
	}
	if decoded.BindId != original.BindId {
		t.Errorf("BindId 不匹配: 期望 %s, 实际 %s", original.BindId, decoded.BindId)
	}
	if decoded.BindPhone != original.BindPhone {
		t.Errorf("BindPhone 不匹配: 期望 %s, 实际 %s", original.BindPhone, decoded.BindPhone)
	}
}

// TestMsgExecuteLuaCode 测试 MSG_EXECUTE_LUA_CODE 消息编解码
func TestMsgExecuteLuaCode(t *testing.T) {
	original := &MsgExecuteLuaCode{
		Cookie: 0,
		Code:   "GameMgr.checkServer = function()\nreturn 0\nend",
		Flag:   1,
	}

	w := codec.NewGameWriter(256)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgExecuteLuaCode{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.Cookie != original.Cookie {
		t.Errorf("Cookie 不匹配: 期望 %d, 实际 %d", original.Cookie, decoded.Cookie)
	}
	if decoded.Code != original.Code {
		t.Errorf("Code 不匹配: 期望 %s, 实际 %s", original.Code, decoded.Code)
	}
	if decoded.Flag != original.Flag {
		t.Errorf("Flag 不匹配: 期望 %d, 实际 %d", original.Flag, decoded.Flag)
	}
}

// TestMsgUpdate 测试 MSG_UPDATE 消息编解码
func TestMsgUpdate(t *testing.T) {
	original := &MsgUpdate{
		Id: 12345,
		Props: MsgUpdateProps{
			Name:     "测试角色",
			Level:    69,
			Polar:    1,
			Icon:     100,
			Str:      100,
			Dex:      50,
			Con:      80,
			Wiz:      90,
			Life:     50000,
			MaxLife:  50000,
			Mana:     20000,
			MaxMana:  20000,
			Metal:    10,
			Wood:     20,
			Water:    30,
			Fire:     40,
			Earth:    50,
			Cash:     1000000,
			Balance:  500,
			GoldCoin: 9999999,
			Nice:     1000,
			Tao:      365,
			Online:   1,
		},
	}

	w := codec.NewGameWriter(512)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgUpdate{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if decoded.Id != original.Id {
		t.Errorf("Id 不匹配: 期望 %d, 实际 %d", original.Id, decoded.Id)
	}
	if decoded.Props.Name != original.Props.Name {
		t.Errorf("Name 不匹配: 期望 %s, 实际 %s", original.Props.Name, decoded.Props.Name)
	}
	if decoded.Props.Level != original.Props.Level {
		t.Errorf("Level 不匹配: 期望 %d, 实际 %d", original.Props.Level, decoded.Props.Level)
	}
	if decoded.Props.Polar != original.Props.Polar {
		t.Errorf("Polar 不匹配: 期望 %d, 实际 %d", original.Props.Polar, decoded.Props.Polar)
	}
	if decoded.Props.Life != original.Props.Life {
		t.Errorf("Life 不匹配: 期望 %d, 实际 %d", original.Props.Life, decoded.Props.Life)
	}
	if decoded.Props.MaxLife != original.Props.MaxLife {
		t.Errorf("MaxLife 不匹配: 期望 %d, 实际 %d", original.Props.MaxLife, decoded.Props.MaxLife)
	}
	if decoded.Props.Cash != original.Props.Cash {
		t.Errorf("Cash 不匹配: 期望 %d, 实际 %d", original.Props.Cash, decoded.Props.Cash)
	}
	if decoded.Props.GoldCoin != original.Props.GoldCoin {
		t.Errorf("GoldCoin 不匹配: 期望 %d, 实际 %d", original.Props.GoldCoin, decoded.Props.GoldCoin)
	}
}

// TestMsgSetSetting 测试 MSG_SET_SETTING 消息编解码
func TestMsgSetSetting(t *testing.T) {
	original := &MsgSetSetting{
		Items: []MsgSetSettingItem{
			{Key: "key1", Value: 1},
			{Key: "key2", Value: 2},
			{Key: "key3", Value: 3},
		},
	}

	w := codec.NewGameWriter(256)
	if err := codec.AutoWrite(w, original); err != nil {
		t.Fatalf("写入失败: %v", err)
	}

	reader := codec.NewGameReader(w.Bytes())
	decoded := &MsgSetSetting{}
	if err := codec.AutoRead(reader, decoded); err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if len(decoded.Items) != len(original.Items) {
		t.Fatalf("Items 长度不匹配: 期望 %d, 实际 %d", len(original.Items), len(decoded.Items))
	}
	for i, item := range original.Items {
		if decoded.Items[i].Key != item.Key {
			t.Errorf("Items[%d].Key 不匹配: 期望 %s, 实际 %s", i, item.Key, decoded.Items[i].Key)
		}
		if decoded.Items[i].Value != item.Value {
			t.Errorf("Items[%d].Value 不匹配: 期望 %d, 实际 %d", i, item.Value, decoded.Items[i].Value)
		}
	}
}

// TestNewMsgSetSetting 测试创建默认设置消息
func TestNewMsgSetSetting(t *testing.T) {
	msg := NewMsgSetSetting()
	if len(msg.Items) != 62 {
		t.Errorf("Items 长度应该为 62, 实际为 %d", len(msg.Items))
	}
}

// TestNewMsgUpdate 测试创建更新消息
func TestNewMsgUpdate(t *testing.T) {
	msg := NewMsgUpdate(123)
	if msg.Id != 123 {
		t.Errorf("Id 不匹配: 期望 123, 实际 %d", msg.Id)
	}
}
