package tools

import "testing"

func TestFormatEthEarn(t *testing.T) {
	t.Log(FormatEthEarn(122222222121221, "5"))
	t.Log(FormatEthEarn(122222222121221, "15"))
	t.Log(FormatEthEarn(122222222121221, "18"))
}

func TestFormatEthEarnString(t *testing.T) {
	t.Log(FormatEthEarnFromString("122222222121221", 8)) // result: 0.00001222
	t.Log(FormatEthEarnFromString("122252222121221", 8)) // result: 0.00001222
}

func TestFormatBtcEarnString(t *testing.T) {
	t.Log(FormatEarnFromString("100000000", 8, "BTC"))       // result: 1.00000000
	t.Log(FormatEarnFromString("100000000", 8, "BCH"))       // result: 1.00000000
	t.Log(FormatEarnFromString("100000000", 8, "BSV"))       // result: 1.00000000
	t.Log(FormatEarnFromString("1000000", 8, "BTC"))         // result: 0.01000000
	t.Log(FormatEarnFromString("122252222121221", 8, "ETH")) // result: 0.00012225
}

func TestDiff2PerTEarn(t *testing.T) {
	r := Diff2PerTEarn(27452707)
	t.Log(r)
}
