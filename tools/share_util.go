package tools

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"zeus/app/kit/global"
)

var ShareToHashrateScales decimal.Decimal = decimal.NewFromInt(4294967296)
var bit2e8 = decimal.NewFromInt(256)

// Bit2Diff 测试未通过
func Bit2Diff(bit int64) int64 {
	nShift := (bit >> 24) & 0xff
	d := bit & 0x00ffffff
	dDiff := 0x0000ffff / (d)
	for nShift < 29 {
		dDiff = dDiff * 256
		nShift++
	}
	for nShift > 29 {
		//bcdiv(dDiff, 256)
		dDiff = decimal.NewFromInt(dDiff).Div(bit2e8).Truncate(0).BigInt().Int64()
		nShift--
	}
	//dDiff = bcmul(dDiff, 1, 0)
	dDiff = decimal.NewFromInt(dDiff).Div(decimal.NewFromInt(1)).Truncate(0).BigInt().Int64()
	return dDiff
}

func ShareToHashRate(share int64, duration int) decimal.Decimal {
	if duration > 0 {
		tmp := decimal.NewFromInt(share).Mul(ShareToHashrateScales)
		re := tmp.Div(decimal.NewFromInt(int64(duration)))
		return re
	}

	return decimal.NewFromInt(0)
}

// 算力 转换为 share
func HashrateToShare(hashrate decimal.Decimal, duration int) int64 {
	if duration > 0 {
		tmp := hashrate.Mul(decimal.NewFromInt(int64(duration)))
		re := tmp.Div(ShareToHashrateScales)
		bigIntV := re.Truncate(0).BigInt()
		reInt := bigIntV.Int64()
		return reInt
	}

	return 0
}

func BtcBlockEarn() int64 {
	// 6.25*pow(10,8)
	return 625000000
}

func BtcBlockEarnBitCoin() float64 {
	return 6.25
}

func BtcShareToEarn(share int64, diff int64) decimal.Decimal {
	return decimal.NewFromInt(share).Mul(decimal.NewFromInt(BtcBlockEarn())).Mul(decimal.NewFromInt(diff))
}

var EH float64 = float64(math.Pow10(18))

func EHashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(EH))
	re, _ := tmp.Float64()
	return re
}

var PH float64 = math.Pow10(15)

func PHashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(PH))
	re, _ := tmp.Float64()
	return re
}

var TH float64 = math.Pow10(12)

func THashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(TH))
	re, _ := tmp.Float64()
	return re
}

var GH float64 = math.Pow10(9)

func GHashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(GH))
	re, _ := tmp.Float64()
	return re
}

var MH float64 = math.Pow10(6)

func MHashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(MH))
	re, _ := tmp.Float64()
	return re
}

var KH float64 = math.Pow10(3)

func KHashrate(hashrate float64) float64 {
	tmp := decimal.NewFromFloat(hashrate).Div(decimal.NewFromFloat(KH))
	re, _ := tmp.Float64()
	return re
}

func FormatHashrate(hashrate float64) string {
	_tmp := fmt.Sprintf("%.0f", hashrate)
	var re float64 = 0
	unit := "H/s"
	if len(_tmp) > 18 {
		re = EHashrate(hashrate)
		unit = "EH/s"
	} else if len(_tmp) > 15 {
		re = PHashrate(hashrate)
		unit = "PH/s"
	} else if len(_tmp) > 12 {
		re = THashrate(hashrate)
		unit = "TH/s"
	} else if len(_tmp) > 9 {
		re = GHashrate(hashrate)
		unit = "GH/s"
	} else if len(_tmp) > 6 {
		re = MHashrate(hashrate)
		unit = "MH/s"
	} else if len(_tmp) > 3 {
		re = KHashrate(hashrate)
		unit = "KH/s"
	}

	return fmt.Sprintf("%.3f %s", re, unit)
}

func FormatTHashrate(hashrate float64) string {
	var re float64 = 0
	unit := "TH/s"
	re = THashrate(hashrate)

	return fmt.Sprintf("%.2f %s", re, unit)
}

func EHashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(EH)).Truncate(3)
	return tmp.String()
}

func PHashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(PH)).Truncate(3)
	return tmp.String()
}

func THashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(TH)).Truncate(3)
	return tmp.String()
}

func GHashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(GH)).Truncate(3)
	return tmp.String()
}

func MHashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(MH)).Truncate(3)
	return tmp.String()
}
func KHashrateDecimal(hashrate decimal.Decimal) string {
	tmp := hashrate.Div(decimal.NewFromFloat(KH)).Truncate(3)
	return tmp.String()
}

func FormatHashrate2(hashrate decimal.Decimal) string {
	_tmp := fmt.Sprintf("%s", hashrate.Truncate(0).String())
	var re string = ""
	unit := "H/s"
	if len(_tmp) > 18 {
		re = EHashrateDecimal(hashrate)
		unit = "EH/s"
	} else if len(_tmp) > 15 {
		re = PHashrateDecimal(hashrate)
		unit = "PH/s"
	} else if len(_tmp) > 12 {
		re = THashrateDecimal(hashrate)
		unit = "TH/s"
	} else if len(_tmp) > 9 {
		re = GHashrateDecimal(hashrate)
		unit = "GH/s"
	} else if len(_tmp) > 6 {
		re = MHashrateDecimal(hashrate)
		unit = "MH/s"
	} else if len(_tmp) > 3 {
		re = KHashrateDecimal(hashrate)
		unit = "KH/s"
	}

	return fmt.Sprintf("%s %s", re, unit)
}

func FormatHashrate3(hashrate string) string {
	if hashrate == "" {
		return "0 H/s"
	}
	_hashrate, _ := decimal.NewFromString(hashrate)
	return FormatHashrate2(_hashrate)
}

func FormatPercent(percent string, keepBit int32) string {
	if percent == "" {
		return "0%"
	}
	if keepBit < 0 {
		keepBit = 0
	}
	format := "%." + strconv.Itoa(int(keepBit)) + "f%%"
	_hashrate, _ := decimal.NewFromString(percent)
	v, _ := _hashrate.Mul(decimal.NewFromInt(100)).Truncate(keepBit).Float64()
	return fmt.Sprintf(format, v)
}

func FormatBTCCoin(coinValue decimal.Decimal, keepBit int32) string {
	return FormatCoin(coinValue, decimal.New(1, 8), global.BTC, keepBit)
}

func FormatCoin(coinValue decimal.Decimal, coinMinUnit decimal.Decimal, coinName string, keepBit int32) string {
	if coinMinUnit.Equal(global.ZERO) {
		return "0 " + coinName
	}
	value, _ := coinValue.Div(coinMinUnit).Truncate(keepBit).Float64()
	format := "%." + strconv.Itoa(int(keepBit)) + "f " + coinName
	return fmt.Sprintf(format, value)
}

func Div(data0 int64, data1 float64) int64 {
	if data1 == 0 {
		return 0
	}
	tmp := decimal.New(data0, 0).Div(decimal.NewFromFloat(data1))
	re := tmp.IntPart()
	return re
}
