package httprequest

import (
	"fmt"
	"paginatejson/lib"
	"strings"
)

type IHttpRequest interface {
	Method() string
	UrlQueryGet(key string) string
}

type BuilderHttpRequestParser struct {
	Config IConfigRequest
}

type HttpRequestMethod string

const (
	GetMethod  HttpRequestMethod = "GET"
	PostMethod HttpRequestMethod = "POST"
)

func GenHttpRequestMethod(inputMethod string) (result HttpRequestMethod, err error) {
	switch {
	case strings.EqualFold(inputMethod, GetMethod.String()):
		{
			result = GetMethod
			break
		}
	case strings.EqualFold(inputMethod, PostMethod.String()):
		{
			result = PostMethod
			break
		}
	default:
		{
			err = fmt.Errorf("http request method %s is unhandled", inputMethod)
			break
		}
	}
	return
}

func (hrm HttpRequestMethod) String() string {
	return string(hrm)
}

type HttpRequestParser struct {
	builder BuilderHttpRequestParser
	request IHttpRequest
	method  HttpRequestMethod
	size    string
	page    string
	sort    string
	order   string
	filters string
	fields  string
	err     error
}

func NewHttpRequestParser(builder BuilderHttpRequestParser) (hrp *HttpRequestParser) {
	hrp = new(HttpRequestParser)

	// builder
	hrp.setBuilder(builder)

	return
}

func (hrp *HttpRequestParser) Parse(request IHttpRequest) {
	// clearErr first
	hrp.clearErr()

	// request
	hrp.setRequest(request)

	// method
	method, errMethod := hrp.genMethod()
	if errMethod != nil {
		hrp.setErr(errMethod)
		return
	}
	hrp.setMethod(method)

	// size
	size := hrp.genSize()
	hrp.setSize(size)

	// page
	page := hrp.genPage()
	hrp.setPage(page)

	// sort
	sort := hrp.genSort()
	hrp.setSort(sort)

	// order
	order := hrp.genOrder()
	hrp.setOrder(order)

	// filters
	filters := hrp.genFilters()
	hrp.setFilters(filters)

	// fields
	fields := hrp.genFields()
	hrp.setFields(fields)
}

func (hrp HttpRequestParser) Method() (method HttpRequestMethod) {
	method = hrp.getMethod()
	return
}

func (hrp HttpRequestParser) Size() (size string) {
	size = hrp.getSize()
	return
}

func (hrp HttpRequestParser) Page() (page string) {
	page = hrp.getPage()
	return
}

func (hrp HttpRequestParser) Sort() (sort string) {
	sort = hrp.getSort()
	return
}

func (hrp HttpRequestParser) Order() (order string) {
	order = hrp.getOrder()
	return
}

func (hrp HttpRequestParser) Filters() (filters string) {
	filters = hrp.getFilters()
	return
}

func (hrp HttpRequestParser) Fields() (fields string) {
	fields = hrp.getFields()
	return
}

func (hrp HttpRequestParser) Err() (err error) {
	err = hrp.getErr()
	return
}

func (hrp *HttpRequestParser) setBuilder(builder BuilderHttpRequestParser) {
	hrp.builder = builder
}

func (hrp HttpRequestParser) getBuilder() (builder BuilderHttpRequestParser) {
	builder = hrp.builder
	return
}

func (hep *HttpRequestParser) setRequest(request IHttpRequest) {
	hep.request = request
}

func (hep HttpRequestParser) getRequest() IHttpRequest {
	return hep.request
}

func (hrp *HttpRequestParser) setMethod(method HttpRequestMethod) {
	hrp.method = method
}

func (hrp HttpRequestParser) getMethod() (method HttpRequestMethod) {
	method = hrp.method
	return
}

func (hrp HttpRequestParser) genMethod() (method HttpRequestMethod, err error) {
	req := hrp.getRequest()
	reqMethod := req.Method()

	met, errMet := GenHttpRequestMethod(reqMethod)
	if errMet != nil {
		err = errMet
		return
	}

	method = met
	return
}

func (hrp *HttpRequestParser) setSize(size string) {
	hrp.size = size
}

func (hrp HttpRequestParser) getSize() (size string) {
	size = hrp.size
	return
}

func (hrp HttpRequestParser) genSize() (size string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	sizeParams := conf.SizeParams()
	req := hrp.getRequest()

	for idxParams := range sizeParams {
		itemParams := sizeParams[idxParams]

		resSize := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resSize) {
			continue
		}

		size = resSize
		break
	}

	return
}

func (hrp *HttpRequestParser) setPage(page string) {
	hrp.page = page
}

func (hrp HttpRequestParser) getPage() (page string) {
	page = hrp.page
	return
}

func (hrp HttpRequestParser) genPage() (page string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	pageParams := conf.PageParams()

	req := hrp.getRequest()

	for idxParams := range pageParams {
		itemParams := pageParams[idxParams]

		resPage := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resPage) {
			continue
		}

		page = resPage
		break
	}

	return
}

func (hrp *HttpRequestParser) setSort(sort string) {
	hrp.sort = sort
}

func (hrp HttpRequestParser) getSort() (sort string) {
	sort = hrp.sort
	return
}

func (hrp HttpRequestParser) genSort() (sort string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	sortParams := conf.SortParams()
	req := hrp.getRequest()

	for idxParams := range sortParams {
		itemParams := sortParams[idxParams]

		resSort := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resSort) {
			continue
		}

		sort = resSort
		break
	}

	return
}

func (hrp *HttpRequestParser) setOrder(order string) {
	hrp.order = order
}

func (hrp HttpRequestParser) getOrder() (order string) {
	order = hrp.order
	return
}

func (hrp HttpRequestParser) genOrder() (order string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	orderParams := conf.OrderParams()
	req := hrp.getRequest()

	for idxParams := range orderParams {
		itemParams := orderParams[idxParams]

		resOrder := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resOrder) {
			continue
		}

		order = resOrder
		break
	}

	return
}

func (hrp *HttpRequestParser) setFilters(filters string) {
	hrp.filters = filters
}

func (hrp HttpRequestParser) getFilters() (filters string) {
	filters = hrp.filters
	return
}

func (hrp HttpRequestParser) genFilters() (filters string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	filtersParams := conf.FiltersParams()
	req := hrp.getRequest()

	for idxParams := range filtersParams {
		itemParams := filtersParams[idxParams]

		resFilters := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resFilters) {
			continue
		}

		filters = resFilters
		break
	}

	return
}

func (hrp *HttpRequestParser) setFields(fields string) {
	hrp.fields = fields
}

func (hrp HttpRequestParser) getFields() (fields string) {
	fields = hrp.fields
	return
}

func (hrp HttpRequestParser) genFields() (fields string) {
	builder := hrp.getBuilder()
	conf := builder.Config
	fieldsParams := conf.FieldsParams()
	req := hrp.getRequest()

	for idxParams := range fieldsParams {
		itemParams := fieldsParams[idxParams]

		resFields := req.UrlQueryGet(itemParams)
		if lib.IsEmptyStr(resFields) {
			continue
		}

		fields = resFields
		break
	}

	return
}

func (hrp *HttpRequestParser) setErr(err error) {
	hrp.err = err
}

func (hrp *HttpRequestParser) clearErr() {
	hrp.err = nil
}

func (hrp HttpRequestParser) getErr() (err error) {
	err = hrp.err
	return
}
