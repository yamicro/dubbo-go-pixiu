package plugin

import (
	"github.com/apache/dubbo-go-pixiu/pkg/config"
	"github.com/apache/dubbo-go-pixiu/pkg/plugin/sc"
)

func BeforeStart()  {

	// spring cloud start before
	sc.BeforeStart(config.GetBootstrap())
}
