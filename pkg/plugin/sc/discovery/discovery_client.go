package discovery

//
type DiscoveryClient interface {

	Description() string

	GetInstances(serviceId string) []*ServiceInstance

	GetServices() []*string
}


type Client interface {

	// 注册自己
	Register() error

	// 取消注册自己
	UnRegister() error

}

