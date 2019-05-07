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

package request

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

// RestTextHandler 处理 text 请求发送内容
type RestTextHandler struct {
	RestHandler
	Values url.Values
}

// ObtainRemoteServer 获取本次 http 请求服务根路径 如：localhost:8080
func (handler *RestTextHandler) ObtainRemoteServer() string {
	return handler.RemoteServer
}

// ObtainUri 获取本次 http 请求服务方法路径 如：/user/login
func (handler *RestTextHandler) ObtainUri() string {
	return handler.RestHandler.Uri
}

// ObtainBody 获取本次 http 请求 body io
func (handler *RestTextHandler) ObtainBody() io.Reader {
	return bytes.NewBufferString(handler.Values.Encode())
}

// ObtainHeader 获取本次 http 请求 header
func (handler *RestTextHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "text/html")
	return handler.Header
}

// ObtainCookies 获取本次 http 请求 cookies
func (handler *RestTextHandler) ObtainCookies() []http.Cookie {
	return handler.Cookies
}

// Post 发起 Post 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Post() (body []byte, err error) {
	return request(http.MethodPost, handler)
}

// Put 发起 Put 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Put() (body []byte, err error) {
	return request(http.MethodPut, handler)
}

// Delete 发起 Delete 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Delete() (body []byte, err error) {
	return request(http.MethodDelete, handler)
}

// Patch 发起 Patch 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Patch() (body []byte, err error) {
	return request(http.MethodPatch, handler)
}

// Options 发起 Options 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Options() (body []byte, err error) {
	return request(http.MethodOptions, handler)
}

// Head 发起 Head 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Head() (body []byte, err error) {
	return request(http.MethodHead, handler)
}

// Connect 发起 Connect 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Connect() (body []byte, err error) {
	return request(http.MethodConnect, handler)
}

// Trace 发起 Trace 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Trace() (body []byte, err error) {
	return request(http.MethodTrace, handler)
}

// Get 发起 Get 请求，body 为请求后的返回内容，err 指出请求出错原因
func (handler *RestTextHandler) Get() (body []byte, err error) {
	return get(handler)
}