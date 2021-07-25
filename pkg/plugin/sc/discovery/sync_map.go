package discovery

import "sync"

type SyncMap struct {
	smap *sync.Map
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		smap: new(sync.Map),
	}
}

func (sm *SyncMap) Get(key string) ([]*ServiceInstance, bool) {
	val, exist := sm.smap.Load(key)
	if !exist {
		return nil, false
	}

	info, ok := val.([]*ServiceInstance)
	if !ok {
		return nil, false
	}

	return info, true
}

func (sm *SyncMap) Put(key string, val []*InstanceInfo) {
	sm.smap.Store(key, val)
}

func (sm *SyncMap) Each(eachFunc func(key string, val []*InstanceInfo) bool) {
	sm.smap.Range(func(key, value interface{}) bool {
		return eachFunc(key.(string), value.([]*InstanceInfo))
	})
}

func (sm *SyncMap) GetMap() *sync.Map {
	return sm.smap
}