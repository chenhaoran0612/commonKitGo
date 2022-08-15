package tools

import "github.com/shopspring/decimal"

var BTC_BLOCK_ONE_LAYER_NUMS decimal.Decimal = decimal.NewFromInt(210000)
var BTC_FIRST_BLOCK_REWRAD decimal.Decimal = decimal.NewFromInt(5000000000)
var TWO decimal.Decimal = decimal.NewFromInt(2)
var PER_DAY_BTC_BLOCK_COUNT decimal.Decimal = decimal.NewFromInt(144)

// see: https://en.bitcoin.it/wiki/Controlled_supply
func GetBTCBlockReward(height int64) decimal.Decimal {
	if height < 0 || height >= 6930000 {
		return decimal.Zero
	}

	blockHeight := decimal.NewFromInt(height)
	pows := blockHeight.Div(BTC_BLOCK_ONE_LAYER_NUMS)
	return BTC_FIRST_BLOCK_REWRAD.Div(TWO.Pow(pows)).Truncate(0)
}
