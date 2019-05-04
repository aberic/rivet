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
	"hash/crc32"
	"math/rand"
)

// 负载均衡接口Hash实现
func init() {
	RegisterBalance("hash", &HashBalance{})
}

type HashBalance struct {
	key string
}

func (p *HashBalance) DoBalance(adds []*Address, key ...string) (add *Address, err error) {
	defKey := fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(adds)
	if lens == 0 {
		err = fmt.Errorf("no balance")
		return
	}
	hashVal := crc32.Checksum([]byte(defKey), crc32.MakeTable(crc32.IEEE))
	index := int(hashVal) % lens
	add = adds[index]
	return

}
