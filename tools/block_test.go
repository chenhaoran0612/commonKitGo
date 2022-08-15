package tools

import "testing"

func TestGetBtcBlockReward(t *testing.T) {
	height := 555394
	t.Log(GetBTCBlockReward(int64(height))) // result : 1250000000

	height = 746811
	t.Log(GetBTCBlockReward(int64(height))) // result : 625000000

}
