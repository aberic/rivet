/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package rivet

import (
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/request"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/log"
)

var (
	serviceName = env.GetEnvDefault("SERVICE_NAME", "null")
)

// Log 提供日志调用入口
func Log() *log.Logger {
	return log.GetLogInstance()
}

// Shunt 提供负载调用入口
func Shunt() *shunt.Shunt {
	return shunt.GetShuntInstance()
}

// Response 提供实例化调用 Do 方法，并内置返回策略
func Response() *response.Response {
	return response.SyncPoolGetResponse()
}

// Request 提供实例化调用请求方法，并内置返回策略
func Request() *request.Request {
	return request.SyncPoolGetRequest()
}
