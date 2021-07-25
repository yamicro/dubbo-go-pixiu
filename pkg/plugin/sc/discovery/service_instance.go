package discovery

type ServiceInstance interface {

	// 获取注册服务实例ID
	GetInstanceId() string

	// 服务实例ID
	GetServiceId() string

	GetHost() string

	GetPort() int

	Meta() map[string]interface{}
}

type InstanceInfo struct {

	ServiceName		string

	// 格式为 host:port
	Addr 			string

	// 此实例附加信息
	MetaInfo 			map[string]string
}


func (info *InstanceInfo) GetInstanceId() string {
	return info.ServiceName
}

func (info *InstanceInfo) GetServiceId() string {
	return info.ServiceName
}

func (info *InstanceInfo) GetHost() string {
	return info.ServiceName
}

func (info *InstanceInfo) GetPort() int {
	return 1111
}

func (info *InstanceInfo) Meta() map[string]interface{} {
	return nil
}