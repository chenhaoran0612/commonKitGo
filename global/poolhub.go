package global

import (
	"errors"
	"github.com/shopspring/decimal"
)

const STAT_DAY_HASHRATE_BY_SHATE = "STAT_DAY_HASHRATE_BY_SHATE" // 通过share统计天算力

const STAT_REPORT = "STAT_REPORT"             // 统计报告事件
const STAT_REPORT_MERGE = "STAT_REPORT_MERGE" // 统计报告合并事件

// 统计报表状态
const (
	DATA_HANDLE_STAT_REPORT_TYPE        = "DATA_HANDLE"        // 数据处理中
	DATA_HANDLE_FILUED_STAT_REPORT_TYPE = "DATA_HANDLE_FAILED" // 数据处理失败
	GENERATE_STAT_REPORT_TYPE           = "GENERATE"           // 已生成
	EXPORT_STAT_REPORT_TYPE             = "EXPORT"             // 已导出
)

const (
	BTC = "BTC"
	BCH = "BCH"
	BSV = "BSV"
	ETH = "ETH"
)

const (
	UNIT_SATOSHI = 100000000           // 单位为 聪
	UNIT_ETH_WEI = 1000000000000000000 // 单位为 ETH 的 Wei
)

const MINUTE2_SECONDS = 2 * 60   // 2  分钟秒数
const MINUTE10_SECONDS = 10 * 60 // 10  分钟秒数
const MINUTE15_SECONDS = 15 * 60 // 15  分钟秒数
const HOUR_SECONDS = 60 * 60
const DAY_SECONDS = 24 * HOUR_SECONDS
const DAY_MILLSECONDS = DAY_SECONDS * 1000

// 一年毫秒数
const YEAR_MILLSECONDS = 365 * 24 * 60 * 60 * 1000

var BTC0_001 decimal.Decimal = decimal.NewFromInt(1e5) // 0.001 BTC
var ETH0_1 decimal.Decimal = decimal.NewFromInt(1e17)  // 0.1 ETH
var ZERO decimal.Decimal = decimal.NewFromInt(0)       // 0
var ONE decimal.Decimal = decimal.NewFromInt(1)        // 1

// s3 下载文件，最大重试次数
const S3_DOWNLOAD_FILE_MAX_RETRY_COUNT = 3

// 突发事件原因类型
const (
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_POWER                            = "POWER"                        // 电源故障
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_TECHNICAL                        = "TECHNICAL"                    // 技术故障 Technical failure
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_NETWORK                          = "NETWORK"                      // 网络故障 Network malfunction
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_CLIMATE                          = "CLIMATE"                      // 气候原因 Climate
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_UNKNOWN                          = "UNKNOWN"                      // 无知的原因 Unknown reason
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_LOWER_HASHRATE                   = "LOWER_HASHRATE"               // 低算力 Lower hashrate
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_LOWER_MACHINE_COUNT              = "LOWER_MACHINE_COUNT"          // 低机器数量 Lower machine count
	STAT_REPORT_ABNORMAL_EVENT_REASON_TYPE_LOWER_HASHRATE_AND_MACHINE_COUNT = "LOWER_HASHRATE_MACHINE_COUNT" // 低算力和机器数量 Lower hashrate and machine count
)

var STAT_REPORT_NOT_EXIST = errors.New("STAT_REPORT_NOT_EXIST")
var STAT_REPORT_NOT_MULTI_LEVEL_MERGE = errors.New("STAT_REPORT_NOT_MULTI_LEVEL_MERGE")

var MINER_IS_EMPTY = errors.New("MINER IS EMPTY")
var MINER_NOT_EXISTS = errors.New("MINER NOT EXISTS")
var MINER_FORMAT_NOT_SUPPORT = errors.New("MINER_FORMAT_NOT_SUPPORT")

var MANUAL_BILL_TYPE_INCOMING = "INCOMING" // 手工录入的账单类型，收入

var MANUAL_BILL_TYPE_EXTRA_PAYOUT = "EXTRA_PAYOUT" // 手工录入的账单类型，额外支出

var STAT_REPORT_ABNORMAL_EVENT_INFO_ERROR = errors.New("STAT_REPORT_ABNORMAL_EVENT_INFO_ERROR")

var USER_NOT_EXISTS = errors.New("USER NOT EXISTS")
var NOT_SUPPORT = errors.New("NOT SUPPORT")
var NOT_FIND_MINER = errors.New("NOT_FIND_MINER")

const (
	HASHRATE_SPAN_TYPE_HOUR     = 0 // 算力 span 小时 类型
	HASHRATE_SPAN_TYPE_DAY      = 1 // 算力 span 天 类型
	HASHRATE_SPAN_TYPE_10MINUTE = 2 // 算力 span 10分钟 类型
)

var (
	HASHRATE_SPAN_TYPE_HOUR_STR     = "1h"  // 算力 span 小时 类型 字符串表示
	HASHRATE_SPAN_TYPE_DAY_STR      = "1d"  // 算力 span 天 类型  字符串表示
	HASHRATE_SPAN_TYPE_10MINUTE_STR = "10m" // 算力 span 10分钟 类型  字符串表示
)

var (
	HASHRATE_SPAN_KEEP_MINER_ID_CLOUMN     = "minerId"     // 算力 span 保留minerId列数据，前提是minerId列有过滤数据
	HASHRATE_SPAN_KEEP_WORKER_ID_CLOUMN    = "workerId"    // 算力 span 保留workerId列数据，前提是workerId列有过滤数据
	HASHRATE_SPAN_KEEP_EARN_CLOUMN         = "earn"        // 算力 span 保留earn列数据，前提是有earn列
	HASHRATE_SPAN_KEEP_REVISED_EARN_CLOUMN = "revisedEarn" // 算力 span 保留revisedEarn列数据，前提是有revisedEarn列
)

var EVENT_EXCEED_MAX_LOADING = errors.New("Exceed the maximum load") // 正在运行的事件超过了最大负载

var ADDRESS_NOT_SUPPORT = errors.New("ADDRESS NOT SUPPORT")

const (
	MANUAL_EARN_RECORD_INCOMING_TYPE = "INCOMING" // 收入
	MANUAL_EARN_RECORD_PAYOUT_TYPE   = "PAYOUT"   // 支出
)

const (
	MAIN_POOL_ID = 0 // 主池ID为 0
)
