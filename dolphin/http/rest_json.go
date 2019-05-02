package http

import (
	"net/http"
	"net/url"
)

type RestJsonHandler struct {
	RestHandler
	Header  http.Header
	Cookies []*http.Cookie
}

func (handler RestJsonHandler) ObtainUri() string {
	return handler.RestHandler.uri
}

func (handler RestJsonHandler) ObtainParam() interface{} {
	return handler.RestHandler.param
}

func (handler RestJsonHandler) ObtainValue() url.Values {
	return nil
}

func (handler RestJsonHandler) ObtainRemoteServer() string {
	return handler.RestHandler.RemoteServer
}

func (handler RestJsonHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "application/json")
	return handler.Header
}

func (handler RestJsonHandler) ObtainCookies() []*http.Cookie {
	return handler.Cookies
}

func (handler *RestJsonHandler) Post() (body []byte, err error) {
	return requestJson(http.MethodPost, handler)
}

func (handler *RestJsonHandler) Put() (body []byte, err error) {
	return requestJson(http.MethodPut, handler)
}

func (handler *RestJsonHandler) Delete() (body []byte, err error) {
	return requestJson(http.MethodDelete, handler)
}

func (handler *RestJsonHandler) Patch() (body []byte, err error) {
	return requestJson(http.MethodPatch, handler)
}

func (handler *RestJsonHandler) Options() (body []byte, err error) {
	return requestJson(http.MethodOptions, handler)
}

func (handler *RestJsonHandler) Head() (body []byte, err error) {
	return requestJson(http.MethodHead, handler)
}

func (handler *RestJsonHandler) Connect() (body []byte, err error) {
	return requestJson(http.MethodConnect, handler)
}

func (handler *RestJsonHandler) Trace() (body []byte, err error) {
	return requestJson(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestJsonHandler) Get() (body []byte, err error) {
	return get(handler)
}
