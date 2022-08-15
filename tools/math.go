package tools

import (
	"github.com/shopspring/decimal"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"zeus/app/kit/global"
)

// 变化率 = (d1-d2)/d2
func ChangedRatio64(d1 float64, d2 float64) float64{
	if d2>0{
		tmp := d1 - d2
		return tmp /d2
	}

	return 0
}

// 变化率 = (d1-d2)/d2
func ChangedRatio32(d1 float32, d2 float32) float32{
	if d2>0{
		tmp := d1 - d2
		return tmp /d2
	}

	return 0
}

// 变化率 = (d1-d2)/d2
func ChangedRatioDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal{
	if d2.Cmp(global.ZERO)>0{
		tmp := d1.Sub(d2)
		return tmp.Div(d2)
	}

	return decimal.NewFromInt(0)
}


// 占比 = d1/d2
func RatioDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal{
	if d2.Cmp(global.ZERO)>0{
		return d1.Div(d2)
	}

	return decimal.NewFromInt(0)
}

// Decimal2Float64 decimal 转换为 float64 return float64
// v 待转换数据
// keepBit 保留的小数点位数
// isRound 是否四舍五入
func Decimal2Float64(v decimal.Decimal, keepBit int32, isRound bool ) float64{
	vStr := "0"
	if isRound{
		vStr = v.Round(keepBit).String()
	}else{
		vStr = v.Truncate(keepBit).String()
	}
	result, _ := strconv.ParseFloat(vStr, 64)
	return result
}

// Decimal2Float32 decimal 转换为 float32 return float32
// v 待转换数据
// keepBit 保留的小数点位数
// isRound 是否四舍五入
func Decimal2Float32(v decimal.Decimal, keepBit int32, isRound bool) float32{
	vStr := "0"
	if isRound{
		vStr = v.Round(keepBit).String()
	}else{
		vStr = v.Truncate(keepBit).String()
	}
	result, _ := strconv.ParseFloat(vStr, 64)
	return float32(result)
}

// 占比 = d1/d2
func Ratio64(d1 float64, d2 float64) float64{
	if d2>0{
		return d1 /d2
	}

	return 0
}


// 占比 = d1/d2
func Ratio32(d1 float32, d2 float32) float32{
	if d2>0{
		return d1 /d2
	}

	return 0
}

// 保留浮点位数 = d1 * pow(10,keepFloatCount) / pow(10,keepFloatCount)
func KeepFloat64(d1 float64, keepFloatCount int) float64{
    if d1 >0 && keepFloatCount>0{
		keepFloatPow := math.Pow10(keepFloatCount)
    	return d1 * keepFloatPow / keepFloatPow
	}
	return 0
}

// 保留浮点位数 = d1 * pow(10,keepFloatCount) / pow(10,keepFloatCount)
func KeepFloat32(d1 float32, keepFloatCount int) float32{
    if d1 >0 && keepFloatCount>0{
		keepFloatPow := math.Pow10(keepFloatCount)
    	return float32(float64(d1) * keepFloatPow / keepFloatPow)
	}
	return 0
}

func Float64ToString(x float64) string{
	if x==0{
		return "0"
	}

	out := strconv.FormatFloat(x, 'f',10, 64)
	if out==""{
		return "0"
	}
	dotIndex := strings.Index(out, ".")

	return out[0:dotIndex]
}

func StringToFloat64(x string) float64{
	out, _ := strconv.ParseFloat(x, 64)
	return out
}


func Float64ToUint64(x float64) (uint64, error){
	s := Float64ToString(x)
	return strconv.ParseUint(s, 10, 64)
}


func RandomRange(min int64, max int64) int64{
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max - min + 1) + min
}