package sc

import (
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	"github.com/apache/dubbo-go-pixiu/pkg/model"
	"github.com/apache/dubbo-go-pixiu/pkg/plugin/sc/discovery"
)

var (
	client discovery.Client
)

func BeforeStart(bootstrap *model.Bootstrap) {

	cloud := bootstrap.SpringCloud
	if cloud.Eureka.Enable {

		eurekaClient, err := discovery.NewEurekaClient(cloud.Eureka.ConfigFile)

		if err != nil {
			logger.Error("init eureka client fail", err)
		}

		client = eurekaClient
	}
}
