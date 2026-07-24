package constant

import "testing"

func TestClientButtonIdConst(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"CharRename", CharRename, 1},
		{"DropTask", DropTask, 2},
		{"GetRankInfo", GetRankInfo, 3},
		{"NotifySendInitDataDone", NotifySendInitDataDone, 39},
		{"NotifyTestYbxsEndTime", NotifyTestYbxsEndTime, 60001},
		{"NotifyMailAllLoaded", NotifyMailAllLoaded, 10012},
		{"NotifyIOSReview", NotifyIOSReview, 50017},
		{"NotifyQmoxiangStatus", NotifyQmoxiangStatus, 20010},
		{"NotifyAutoDisconnect", NotifyAutoDisconnect, 20011},
		{"NotifyOpenExorcism", NotifyOpenExorcism, 20008},
		{"NotifyCloseExorcism", NotifyCloseExorcism, 20009},
		{"CombatGetCurrentRound", CombatGetCurrentRound, 30041},
		{"NotifyFetchLivenessBonus", NotifyFetchLivenessBonus, 33},
		{"GetLivenessInfo", NotifyGetLivenessInfo, 32},
		// 更多常量可按需要添加测试
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("constant %s = %d, expected %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}
