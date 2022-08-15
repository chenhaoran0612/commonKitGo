package open_platform

type PlatFormErrorEnum int32

//系统错误码
const (
	SYSETM_TOKEN_NOT_EMPTY PlatFormErrorEnum = 10001
	SYSETM_TOKEN_EXPIRE PlatFormErrorEnum = 10002
	SYSETM_TOKEN_INVALID PlatFormErrorEnum = 10003
)

//通用错误码
const (
	SORT_COLUMN_NOT_EXIST PlatFormErrorEnum = 20001
	SORT_TYPE_NOT_EXIST PlatFormErrorEnum = 20002
	SORT_COLUMN_NOT_EMPTY PlatFormErrorEnum = 20003
	SORT_TYPE_NOT_EMPTY PlatFormErrorEnum = 20004

	DAY_FORMAT_NOT_SUPPORT PlatFormErrorEnum = 20005
	DAY_RANGE_NOT_SUPPORT PlatFormErrorEnum = 20006
)

//Rawpool Account 微服务错误码
const (
	ACCOUNT_UID_TYPE_ERROR PlatFormErrorEnum = 30001
	ACCOUNT_POOL_ID_NOT_EMPTY PlatFormErrorEnum = 30002
	ACCOUNT_EMAIL_TYPE_ERROR PlatFormErrorEnum = 30003
	ACCOUNT_NOT_EXIST  PlatFormErrorEnum = 30004
)

//Rawpool Poolhub 微服务错误码
const (
	MINER_NOT_EXIST PlatFormErrorEnum = 40001
)

//GATEWAY 错误码
const (
	HASH_RATE_CHART_TYPE_NOT_EXISTS PlatFormErrorEnum = 50001
	HASH_RATE_CHART_VALUE_NOT_NUMBER PlatFormErrorEnum = 50002
)

//错误提示消息
func (p PlatFormErrorEnum) String() string {
	switch p {
	case SYSETM_TOKEN_NOT_EMPTY: return "token not empty"
	case SYSETM_TOKEN_EXPIRE: return "token expire"
	case SYSETM_TOKEN_INVALID: return "token invalid"
	case SORT_COLUMN_NOT_EXIST: return "sort column does not exist"
	case SORT_TYPE_NOT_EXIST: return "sort type does not exist"
	case SORT_COLUMN_NOT_EMPTY: return "sort column does not empty"
	case SORT_TYPE_NOT_EMPTY: return "sort type does not empty"
	case DAY_FORMAT_NOT_SUPPORT: return "day format not support"
	case DAY_RANGE_NOT_SUPPORT: return "day range not support"
	case ACCOUNT_UID_TYPE_ERROR: return "account uid type error"
	case ACCOUNT_POOL_ID_NOT_EMPTY: return "account pool id not empty"
	case ACCOUNT_EMAIL_TYPE_ERROR: return "email type error"
	case ACCOUNT_NOT_EXIST: return "account not exist"
	case MINER_NOT_EXIST: return "miner not exist"
	case HASH_RATE_CHART_TYPE_NOT_EXISTS: return "search type not exist"
	case HASH_RATE_CHART_VALUE_NOT_NUMBER: return "search value not number"
	default:         return "UNKNOWN"
	}
}