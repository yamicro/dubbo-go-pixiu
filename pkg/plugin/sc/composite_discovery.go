package sc

import (
	"github.com/apache/dubbo-go-pixiu/pkg/router"
)


type CompositeDiscoveryService struct {
	router *router.Route
}

func NewCompositeDiscoveryService() *CompositeDiscoveryService {

	return &CompositeDiscoveryService{
		router: router.NewRoute(),
	}
}

func (ds *CompositeDiscoveryService) Description() string {
	return ""
}

func (ds *CompositeDiscoveryService) GetInstances(serviceId string) []*ServiceInstance {
	return nil
}

func (ds *CompositeDiscoveryService) GetServices() []*string {
	return nil
}