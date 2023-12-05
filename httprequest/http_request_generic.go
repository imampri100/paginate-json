package httprequest

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/valyala/fasthttp"
)

type BuilderHttpRequestGeneric struct {
	HttpRequestRaw interface{}
}

type HttpRequestGeneric struct {
	builder         BuilderHttpRequestGeneric
	reflectType     reflect.Type
	httpRequestKind HttpRequestKind
	httpRequestType HttpRequestType
	netHttpRequest  *http.Request
	fastHttpRequest *fasthttp.Request
	err             error
}

func NewHttpRequestGeneric(builder BuilderHttpRequestGeneric) (hrg *HttpRequestGeneric) {
	hrg = new(HttpRequestGeneric)

	// builder
	hrg.setBuilder(builder)

	// reflectType
	reflectType, errReflect := hrg.genReflectType()
	if errReflect != nil {
		hrg.setErr(errReflect)
		return
	}
	hrg.setReflectType(reflectType)

	// httpRequestKind, httpRequestType
	httpReqKind, httpReqType, errReqKindType := hrg.genHttpRequestKind()
	if errReqKindType != nil {
		hrg.setErr(errReqKindType)
		return
	}
	hrg.setHttpRequestType(httpReqType)
	hrg.setHttpRequestKind(httpReqKind)

	// netHttpRequest
	netHttpReq := hrg.genNetHttpRequest()
	hrg.setNetHttpRequest(netHttpReq)

	// fastHttpRequest
	fastHttpReq, errFastHttp := hrg.genFastHttpRequest()
	if errFastHttp != nil {
		hrg.setErr(errFastHttp)
		return
	}
	hrg.setFastHttpRequest(fastHttpReq)

	return
}

func (hrg HttpRequestGeneric) Me() HttpRequestGeneric {
	return hrg
}

func (hrg HttpRequestGeneric) Method() (method string) {
	reqType := hrg.getHttpRequestType()

	switch reqType {
	case NetHttpRequestType:
		{
			netHttp := hrg.getNetHttpRequest()
			method = netHttp.Method
			break
		}
	case FastHttpRequestType:
		{
			fastHttp := hrg.getFastHttpRequest()
			method = string(fastHttp.Header.Method())
			break
		}
	}

	return
}

func (hrg HttpRequestGeneric) UrlQueryGet(key string) (result string) {
	reqType := hrg.getHttpRequestType()

	switch reqType {
	case NetHttpRequestType:
		{
			netHttp := hrg.getNetHttpRequest()
			query := netHttp.URL.Query()
			result = query.Get(key)
			break
		}
	case FastHttpRequestType:
		{
			fastHttp := hrg.getFastHttpRequest()
			query := fastHttp.URI().QueryArgs()
			resBte := query.Peek(key)
			result = string(resBte)
			break
		}
	}

	return
}

func (hrg HttpRequestGeneric) Err() (err error) {
	err = hrg.getErr()
	return
}

func (hrg *HttpRequestGeneric) setBuilder(builder BuilderHttpRequestGeneric) {
	hrg.builder = builder
}

func (hrg HttpRequestGeneric) getBuilder() (builder BuilderHttpRequestGeneric) {
	builder = hrg.builder
	return
}

func (hrg *HttpRequestGeneric) setReflectType(reflectType reflect.Type) {
	hrg.reflectType = reflectType
}

func (hrg HttpRequestGeneric) getReflectType() (reflectType reflect.Type) {
	reflectType = hrg.reflectType
	return
}

func (hrg HttpRequestGeneric) genReflectType() (reflectType reflect.Type, err error) {
	builder := hrg.getBuilder()
	httpReq := builder.HttpRequestRaw
	if httpReq == nil {
		err = errors.New("http request cannot be nil")
		return
	}

	reflectType = reflect.TypeOf(httpReq)
	return
}

func (hrg *HttpRequestGeneric) setHttpRequestKind(httpRequestKind HttpRequestKind) {
	hrg.httpRequestKind = httpRequestKind
}

func (hrg HttpRequestGeneric) getHttpRequestKind() (httpRequestKind HttpRequestKind) {
	httpRequestKind = hrg.httpRequestKind
	return
}

func (hrg HttpRequestGeneric) genHttpRequestKind() (httpRequestKind HttpRequestKind, httpRequestType HttpRequestType, err error) {
	reflectType := hrg.getReflectType()

	reflectNetHttp := reflect.TypeOf(http.Request{})
	reflectNetHttpPtr := reflect.TypeOf(&http.Request{})
	reflectFastHttp := reflect.TypeOf(fasthttp.Request{})
	reflectFastHttpPtr := reflect.TypeOf(&fasthttp.Request{})

	switch {
	case reflectType.ConvertibleTo(reflectNetHttp):
		{
			httpRequestKind = StructHttpRequestKind
			httpRequestType = NetHttpRequestType
			break
		}
	case reflectType.ConvertibleTo(reflectNetHttpPtr):
		{
			httpRequestKind = PointerHttpRequestKind
			httpRequestType = NetHttpRequestType
			break
		}
	case reflectType.ConvertibleTo(reflectFastHttp):
		{
			httpRequestKind = StructHttpRequestKind
			httpRequestType = FastHttpRequestType
			break
		}
	case reflectType.ConvertibleTo(reflectFastHttpPtr):
		{
			httpRequestKind = PointerHttpRequestKind
			httpRequestType = FastHttpRequestType
			break
		}
	default:
		{
			err = fmt.Errorf("http request with type of %s is unhandled", reflectType.String())
			break
		}
	}

	return
}

func (hrg *HttpRequestGeneric) setHttpRequestType(httpRequestType HttpRequestType) {
	hrg.httpRequestType = httpRequestType
}

func (hrg HttpRequestGeneric) getHttpRequestType() (httpRequestType HttpRequestType) {
	httpRequestType = hrg.httpRequestType
	return
}

func (hrg *HttpRequestGeneric) setNetHttpRequest(netHttpRequest *http.Request) {
	hrg.netHttpRequest = netHttpRequest
}

func (hrg HttpRequestGeneric) getNetHttpRequest() (netHttpRequest *http.Request) {
	netHttpRequest = hrg.netHttpRequest
	return
}

func (hrg HttpRequestGeneric) genNetHttpRequest() (netHttpRequest *http.Request) {
	builder := hrg.getBuilder()
	httpReq := builder.HttpRequestRaw

	reqType := hrg.getHttpRequestType()
	if reqType != NetHttpRequestType {
		return
	}

	reqKind := hrg.getHttpRequestKind()

	switch reqKind {
	case StructHttpRequestKind:
		{
			concreteHttpReq := httpReq.(http.Request)
			netHttpRequest = &concreteHttpReq
			break
		}
	case PointerHttpRequestKind:
		{
			concreteHttpReq := httpReq.(*http.Request)
			netHttpRequest = concreteHttpReq
			break
		}
	}

	return
}

func (hrg *HttpRequestGeneric) setFastHttpRequest(fastHttpRequest *fasthttp.Request) {
	hrg.fastHttpRequest = fastHttpRequest
}

func (hrg HttpRequestGeneric) getFastHttpRequest() (fastHttpRequest *fasthttp.Request) {
	fastHttpRequest = hrg.fastHttpRequest
	return
}

func (hrg HttpRequestGeneric) genFastHttpRequest() (fastHttpRequest *fasthttp.Request, err error) {
	builder := hrg.getBuilder()
	httpReq := builder.HttpRequestRaw

	reqType := hrg.getHttpRequestType()
	if reqType != FastHttpRequestType {
		return
	}

	reqKind := hrg.getHttpRequestKind()

	switch reqKind {
	case StructHttpRequestKind:
		{
			err = fmt.Errorf("http request cannot use fasthttp.Request without pointer. Copying value from fasthttp.Request is forbidden")
			break
		}
	case PointerHttpRequestKind:
		{
			concreteHttpReq := httpReq.(*fasthttp.Request)
			fastHttpRequest = concreteHttpReq
			break
		}
	}

	return
}

func (hrg *HttpRequestGeneric) setErr(err error) {
	hrg.err = err
}

func (hrg *HttpRequestGeneric) clearErr() {
	hrg.err = nil
}

func (hrg HttpRequestGeneric) getErr() (err error) {
	err = hrg.err
	return
}

type UrlValues struct {
}

type HttpRequestKind string

const (
	StructHttpRequestKind  HttpRequestKind = "struct"
	PointerHttpRequestKind HttpRequestKind = "ptr"
)

type HttpRequestType string

const (
	NetHttpRequestType  HttpRequestType = "net/http.Request"
	FastHttpRequestType HttpRequestType = "fasthttp.Request"
)
