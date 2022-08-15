package tools

import (
	"bytes"
	"crypto/md5"
	r "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"hash/fnv"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
)

var snowNode *snowflake.Node

func init() {
	name, _ := os.Hostname()
	h := fnv.New32()
	h.Write([]byte(name))
	sID := h.Sum32()
	ID := sID % 1024
	snowNode, _ = snowflake.NewNode(int64(ID))
}

const (
	//Mobile ...
	Mobile = "mobile"
	//Web ...
	Web = "web"
)

const (
	//AlipaySuccess ...
	AlipaySuccess = "TRADE_SUCCESS"
	//WxSuccess ...
	WxSuccess = "SUCCESS"
	//Wechat..
	Wechat = "Wechat"
	//Alipay..
	Alipay = "Alipay"
)

//数据库 in ()使用
func ArrayToString(data []string) string {
	result := strings.Builder{}
	for _, v := range data {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write([]byte(fmt.Sprintf(" '%s'", v)))
	}
	return result.String()
}

func ArrayIntToString(data []int) string {
	result := strings.Builder{}
	for _, v := range data {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write([]byte(fmt.Sprintf(" '%d'", v)))
	}
	return result.String()
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

//雪花ID
func NextSnowflakeId() int64 {
	return int64(snowNode.Generate())
}

// StringInSlice 字符串是否在数组中
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//生成 Trade NO
func GenerateTradeNo(clientUserId int64) string {
	st := time.Now().Format("20060102150405")
	return st + "-" + fmt.Sprint(clientUserId)[10:19] + "-" + RandomDigit(4)
}

// 生成退款单号
func GenerateRefundNo() string {
	st := time.Now().Format("20060102150405")
	return "Refund-" + st + "-" + RandomDigit(4)
}

// 生成对应微信支付需要的时间格式
func GenerateWxTime(t time.Time) string {
	st := t.Format("20060102150405")
	return st
}

func RandomDigit(digit int) string {
	if digit < 1 || digit > 10 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	var target string
	switch digit {
	case 1:
		d := rand.Intn(10)
		target = fmt.Sprintf("%d", d)
	case 2:
		d := rand.Intn(100)
		target = fmt.Sprintf("%02d", d)
	case 3:
		d := rand.Intn(1000)
		target = fmt.Sprintf("%03d", d)
	case 4:
		d := rand.Intn(10000)
		target = fmt.Sprintf("%04d", d)
	case 5:
		d := rand.Intn(100000)
		target = fmt.Sprintf("%05d", d)
	case 6:
		d := rand.Intn(1000000)
		target = fmt.Sprintf("%06d", d)
	case 7:
		d := rand.Intn(10000000)
		target = fmt.Sprintf("%07d", d)
	case 8:
		d := rand.Intn(100000000)
		target = fmt.Sprintf("%08d", d)
	case 9:
		d := rand.Intn(1000000000)
		target = fmt.Sprintf("%09d", d)
	default:

	}
	return target
}

//MD5算法
func Md5(str string) (result string) {
	m := md5.New()
	m.Write([]byte(str))
	result = hex.EncodeToString(m.Sum(nil))
	return
}

//MD5算法
func Md5ByBytes(data []byte)  string {
	m := md5.New()
	m.Write(data)
	return  hex.EncodeToString(m.Sum(nil))
}

func Sha256(str string) (result string) {
	if len(str) == 0 {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(str))
	result = hex.EncodeToString(h.Sum(nil))
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func StringToFloat32(str string) float32{
	re, err := strconv.ParseFloat(str, 64)
	if err!=nil{return 0}
	return float32(re)
}

func GetSysPath() (dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

func GUID() string {
	b := make([]byte, 48)
	io.ReadFull(r.Reader, b)
	data := []byte(base64.URLEncoding.EncodeToString(b))
	has := md5.Sum(data)
	return fmt.Sprintf("%X", has)
}

func SourceToTargetByCopier(source interface{}, target interface{}) {
	copier.Copy(target, source)
}

func StrToGwei(s string) (d uint64 , err error) {
	if len(s) < 9 {
		return 0 , nil
	}
	d , err = strconv.ParseUint(s[0:len(s) - 9] , 10, 64)
	return
}


//使用反射，转换结构体
func SourceToTarget(sourceStruct interface{}, targetStruct interface{}) {
	source := structToMap(sourceStruct)
	targetV := reflect.ValueOf(targetStruct).Elem()
	targetT := reflect.TypeOf(targetStruct).Elem()
	for i := 0; i < targetV.NumField(); i++ {
		fieldName := targetT.Field(i).Name
		sourceVal := source[fieldName]
		if !sourceVal.IsValid() {
			continue
		}
		targetVal := targetV.Field(i)

		// 目标类型和源类型不一致，跳过赋值
		if targetVal.Type() != sourceVal.Type() {
			continue
		}
		targetVal.Set(sourceVal)
	}
}

func structToMap(structName interface{}) map[string]reflect.Value {
	t := reflect.TypeOf(structName).Elem()
	v := reflect.ValueOf(structName).Elem()
	fieldNum := t.NumField()
	resMap := make(map[string]reflect.Value, fieldNum)
	for i := 0; i < fieldNum; i++ {
		resMap[t.Field(i).Name] = v.Field(i)
	}
	return resMap
}

var STRUCT_TO_MAP_ONLY_SUPPORT = errors.New("StructToMap only support struct or struct pointer")

// StructToMap 结构体转为Map[string]interface{}
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	// 处理指针类型，只处理一次
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {  // 非结构体返回错误提示
		return nil, STRUCT_TO_MAP_ONLY_SUPPORT
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}else{
			out[fi.Name] = v.Field(i).Interface()
		}
	}
	return out, nil
}
// interface 转换为 string, 只支持基本数据类型
func InterfaceToString(data interface{}) string{
	switch data.(type) {
	case int64: return fmt.Sprintf("%d", data.(int64))
	case int32: return fmt.Sprintf("%d", data.(int32))
	case int: return fmt.Sprintf("%d", data.(int))
	case byte: return fmt.Sprintf("%s", string(data.(byte)))
	case float32: return fmt.Sprintf("%f", data.(float32))
	case float64: return fmt.Sprintf("%f", data.(float64))
	case string: return fmt.Sprintf("%s", data.(string))
	}
	return ""
}
//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}



//GetPlatform 获取实际运行平台
func GetPlatform(platform string) string {
	switch platform {
	case "Ios", "Android", "Mobile", "WinPhone":
		return Mobile
	default:
		return Web
	}
}

//Go 使用这个避免了协成panic导致程序退出
func Go(f func(v []interface{}), v ...interface{}) {
	go func(v []interface{}) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f(v)
	}(v)
}
