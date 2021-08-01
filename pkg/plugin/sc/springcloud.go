package sc

import (
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	"github.com/apache/dubbo-go-pixiu/pkg/model"
)

var (
	client Client
)

func BeforeStart(bootstrap *model.Bootstrap) {

	cloud := bootstrap.SpringCloud
	if cloud.Eureka.Enable {

		eurekaClient, err := NewEurekaClient(cloud.Eureka.ConfigFile)

		if err != nil {
			logger.Error("init eureka client fail", err)
		}

		// 定时刷新
		eurekaClient.StartPeriodicalRefresh()

		client = eurekaClient
	}
}
