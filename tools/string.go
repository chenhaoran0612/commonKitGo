package tools

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func FirstToUpper(data string) string{
	if len(data)>1{
		first := data[0]
		if first>='a' && first<='z'{
			return string(byte(first-32))+data[1:]
		}
	}

	return data
}

func FirstToLower(data string) string{
	if len(data)>1{
		first := data[0]
		if first>='A' && first<='Z'{
			return string(byte(first+32))+data[1:]
		}
	}

	return data
}

func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RandomNum(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func GetTraceId() string{
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}



// FindSplit 通过splitChar分割dataArr数据
func FindSplit(dataArr string, splitChar byte)[]string{

	result := make([]string, 0)
	for{
		splitIndex := FindSplitAlpha(splitChar , dataArr)
		if splitIndex==-1{
			result = append(result, dataArr)
			break
		}


		result = append(result, dataArr[:splitIndex])


		if splitIndex+1 == len(dataArr){
			break
		}

		dataArr = dataArr[splitIndex+1:]

	}

	return  result
}

// FindSplitAlpha 找到第一个分割字符的位置，splitChar 为 分割字符，dataArr 为 数据数组
func FindSplitAlpha(splitChar byte, dataArr string) int {
	for i, dataByte := range []byte(dataArr){
		if splitChar == dataByte && i>0 && dataArr[i-1]!='\\'{
			return  i
		}

		if splitChar == dataByte && i==0{
			return  i
		}
	}

	return -1
}


func StringSplit2Map(data string, arrSplit byte, mapSplit byte) map[string]string{
	datas := FindSplit(data, arrSplit)
	result := make(map[string]string)
	index := 0
	length := len(datas)
	for{
		if index == length{
			break
		}


		commaIndex := FindSplitAlpha(mapSplit, datas[index] )
		if commaIndex==-1{
			index = index+1
			continue
		}
		result[datas[index][0:commaIndex]] = datas[index][commaIndex+1:]
		index = index+1
	}

	return result
}