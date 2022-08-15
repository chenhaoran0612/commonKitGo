package app_config_center

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"github.com/chenhaoran0612/commonKitGo/tools"
)

// AddProviderService 添加提供者服务配置
// providerName 提供者名称
// interfaceName 接口名称
// methodNames 方法名称，可以多个
func AddProviderService(providerName string, interfaceName string, methodNames ...string) {
	AddProviderServiceNeedConfig(providerName, interfaceName, `retries:0 loadBalance:random protocol:grpc registry:nacos cluster:failover`, methodNames...)
}

// AddProviderServiceNeedConfig 添加提供者服务配置
// providerName 提供者名称
// configArg 配置参数， ：如：`retries:0 loadBalance:random protocol:grpc registry:nacos cluster:failover warmup:100`
// interfaceName 接口名称
// methodNames 方法名称，可以多个
func AddProviderServiceNeedConfig(providerName string, interfaceName string, configArg string, methodNames ...string) {
	_services := config.GetProviderConfig().Services
	if _services == nil {
		_services = make(map[string]*config.ServiceConfig)
	}

	// 赋予配置默认参数
	interfaceConfigArgMap := tools.StringSplit2Map(configArg, ' ', ':')
	retries, ok := interfaceConfigArgMap["retries"]
	if false == ok {
		retries = "0"
	}
	loadBalance, ok := interfaceConfigArgMap["loadBalance"]
	if false == ok {
		loadBalance = "random"
	}
	protocol, ok := interfaceConfigArgMap["Protocol"]
	if false == ok {
		protocol = "grpc"
	}
	registry, ok := interfaceConfigArgMap["Registry"]
	if false == ok {
		registry = "nacos"
	}
	cluster, ok := interfaceConfigArgMap["cluster"]
	if false == ok {
		cluster = "failover"
	}
	warmup, ok := interfaceConfigArgMap["warmup"]
	if false == ok {
		warmup = "100"
	}

	for _, methodName := range methodNames {

		if len(methodName) == 0 {
			continue
		}

		if serviceConfig, ok := _services[providerName]; ok {
			existMethod := false
			for _, method := range serviceConfig.Methods {
				if method.Name == methodName {
					existMethod = true
					break
				}
			}

			if false == existMethod {
				method := &config.MethodConfig{Name: methodName, Retries: retries, LoadBalance: loadBalance, InterfaceName: interfaceName, InterfaceId: providerName}
				serviceConfig.Methods = append(serviceConfig.Methods, method)
			}
		} else {
			serviceConfig := config.NewServiceConfig(providerName, context.Background())
			serviceConfig.InterfaceName = interfaceName
			serviceConfig.Protocol = protocol
			serviceConfig.Registry = registry
			serviceConfig.Cluster = cluster
			serviceConfig.Loadbalance = loadBalance
			serviceConfig.Warmup = warmup
			serviceConfig.GrpcMaxMessageSize = 4
			method := &config.MethodConfig{Name: methodName, Retries: retries, LoadBalance: loadBalance, InterfaceName: interfaceName, InterfaceId: providerName}
			serviceConfig.Methods = append(serviceConfig.Methods, method)
			_services[providerName] = serviceConfig
		}
	}

	config1 := config.GetProviderConfig()
	config1.Services = _services
	config.SetProviderConfig(config1)

}

// AddConsumerService 添加消费者服务配置
// consumerName 消费者名称
// interfaceName 接口名称
// methodNames 方法名称，可以多个
func AddConsumerService(consumerName string, interfaceName string, methodNames ...string) {
	AddConsumerServiceNeedConfig(consumerName, interfaceName, `retries:0 loadBalance:random protocol:grpc registry:nacos cluster:failover`, methodNames...)
}

// AddConsumerServiceNeedConfig 添加消费者服务配置
// consumerName 消费者名称
// configArg 配置参数， ：如：`retries:0 loadBalance:random protocol:grpc registry:nacos cluster:failover`
// interfaceName 接口名称
// methodNames 方法名称，可以多个
func AddConsumerServiceNeedConfig(consumerName string, interfaceName string, configArg string, methodNames ...string) {
	_services := config.GetConsumerConfig().References
	if _services == nil {
		_services = make(map[string]*config.ReferenceConfig)
	}

	// 赋予配置默认参数
	interfaceConfigArgMap := tools.StringSplit2Map(configArg, ' ', ':')
	retries, ok := interfaceConfigArgMap["retries"]
	if false == ok {
		retries = "0"
	}
	loadBalance, ok := interfaceConfigArgMap["loadBalance"]
	if false == ok {
		loadBalance = "random"
	}
	protocol, ok := interfaceConfigArgMap["Protocol"]
	if false == ok {
		protocol = "grpc"
	}
	registry, ok := interfaceConfigArgMap["Registry"]
	if false == ok {
		registry = "nacos"
	}
	cluster, ok := interfaceConfigArgMap["cluster"]
	if false == ok {
		cluster = "failover"
	}

	for _, methodName := range methodNames {

		if len(methodName) == 0 {
			continue
		}

		if serviceConfig, ok := _services[consumerName]; ok {
			existMethod := false
			for _, method := range serviceConfig.Methods {
				if method.Name == methodName {
					existMethod = true
					break
				}
			}

			if false == existMethod {
				method := &config.MethodConfig{Name: methodName, Retries: retries, LoadBalance: loadBalance, InterfaceName: interfaceName, InterfaceId: consumerName}
				serviceConfig.Methods = append(serviceConfig.Methods, method)
			}
		} else {
			serviceConfig := config.NewReferenceConfig(consumerName, context.Background())
			serviceConfig.InterfaceName = interfaceName
			serviceConfig.Protocol = protocol
			serviceConfig.Registry = registry
			serviceConfig.Cluster = cluster
			serviceConfig.Loadbalance = loadBalance
			method := &config.MethodConfig{Name: methodName, Retries: retries, LoadBalance: loadBalance, InterfaceName: interfaceName, InterfaceId: consumerName}
			serviceConfig.Methods = append(serviceConfig.Methods, method)
			_services[consumerName] = serviceConfig
		}
	}

	config1 := config.GetConsumerConfig()
	config1.References = _services
	config.SetConsumerConfig(config1)

}
