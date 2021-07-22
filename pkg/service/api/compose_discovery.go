package api

import (
	"errors"
	"github.com/apache/dubbo-go-pixiu/pkg/router"
	"github.com/dubbogo/dubbo-go-pixiu-filter/pkg/api/config"
	fr "github.com/dubbogo/dubbo-go-pixiu-filter/pkg/router"
)

const (
	LocalMemoryApiDiscoveryService = "api.ds.local_memory"
	SC = "api.ds.spring-cloud"
)

// ComposeDiscoveryService compose multiple discovery service impl, discovery local by config and any or multiple registration center.
type ComposeDiscoveryService struct {
	router *router.Route
}

func NewComposeDiscoveryService() *ComposeDiscoveryService {

	return &ComposeDiscoveryService{
		router: router.NewRoute(),
	}
}

// AddAPI adds a method to the router tree
func (l *ComposeDiscoveryService) AddAPI(api fr.API) error {
	return l.router.PutAPI(api)
}

// GetAPI returns the method to the caller
func (l *ComposeDiscoveryService) GetAPI(url string, httpVerb config.HTTPVerb) (fr.API, error) {
	if api, ok := l.router.FindAPI(url, httpVerb); ok {
		return *api, nil
	}

	return fr.API{}, errors.New("not found")
}

// ClearAPI clear all api
func (l *ComposeDiscoveryService) ClearAPI() error {
	return l.router.ClearAPI()
}

// RemoveAPIByPath remove all api belonged to path
func (l *ComposeDiscoveryService) RemoveAPIByPath(deleted config.Resource) error {
	_, groupPath := getDefaultPath()
	fullPath := getFullPath(groupPath, deleted.Path)

	l.router.DeleteNode(fullPath)
	return nil
}

// RemoveAPIByPath remove all api
func (l *ComposeDiscoveryService) RemoveAPI(fullPath string, method config.Method) error {
	l.router.DeleteAPI(fullPath, method.HTTPVerb)
	return nil
}

// ResourceChange handle modify resource event
func (l *ComposeDiscoveryService) ResourceChange(new config.Resource, old config.Resource) bool {

	return false
}

// ResourceAdd handle add resource event
func (l *ComposeDiscoveryService) ResourceAdd(res config.Resource) bool {

	return false
}

// ResourceDelete handle delete resource event
func (l *ComposeDiscoveryService) ResourceDelete(deleted config.Resource) bool {

	return false
}

// MethodChange handle modify method event
func (l *ComposeDiscoveryService) MethodChange(res config.Resource, new config.Method, old config.Method) bool {

	return false
}

// MethodAdd handle add method event
func (l *ComposeDiscoveryService) MethodAdd(res config.Resource, method config.Method) bool {

	return false
}

// MethodDelete handle delete method event
func (l *ComposeDiscoveryService) MethodDelete(res config.Resource, method config.Method) bool {

	return false
}