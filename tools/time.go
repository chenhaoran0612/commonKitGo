package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var HOUR_REG, _ = regexp.Compile("\\d{10}")

var MINUTE_FORMAT = "200601021504"
var HOUR_FORMAT = "2006010215"
var DAY_FORMAT = "20060102"
var DAY_FORMAT2 = "2006-01-02"
var HOUR_FORMAT2 = "2006-01-02 15H"
var SECONDS_FORMAT = "2006-01-02 15:04:05"

// loggerName as sharelogger.2021072012.csv.zst; result = 2021072012
func ParseHourFromLoggerName(loggerName string) string {
	return HOUR_REG.FindString(loggerName)
}

var BEIBING_TZ = GetLocationByHour(8) // 北京 时区
var UTC0_TZ = GetLocationByHour(0)    // utc+0 时区

// 支持的时区列表
var SUPPORT_TIMEZONE_MAP map[int]*time.Location = map[int]*time.Location{
	8: BEIBING_TZ,
	0: UTC0_TZ,
}

// 解析北京时间为unix时间戳
// 例子：timeStr = 2021-07-20 18:50:30 ; timeFormat = 2006-01-02 15:04:05, result = 1626778230
func ParseBeijingTimeToUnix(timeStr string, timeFormat string) int64 {
	return ParseTimeToUnix(timeStr, timeFormat, BEIBING_TZ)
}

// 解析UTC+0 时间为unix时间戳
// 例子：timeStr = 2021-07-20 18:50:30 ; timeFormat = 2006-01-02 15:04:05, result = 1626778230
func ParseUTCTimeToUnix(timeStr string, timeFormat string) int64 {
	return ParseTimeToUnix(timeStr, timeFormat, UTC0_TZ)
}

func ParseTimeToUnix(timeStr string, timeFormat string, timezone *time.Location) int64 {
	time2, _ := time.ParseInLocation(timeFormat, timeStr, timezone)
	return time2.Unix()
}

// 格式化为utc天 ，例子：unixTime = 1626778230 ；result = 20210720
func FormatUtcDay(unixTime int64) (int, error) {
	return FormatDay(unixTime, UTC0_TZ)
}

// 格式化为北京时区天 ，例子：unixTime = 1626778230 ；result = 20210720
func FormatBeijingDay(unixTime int64) (int, error) {
	return FormatDay(unixTime, BEIBING_TZ)
}

// 格式化为utc小时 ，例子：unixTime = 1626778230 ；result = 2021072010
func FormatUtcHour(unixTime int64) (int64, error) {
	return FormatHour(unixTime, UTC0_TZ)
}

// 格式化为北京时区小时 ，例子：unixTime = 1626778230 ；result = 2021072018
func FormatBeiJingHour(unixTime int64) (int64, error) {
	return FormatHour(unixTime, BEIBING_TZ)
}

// 格式化 时区小时 ，例子：unixTime = 1626778230, 北京时区 ；result = 2021072018
func FormatHour(unixTime int64, timezone *time.Location) (int64, error) {
	return FormatUnixToInt64(unixTime, timezone, HOUR_FORMAT)
}

// 格式化 时区小时 ，例子：unixTime = 1626778230, 北京时区 ；result = 20210720
func FormatDay(unixTime int64, timezone *time.Location) (int, error) {
	tmp := time.Unix(unixTime, 0)
	re := tmp.In(timezone).Format(DAY_FORMAT)
	reInt, err := strconv.Atoi(re)
	return reInt, err
}

// 格式化 unix 时间 到 int64 ，例子：unixTime = 1626778230, 北京时区 ；result = 20210720
func FormatUnixToInt64(unixTime int64, timezone *time.Location, format string) (int64, error) {
	tmp := time.Unix(unixTime, 0)
	re := tmp.In(timezone).Format(format)
	reInt, err := strconv.ParseInt(re, 10, 64)
	return reInt, err
}

// 格式化 unix 时间 到 int64 ，例子：unixTime = 1626778230, 北京时区 format=2006-01-02 15:04:05；result = 20210720
func FormatUnixToString(unixTime int64, timezone *time.Location, format string) string {
	tmp := time.Unix(unixTime, 0)
	re := tmp.In(timezone).Format(format)
	return re
}

func GetNowTimeMillseconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetNowTimeSeconds() int64 {
	return time.Now().Unix()
}

func GetLocationByHour(hour int) *time.Location {
	locationName := "UTC+%d"
	if hour < 0 {
		locationName = "UTC%d"
	}
	return time.FixedZone(fmt.Sprintf(locationName, hour), hour*60*60)
}

//func GetNowTimeSpan(seconds int64) int64{
//	now := GetNowTimeSeconds()
//	modSeconds := now%seconds
//	return now - modSeconds
//}
