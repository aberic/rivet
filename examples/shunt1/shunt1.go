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

package main

import (
	"fmt"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/server"
	"github.com/ennoo/rivet/shunt"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var adds []*server.Service

func main() {
	rivet.Initialize(true, true, true)
	rivet.Shunt.Register("test", &shunt.RoundRobinBalance{Position: 0})
	rivet.Shunt.Register("test1", &shunt.RandomBalance{})
	rivet.Shunt.Register("test2", &shunt.HashBalance{Key: []string{}})
	//addAddress()
	rivet.ListenAndServe(rivet.SetupRouter(testShunt1), "8083")
}

func addAddress() {
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := server.NewService(host, port)
		adds = append(adds, one)
	}
}

func b(serviceName string) {
	for {
		add, err := shunt.RunShunt(serviceName)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(add)
		time.Sleep(time.Second)
	}
}

func testShunt1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet3")
	vRepo.GET("/shunt/:serviceName", shunt3)
	vRepo.POST("/shunt", shunt4)
}

func shunt3(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		serviceName := context.Param("serviceName")
		rivet.Shunt.Register(serviceName, &shunt.RoundRobinBalance{Position: 0})
		b(serviceName)
		result.SaySuccess(context, "test2")
	})
}

func shunt4(context *gin.Context) {
	rivet.Request().Callback(context, http.MethodPost, "test", "rivet/shunt", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}
