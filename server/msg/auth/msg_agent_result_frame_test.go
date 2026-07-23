package auth

import (
	"fmt"
	"testing"
	"wgame-server/server/msg"
)

func TestMsgAgentResultFrame(t *testing.T) {
	m := &MsgAgentResult{
		Result:       1,
		ID:           1,
		Privilege:    0,
		IP:           "192.168.1.63",
		Port:         8800,
		ServerName:   "test0",
		ServerStatus: 1,
		Msg:          "ok",
	}

	data, err := msg.WriteFrame(m, -1, 0)
	if err != nil {
		t.Fatalf("WriteFrame failed: %v", err)
	}

	fmt.Println("=== Full Frame (with Header) ===")
	fmt.Printf("Total length: %d\n", len(data))
	fmt.Printf("Hex: %x\n", data)
	fmt.Println()

	fmt.Println("Frame structure:")
	fmt.Printf("  [00-01] magic: 0x%04X = %d\n", (data[0]<<8)|data[1], (data[0]<<8)|data[1])
	fmt.Printf("  [02-03] tableIndex: 0x%04X = %d\n", (data[2]<<8)|data[3], (data[2]<<8)|data[3])
	fmt.Printf("  [04-07] tickCount: 0x%08X = %d\n", (data[4]<<24)|(data[5]<<16)|(data[6]<<8)|data[7], (data[4]<<24)|(data[5]<<16)|(data[6]<<8)|data[7])
	fmt.Printf("  [08-09] bodyLen: 0x%04X = %d\n", (data[8]<<8)|data[9], (data[8]<<8)|data[9])
	fmt.Printf("  [10-11] cmd: 0x%04X = %d\n", (data[10]<<8)|data[11], (data[10]<<8)|data[11])
	fmt.Println()

	offset := 12
	fmt.Println("Body fields:")
	fmt.Printf("  [%03d-%03d] auth_key: 0x%08X = %d\n", offset, offset+3, bytesToInt(data, offset), bytesToInt(data, offset))
	offset += 4
	fmt.Printf("  [%03d-%03d] result: 0x%08X = %d\n", offset, offset+3, bytesToInt(data, offset), bytesToInt(data, offset))
	offset += 4
	fmt.Printf("  [%03d-%03d] privilege: 0x%04X = %d\n", offset, offset+1, bytesToShort(data, offset), bytesToShort(data, offset))
	offset += 2

	strLen := int(data[offset]) & 0xFF
	fmt.Printf("  [%03d] ip len: %d\n", offset, strLen)
	offset += 1
	fmt.Printf("  [%03d-%03d] ip: %s\n", offset, offset+strLen-1, string(data[offset:offset+strLen]))
	offset += strLen

	fmt.Printf("  [%03d-%03d] port: 0x%04X = %d\n", offset, offset+1, bytesToShort(data, offset), bytesToShort(data, offset))
	offset += 2
	fmt.Printf("  [%03d-%03d] seed: 0x%08X = %d\n", offset, offset+3, bytesToInt(data, offset), bytesToInt(data, offset))
	offset += 4
	fmt.Printf("  [%03d-%03d] id: 0x%04X = %d\n", offset, offset+1, bytesToShort(data, offset), bytesToShort(data, offset))
	offset += 2

	strLen = int(data[offset]) & 0xFF
	fmt.Printf("  [%03d] serverName len: %d\n", offset, strLen)
	offset += 1
	fmt.Printf("  [%03d-%03d] serverName: %s\n", offset, offset+strLen-1, string(data[offset:offset+strLen]))
	offset += strLen

	fmt.Printf("  [%03d] serverStatus: 0x%02X = %d\n", offset, data[offset]&0xFF, data[offset]&0xFF)
	offset += 1

	strLen = int(data[offset]) & 0xFF
	fmt.Printf("  [%03d] msg len: %d\n", offset, strLen)
	offset += 1
	fmt.Printf("  [%03d-%03d] msg: %s\n", offset, offset+strLen-1, string(data[offset:offset+strLen]))
	offset += strLen

	fmt.Println()
	fmt.Printf("Total parsed: %d bytes\n", offset)
}

func bytesToInt(bytes []byte, offset int) int {
	return int(((bytes[offset] & 0xFF) << 24) |
		((bytes[offset+1] & 0xFF) << 16) |
		((bytes[offset+2] & 0xFF) << 8) |
		(bytes[offset+3] & 0xFF))
}

func bytesToShort(bytes []byte, offset int) int {
	return int(((bytes[offset] & 0xFF) << 8) | (bytes[offset+1] & 0xFF))
}
