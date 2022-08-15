package app_config_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/common/extension"
	"github.com/apache/dubbo-go/config"
	"github.com/apache/dubbo-go/config_center"
	"sync"
)

var DYNAMIC_CONFIG config_center.DynamicConfiguration
var dynamic_once 	sync.Once

// 初始化nacos动态配置
func InitNacosConfig(baseConfig *config.BaseConfig){

	dynamic_once.Do(func() {
		configCenterConfig := baseConfig.ConfigCenterConfig
		regurl, _ := common.NewURL(fmt.Sprintf("registry://%s", configCenterConfig.Address))
		dyconfig := extension.GetConfigCenterFactory("nacos")
		DYNAMIC_CONFIG,_ = dyconfig.GetDynamicConfiguration(regurl)
	})

}

// 获取指定nacos组中的dataId json配置信息
func GetNacosJSONConfig( dataId string, group string) (map[string]interface{}, error){

	nacosValue, err := DYNAMIC_CONFIG.GetProperties(dataId, config_center.WithGroup(group))
	if err!=nil{return nil, err}
	if nacosValue==""{
		return nil, errors.New("dynmaic config is empty")
	}

	var ret map[string]interface{}
	err = json.Unmarshal([]byte(nacosValue), &ret)
	return ret, err
}

// 获取指定nacos组中的dataId string配置信息
func GetNacosStringConfig( dataId string, group string) (string, error){
	nacosValue, err := DYNAMIC_CONFIG.GetProperties(dataId, config_center.WithGroup(group))
	if err!=nil{return "", err}
	return nacosValue, err
}


// 获取指定nacos组中的dataId bool配置信息
func GetNacosBoolConfig( dataId string, group string) (bool, error){
	nacosValue, err := DYNAMIC_CONFIG.GetProperties(dataId, config_center.WithGroup(group))
	if err!=nil{return false, err}
	if nacosValue=="true" || nacosValue=="1" || nacosValue=="yes" || nacosValue=="on"{
		return true,nil
	}

	return false, err
}

