package discovery

import (
	"errors"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/plugins"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/ratelimit"
	"github.com/apache/dubbo-go-pixiu/pkg/router"
	"github.com/dubbogo/dubbo-go-pixiu-filter/pkg/api/config"
	ratelimitConf "github.com/dubbogo/dubbo-go-pixiu-filter/pkg/api/config/ratelimit"
	fr "github.com/dubbogo/dubbo-go-pixiu-filter/pkg/router"
)

// DiscoveryServiceAdapter is the local cached API discovery service
type DiscoveryServiceAdapter struct {

	router *router.Route
}

func NewDiscoveryServiceAdapter() *DiscoveryServiceAdapter {
	return &DiscoveryServiceAdapter{
		router: router.NewRoute(),
	}
}

// AddAPI adds a method to the router tree
func (l *DiscoveryServiceAdapter) AddAPI(api fr.API) error {
	return l.router.PutAPI(api)
}

// GetAPI returns the method to the caller
func (l *DiscoveryServiceAdapter) GetAPI(url string, httpVerb config.HTTPVerb) (fr.API, error) {
	if api, ok := l.router.FindAPI(url, httpVerb); ok {
		return *api, nil
	}

	// pi 从 client 中获取服务

	return fr.API{}, errors.New("not found")
}

// ClearAPI clear all api
func (l *DiscoveryServiceAdapter) ClearAPI() error {
	return l.router.ClearAPI()
}

// RemoveAPIByPath remove all api belonged to path
func (l *DiscoveryServiceAdapter) RemoveAPIByPath(deleted config.Resource) error {
	return nil
}

// RemoveAPIByPath remove all api
func (l *DiscoveryServiceAdapter) RemoveAPI(fullPath string, method config.Method) error {
	l.router.DeleteAPI(fullPath, method.HTTPVerb)
	return nil
}



// ResourceChange handle modify resource event
func (l *DiscoveryServiceAdapter) ResourceChange(new config.Resource, old config.Resource) bool {

	return false
}

// ResourceAdd handle add resource event
func (l *DiscoveryServiceAdapter) ResourceAdd(res config.Resource) bool {

	return false
}

// ResourceDelete handle delete resource event
func (l *DiscoveryServiceAdapter) ResourceDelete(deleted config.Resource) bool {

	return false
}

// MethodChange handle modify method event
func (l *DiscoveryServiceAdapter) MethodChange(res config.Resource, new config.Method, old config.Method) bool {

	return false
}

// MethodAdd handle add method event
func (l *DiscoveryServiceAdapter) MethodAdd(res config.Resource, method config.Method) bool {

	return false
}

// MethodDelete handle delete method event
func (l *DiscoveryServiceAdapter) MethodDelete(res config.Resource, method config.Method) bool {

	return false
}

func (l *DiscoveryServiceAdapter) PluginPathChange(filePath string) {
	plugins.OnFilePathChange(filePath)
}

func (l *DiscoveryServiceAdapter) PluginGroupChange(group []config.PluginsGroup) {
	plugins.OnGroupUpdate(group)
}

func (l *DiscoveryServiceAdapter) RateLimitChange(c *ratelimitConf.Config) {
	ratelimit.OnUpdate(c)
}