package tools

import (
	"github.com/holiman/uint256"
)

func AddUInt(x uint64, y uint64) *uint256.Int{
	xU := uint256.NewInt(x)
	return xU.AddUint64(xU, y)
}

func MulUInt(x uint64, y uint64) *uint256.Int{
	xU := uint256.NewInt(x)
	yU := uint256.NewInt(y)
	return xU.Mul(xU, yU)
}

func IsUInt(x *uint256.Int) bool {
	if x.IsUint64(){
		return true
	}

	return false
}

func ToUInt256Hex(x *uint256.Int) string{
	return x.Hex()
}

func ToUInt256String(x *uint256.Int) string{
	return ToUInt256Hex(x)
}

func ToUInt256ToUInt(x *uint256.Int) (uint64, bool){
	if IsUInt(x)!=true{
		return 0, true
	}
	out := x.Uint64()
	return out,false
}