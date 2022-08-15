package tools

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func UInt64ToDecimal(v uint64) decimal.Decimal{
	bInt := &big.Int{}
	bInt.SetUint64(v)
	return decimal.NewFromBigInt(bInt, 0)
}
