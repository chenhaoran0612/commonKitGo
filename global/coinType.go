package global

var CoinTypeMap = map[string]*CoinTypeInfo{
	"BTC": {
		Name: "BTC",
	},
	"ETH": {
		Name: "ETH",
	},
}

type CoinTypeInfo struct {
	Name string
}

func CheckCoinType(coinTypeName string) bool {
	_, ok := CoinTypeMap[coinTypeName]
	return ok
}
