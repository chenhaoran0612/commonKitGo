package global

const ETH_MAIN_POOL_NAME = "Ehashpool"
const ETEST_MAIN_POOL_NAME = "Etestpool"
const JAX_POOL_NAME = "JAXPool"
const ETEST_POOL_ID = 1
const JAX_POOL_ID = 2

// 子池ID to 子池名称
var PoolIdToName = map[int64]string{
	MAIN_POOL_ID:  ETH_MAIN_POOL_NAME,
	ETEST_POOL_ID: ETEST_MAIN_POOL_NAME,
	JAX_POOL_ID:   JAX_POOL_NAME,
}

func PoolId2Name(poolId int64) string {
	poolName, ok := PoolIdToName[poolId]
	if ok {
		return poolName
	}

	return ""
}

// 子池名称 to 子池ID
var PoolNameToId = map[string]int64{
	ETH_MAIN_POOL_NAME:   MAIN_POOL_ID,
	ETEST_MAIN_POOL_NAME: ETEST_POOL_ID,
	JAX_POOL_NAME:        JAX_POOL_ID,
}

func PoolName2Id(poolName string) int64 {
	poolId, ok := PoolNameToId[poolName]
	if ok {
		return poolId
	}

	return -1
}
