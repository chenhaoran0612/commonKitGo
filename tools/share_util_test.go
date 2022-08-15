package tools

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"zeus/app/kit/global"
)

func TestBit2Diff(t *testing.T) {
	Bit2Diff(386547904)
}

// assert 文档： https://github.com/stretchr/testify
func TestShareToHashRate(t *testing.T) {

	share65535 := ShareToHashRate(65535, 60)
	t.Log(share65535)
	assert.Equal(t, share65535, 4691178029056.00000, "true")

	// 此数据有问题
	share65535max := ShareToHashRate(9223372036854775807, 60)
	t.Log(share65535max)
	assert.NotEqual(t, share65535max, 660234687618869479874616797.866666666666, "false")

}

func TestShareToHashrate2(t *testing.T) {
	hashrate := ShareToHashRate(29570168636357, global.MINUTE10_SECONDS)
	t.Log(FormatHashrate2(hashrate)) // result : 211.671 EH/s
}

func TestEthDiffToHashrate(t *testing.T) {
	share65535 := ShareToHashRate(11830924580848856, 60)
	t.Log(share65535)
}

func TestHashrateToShare(t *testing.T) {
	hashrate := 172025 * math.Pow10(12)
	share := HashrateToShare(decimal.NewFromFloat(hashrate), global.DAY_SECONDS)

	t.Log(share)
	hashrate2 := ShareToHashRate(3420319454817/24, global.HOUR_SECONDS)
	t.Log(hashrate2)
}

func TestBtcShareToEarn2(t *testing.T) {
	earn := BtcShareToEarn(65535, 13672594272814)
	t.Log(earn)
	assert.Equal(t, earn, 2.9957281100223874)

	earn2 := BtcShareToEarn(65535, 9223372036854775807)
	t.Log(earn2)
	assert.Equal(t, earn2, 4.4408243359e-06)
}

func TestMaxValue(t *testing.T) {
	// integer max
	fmt.Printf("max int64 = %+v\n", math.MaxInt64)
	fmt.Printf("max int32 = %+v\n", math.MaxInt32)
	fmt.Printf("max int16 = %+v\n", math.MaxInt16)

	// integer min
	fmt.Printf("min int64 = %+v\n", math.MinInt64)
	fmt.Printf("min int32 = %+v\n", math.MinInt32)

	fmt.Printf("max flloat64= %+v\n", math.MaxFloat64)
	fmt.Printf("max float32= %+v\n", math.MaxFloat32)
}

func TestFormatPercent(t *testing.T) {

	fmt.Println(FormatPercent("0.1", 2))
	fmt.Println(FormatPercent("0.11224", 2))
	fmt.Println(FormatPercent("0.11225", 2))
	fmt.Println(FormatPercent("0.11225", 0))
	fmt.Println(FormatPercent("0.11225", -1))

}

func TestFormatBTC(t *testing.T) {

	fmt.Println(FormatBTCCoin(decimal.NewFromInt(1), 8))
	fmt.Println(FormatBTCCoin(decimal.New(1, 8), 8))
	fmt.Println(FormatBTCCoin(decimal.New(-1, 8), 8))
	fmt.Println(FormatBTCCoin(decimal.New(-0, 8), 8))
	v, _ := decimal.NewFromString("")
	fmt.Println(FormatBTCCoin(v, 8))

}
