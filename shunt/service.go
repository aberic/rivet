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

package shunt

type Balance struct {
	Name    string  `json:"name"`
	Service Service `json:"service"`
}

type Services struct {
	Services []*Service `json:"services"`
}

// 服务器信息
type Service struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (services *Services) Add(service Service) {
	services.Services = append(services.Services, &service)
}

func NewAddress(host string, port int) *Service {
	return &Service{
		Host: host,
		Port: port,
	}
}

func (a *Service) GetHost() string {
	return a.Host
}

func (a *Service) GetPort() int {
	return a.Port
}
