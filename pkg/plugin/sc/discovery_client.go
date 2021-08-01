package sc

type Client interface {

	// 注册自己
	Register() error

	// 取消注册自己
	UnRegister() error

	Description() string

	// 从本地缓存中查询指定服务的全部实例信息
	Get(string) []*InstanceInfo

	GetInstances(serviceId string) []*InstanceInfo

	GetServices() []*string

	// 获取内部保存的注册表
	GetInternalRegistryStore() *SyncMap

	// 更新内部保存的注册表
	SetInternalRegistryStore(*SyncMap)

	// 启动注册信息定时刷新逻辑
	StartPeriodicalRefresh() error
}

