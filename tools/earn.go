package tools

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func FormatEthEarn(earn float64, keepBit string) string {
	if earn == 0 {
		return "0"
	}
	return fmt.Sprintf("%."+keepBit+"f", earn/math.Pow10(18))
}

var POW18 decimal.Decimal = decimal.New(1, 18)
var POW8 decimal.Decimal = decimal.New(1, 8)
var POW32 decimal.Decimal = decimal.NewFromFloat(math.Pow(2, 32))

func FormatEthEarnFromString(earn string, keepBit int32) string {
	if earn == "" {
		return "0"
	}

	earnDecimal, err := decimal.NewFromString(earn)
	if err != nil {
		return "0"
	}
	return FormatEthEarnFromDecimal(earnDecimal, keepBit)
}

func FormatEarnFromString(earn string, keepBit int32, coinType string) string {
	if earn == "" {
		return "0"
	}

	earnDecimal, err := decimal.NewFromString(earn)
	if err != nil {
		return "0"
	}
	switch coinType {
	case "BTC", "BSV", "BCH":
		return FormatBtcEarnFromDecimal(earnDecimal, keepBit)
	case "ETH":
		return FormatEthEarnFromDecimal(earnDecimal, keepBit)
	default:
		return "FormatEarnFromStringError"
	}
	return FormatEthEarnFromDecimal(earnDecimal, keepBit)
}

func Diff2PerTEarn(diff int64) string {
	if diff == 0 {
		return "0"
	}
	return decimal.NewFromInt(86400).Mul(decimal.NewFromInt(1e12)).Div(decimal.NewFromInt(diff)).Div(POW32).Mul(decimal.NewFromFloat(BtcBlockEarnBitCoin())).String()
}

func FormatEthEarnFromDecimal(earn decimal.Decimal, keepBit int32) string {
	return earn.Div(POW18).Truncate(keepBit).StringFixedBank(keepBit)
}

func FormatBtcEarnFromDecimal(earn decimal.Decimal, keepBit int32) string {
	return earn.Div(POW8).Truncate(keepBit).StringFixedBank(keepBit)
}
