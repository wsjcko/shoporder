package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
)

//设置配置中心
func GetConsulConfig(addr string, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心的地址
		consul.WithAddress(addr),
		//设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		//是否移除前缀，这里是设置为true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	//配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	//加载配置
	err = config.Load(consulSource)
	return config, err
}
