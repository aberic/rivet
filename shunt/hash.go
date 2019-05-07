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

import (
	"fmt"
	"github.com/ennoo/rivet/server"
	"hash/crc32"
	"math/rand"
)

// HashBalance 负载均衡 hash 策略实体
type HashBalance struct {
	key string
}

// RunBalance 负载均衡 round 策略实现
func (p *HashBalance) RunBalance(serviceName string, key ...string) (add *server.Service, err error) {
	services := server.ServiceGroup[serviceName].Services
	defKey := fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(services)
	if lens == 0 {
		err = fmt.Errorf("no balance")
		return
	}
	hashVal := crc32.Checksum([]byte(defKey), crc32.MakeTable(crc32.IEEE))
	index := int(hashVal) % lens
	add = services[index]
	return
}
