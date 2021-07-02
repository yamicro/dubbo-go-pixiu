/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package initialize

import (
	"github.com/apache/dubbo-go-pixiu/pkg/filter/accesslog"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/api"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/authority"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/logger"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/plugins"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/ratelimit"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/recovery"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/remote"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/response"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/timeout"
	sa "github.com/apache/dubbo-go-pixiu/pkg/service/api"
)

import (
	"github.com/dubbogo/dubbo-go-pixiu-filter/pkg/api/config"
)

// Run start init.
func Run(config config.APIConfig) {

	// pi start| 1. filter init
	filterInit(&config)
	// pi start| 2. api discovery service init
	apiDiscoveryServiceInit()
}

func filterInit(config *config.APIConfig) {
	accesslog.Init() 	// 访问日志组件初始化
	api.Init()			// api 组件初始化
	authority.Init()	// 权限组件初始化
	logger.Init()		// 日志组件初始化
	recovery.Init()		// 自动恢复组件初始化
	remote.Init()		// 远程调用组件初始化
	response.Init()		// 调用响应体处理组件初始化
	timeout.Init()		// 超时组件初始化
	ratelimit.Init(&config.RateLimit) // 限流组件初始化
	plugins.Init(config.PluginsGroup, config.PluginFilePath, config.Resources) // 初始化插件组
}

func apiDiscoveryServiceInit() {
	sa.Init()
}
