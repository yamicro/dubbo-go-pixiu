package sc

import (
	"fmt"
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	"sync"
	"time"
)
const REGISTRY_REFRESH_INTERVAL = 30

type periodicalRefreshClient struct {
	client Client
}

func newPeriodicalRefresh(c Client) *periodicalRefreshClient {
	return &periodicalRefreshClient{c}
}

// 向eureka查询注册列表, 刷新本地列表
func (r *periodicalRefreshClient) StartPeriodicalRefresh() error {
	logger.Info("refresh registry every %d sec", REGISTRY_REFRESH_INTERVAL)

	refreshRegistryChan := make(chan error)

	isBootstrap := true
	go func() {
		ticker := time.NewTicker(REGISTRY_REFRESH_INTERVAL * time.Second)

		for {
			logger.Info("registry refresh started")
			err := r.doRefresh()
			if nil != err {
				// 如果是第一次查询失败, 退出程序
				if isBootstrap {
					refreshRegistryChan <- fmt.Errorf("failed to refresh registry")
					return

				} else {
					logger.Error(err)
				}

			}
			logger.Info("done refreshing registry")

			if isBootstrap {
				isBootstrap = false
				close(refreshRegistryChan)
			}

			<-ticker.C
		}
	}()

	return <- refreshRegistryChan
}

func (r *periodicalRefreshClient) doRefresh() error {
	instances := r.client.GetInstances("213") // pi

	if nil == instances {
		logger.Info("no instance found")
		return nil
	}

	logger.Infof("total app count: %d", len(instances))

	newRegistryMap := r.groupByService(instances)

	r.refreshRegistryMap(newRegistryMap)

	return nil

}


// 将所有实例按服务名进行分组
func (r *periodicalRefreshClient) groupByService(instances []*InstanceInfo) *sync.Map {
	servMap := new(sync.Map)
	for _, ins := range instances {
		infosGeneric, exist := servMap.Load(ins.ServiceName)
		if !exist {
			infosGeneric = make([]*InstanceInfo, 0, 5)
			infosGeneric = append(infosGeneric.([]*InstanceInfo), ins)

		} else {
			infosGeneric = append(infosGeneric.([]*InstanceInfo), ins)
		}
		servMap.Store(ins.ServiceName, infosGeneric)
	}
	return servMap
}


// 更新本地注册列表
// s: gogate server对象
// newRegistry: 刚从eureka查出的最新服务列表
func (r *periodicalRefreshClient) refreshRegistryMap(newRegistry *sync.Map) {
	if nil == r.client.GetInternalRegistryStore() {
		r.client.SetInternalRegistryStore(NewSyncMap())
	}

	// 找出本地列表存在, 但新列表中不存在的服务
	exclusiveKeys, _ := FindExclusiveKey(r.client.GetInternalRegistryStore().GetMap(), newRegistry)
	// 删除本地多余的服务
	DelKeys(r.client.GetInternalRegistryStore().GetMap(), exclusiveKeys)
	// 将新列表中的服务合并到本地列表中
	MergeSyncMap(newRegistry, r.client.GetInternalRegistryStore().GetMap())
}
