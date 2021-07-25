package discovery

import (
	"fmt"
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	"github.com/wanghongfei/go-eureka-client/eureka"
	"net"
	"strconv"
	"time"
)

type EurekaClient struct {

	client *eureka.Client

	// 保存服务地址
	// key: 服务名:版本号, 版本号为eureka注册信息中的metadata[version]值
	// val: []*InstanceInfo
	registryMap 			*SyncMap
}

var gogateApp *eureka.InstanceInfo
var instanceId = ""

var ticker *time.Ticker
var tickerCloseChan chan struct{}

func (c *EurekaClient) Register() error {
	//ip, err := GetFirstNoneLoopIp()
	//if nil != err {
	//	logger.Errorf("fail ", err)
	//	return err
	//}
	//
	//
	//instanceId = ip + ":" + strconv.Itoa(conf.App.ServerConfig.Port)
	//
	//gogateApp = eureka.NewInstanceInfo(
	//	instanceId,
	//	conf.App.ServerConfig.AppName,
	//	ip,
	//	conf.App.ServerConfig.Port,
	//	conf.App.EurekaConfig.EvictionDuration,
	//	false,
	//)
	//gogateApp.Metadata = &eureka.MetaData{
	//	Class: "",
	//	Map: map[string]string {"version": conf.App.Version},
	//}
	//
	//err = c.client.RegisterInstance(conf.App.ServerConfig.AppName, gogateApp)
	//if nil != err {
	//	return perr.WrapSystemErrorf(err, "failed to register to eureka")
	//}
	//
	//// 心跳
	//go func() {
	//	ticker = time.NewTicker(time.Second * time.Duration(conf.App.EurekaConfig.HeartbeatInterval))
	//	tickerCloseChan = make(chan struct{})
	//
	//	for {
	//		select {
	//		case <-ticker.C:
	//			c.heartbeat()
	//
	//		case <-tickerCloseChan:
	//			Log.Info("heartbeat stopped")
	//			return
	//
	//		}
	//	}
	//}()

	return nil
}

func (c *EurekaClient) UnRegister() error {

	logger.Info("unregistering %s", instanceId)
	err := c.client.UnregisterInstance("gogate", instanceId)

	if nil != err {
		logger.Error(err)
		return err
	}

	logger.Info("done unregistration")
	return nil
}

func GetFirstNoneLoopIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if nil != err {
		return "", fmt.Errorf("failed to fetch interfaces => %w", err)
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no first-none-loop ip found")
}

func (c *EurekaClient) Description() string {
	return ""
}

func (c *EurekaClient) GetServices() []*string {
	return nil
}

func (c *EurekaClient) GetInstances() ([]*InstanceInfo, error) {
	apps, err := c.client.GetApplications()
	if nil != err {
		fmt.Errorf("error find service")
		return nil, err
	}

	var instances []*InstanceInfo
	for _, app := range apps.Applications {
		// 服务名
		servName := app.Name

		// 遍历每一个实例
		for _, ins := range app.Instances {
			// 跳过无效实例
			if nil == ins.Port || ins.Status != "UP" {
				continue
			}

			addr := ins.HostName + ":" + strconv.Itoa(ins.Port.Port)
			var meta map[string]string
			if nil != ins.Metadata {
				meta = ins.Metadata.Map
			}

			instances = append(
				instances,
				&InstanceInfo{
					ServiceName: servName,
					Addr: addr,
					MetaInfo: meta,
				},
			)
		}
	}

	return instances, nil
}