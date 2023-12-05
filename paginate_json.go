package paginatejson

import (
	"errors"
	"paginatejson/configrequest"
	"paginatejson/httprequest"
	"paginatejson/lib"
	"strconv"
)

type ICacheAdapter interface {
	// Set cache with key
	Set(key string, value string) error
	// Get cache by key
	Get(key string) (string, error)
	// IsValid check if cache is valid
	IsValid(key string) bool
	// Clear clear cache by key
	Clear(key string) error
	// ClearPrefix clear cache by key prefix
	ClearPrefix(keyPrefix string) error
	// Clear all cache
	ClearAll() error
}

type DefaultOperator string

const (
	Or  DefaultOperator = "or"
	And DefaultOperator = "and"
)

func (dfo DefaultOperator) String() string {
	return string(dfo)
}

// Config for customize pagination result
type Config struct {
	Operator             DefaultOperator // default: or
	DefaultSize          int64           // default: 10
	CustomParamEnabled   bool            // default: false
	SortParams           []string        // default: []string{"sort"}
	PageParams           []string        // default: []string{"page"}
	OrderParams          []string        // default: []string{"order"}
	SizeParams           []string        // default: []string{"size"}
	FiltersParams        []string        // default: []string{"filters"}
	FieldsParams         []string        // default: []string{"fields"}
	FieldSelectorEnabled bool            // default: false
	// CacheAdapter         ICacheAdapter                          `json:"-"` // required; Not implemented, because didn't get usecase
	// CachePrefix          string                                 // default: paginate_json; Not implemented, because didn't get usecase
	// CacheCompression     bool                                   // default: true; Not implemented, because didn't get usecase
	JSONMarshal   func(v interface{}) ([]byte, error)    `json:"-"`
	JSONUnmarshal func(data []byte, v interface{}) error `json:"-"`
	ErrorEnabled  bool                                   // default: false
}

type IQuery interface {
	Query(httpRawRequest interface{}) (result httprequest.IQuery)
}

type PaginateJson struct {
	config             Config
	configRequest      IConfigRequest
	rawData            interface{}
	requestData        []interface{}
	httpRequest        interface{}
	httpRequestHandler *httprequest.HttpRequest
	pageRequest        PageRequest
	result             Page
	err                error
}

func NewPaginateJson(config Config) (pj *PaginateJson) {
	pj = new(PaginateJson)
	pj.WithConfig(config)
	return
}

func (pj *PaginateJson) WithConfig(config Config) {
	// clearErr first
	pj.clearErr()

	// config
	pj.setConfig(config)

	// configRequest
	configRequest, errConfigRequest := pj.genConfigRequest()
	if errConfigRequest != nil {
		pj.setErr(errConfigRequest)
		return
	}
	pj.setConfigRequest(configRequest)
}

func (pj *PaginateJson) With(data interface{}) (handler IRequestHandler) {
	pj.clearErr()

	pj.setRawData(data)

	reqData, isValid := pj.genRequestData()
	if !isValid {
		pj.setErr(errors.New("request data is invalid. Make sure your data is slice data"))
		return
	}
	pj.setRequestData(reqData)

	handler = pj
	return
}

func (pj *PaginateJson) HttpRequest(req interface{}) (resp IResponseHandler) {
	pj.setHttpRequest(req)

	httpHandler, errHandler := pj.genHttpRequestHandler()
	if errHandler != nil {
		pj.setErr(errHandler)
		return
	}
	pj.setHttpRequestHandler(httpHandler)

	pageRequest, errPageRequest := pj.genPageRequest()
	if errPageRequest != nil {
		pj.setErr(errPageRequest)
		return
	}
	pj.setPageRequest(pageRequest)

	resp = pj
	return
}

func (pj *PaginateJson) ManualRequest(req PageRequest) (resp IResponseHandler) {
	pj.setPageRequest(req)

	resp = pj
	return
}

func (pj *PaginateJson) Response() (page Page) {
	// filtering
	resFilter, errFilter := pj.filtering()
	if errFilter != nil {
		pj.setErr(errFilter)
		return
	}
	pj.setRequestData(resFilter)

	// sort
	resSort, errSort := pj.sorting()
	if errSort != nil {
		pj.setErr(errSort)
		return
	}
	pj.setRequestData(resSort)

	// only show fields by fields
	resField, errField := pj.fielding()
	if errField != nil {
		pj.setErr(errField)
		return
	}
	pj.setRequestData(resField)

	// paging
	resPaging := pj.paging()
	pj.setResult(resPaging)

	// attach error
	resAttach := pj.attachErrToResult()
	pj.setResult(resAttach)

	page = resAttach
	return
}

func (pj PaginateJson) Err() (err error) {
	err = pj.getErr()
	return
}

func (pj PaginateJson) filtering() (result []interface{}, err error) {
	listIface := pj.getRequestData()
	pageRes := pj.getPageRequest()

	configRequest := pj.getConfigRequest()

	// filtering
	filters := pageRes.Filters
	resFilter, errFilter := Filtering(listIface, filters, configRequest)
	if errFilter != nil {
		err = errFilter
		return
	}

	result = resFilter
	return
}

func (pj PaginateJson) sorting() (result []interface{}, err error) {
	listIface := pj.getRequestData()
	pageRes := pj.getPageRequest()

	// sort
	sort := pageRes.Sort
	order := pageRes.Order
	resSort, errSort := Sorting(listIface, sort, order)
	if errSort != nil {
		err = errSort
		return
	}

	result = resSort
	return
}

func (pj PaginateJson) fielding() (result []interface{}, err error) {
	listIface := pj.getRequestData()
	pageRes := pj.getPageRequest()

	// only show fields by fields
	fields := pageRes.Fields
	resField, errField := Fielding(listIface, fields)
	if errField != nil {
		err = errField
		return
	}

	result = resField
	return
}

func (pj PaginateJson) paging() (result Page) {
	listIface := pj.getRequestData()
	pageRes := pj.getPageRequest()

	// offset by page and size
	currentPage := pageRes.Page
	size := pageRes.Size

	if lib.IsEmptyInt(size) {
		configRequest := pj.getConfigRequest()
		defaultSize := configRequest.DefaultSize()
		size = int(defaultSize)
	}

	_, pagination := OffsetByPageSize(listIface, currentPage, size)
	result = pagination
	return
}

func (pj PaginateJson) attachErrToResult() (result Page) {
	result = pj.getResult()

	err := pj.getErr()
	if err == nil {
		return
	}

	configReq := pj.getConfigRequest()

	errEnabled := configReq.ErrorEnabled()
	if !errEnabled {
		return
	}

	result.ErrorMessage = err.Error()
	return
}

func (pj *PaginateJson) setConfig(config Config) {
	pj.config = config
}

func (pj PaginateJson) getConfig() (config Config) {
	config = pj.config
	return
}

func (pnj *PaginateJson) setConfigRequest(configRequest IConfigRequest) {
	pnj.configRequest = configRequest
}

func (pnj PaginateJson) getConfigRequest() IConfigRequest {
	return pnj.configRequest
}

func (pnj PaginateJson) genConfigRequest() (configRequest IConfigRequest, err error) {
	config := pnj.getConfig()

	builderCfr := configrequest.BuilderConfigRequest{
		DefaultConjunction:   config.Operator.String(),
		DefaultSize:          config.DefaultSize,
		CustomParamEnabled:   config.CustomParamEnabled,
		SortParams:           config.SortParams,
		PageParams:           config.PageParams,
		OrderParams:          config.OrderParams,
		SizeParams:           config.SizeParams,
		FiltersParams:        config.FiltersParams,
		FieldSelectorEnabled: config.FieldSelectorEnabled,
		FieldsParams:         config.FieldsParams,
		// CacheAdapter:         config.CacheAdapter,
		// CacheCompression:     config.CacheCompression,
		// CachePrefix:          config.CachePrefix,
		JsonMarshal:   config.JSONMarshal,
		JsonUnmarshal: config.JSONUnmarshal,
		ErrorEnabled:  config.ErrorEnabled,
	}

	cfr := configrequest.NewConfigRequest(builderCfr)
	if cfr.Err() != nil {
		err = cfr.Err()
		return
	}

	configRequest = cfr
	return
}

func (pj *PaginateJson) setRawData(rawData interface{}) {
	pj.rawData = rawData
}

func (pj PaginateJson) getRawData() (rawData interface{}) {
	rawData = pj.rawData
	return
}

func (pj *PaginateJson) setRequestData(requestData []interface{}) {
	pj.requestData = requestData
}

func (pj PaginateJson) getRequestData() (requestData []interface{}) {
	requestData = pj.requestData
	return
}

func (pj PaginateJson) genRequestData() (requestData []interface{}, isValid bool) {
	rawData := pj.getRawData()

	reqData, ok := rawData.([]interface{})
	if !ok {
		return
	}

	isValid = true
	requestData = reqData
	return
}

func (pj *PaginateJson) setHttpRequest(httpRequest interface{}) {
	pj.httpRequest = httpRequest
}

func (pj PaginateJson) getHttpRequest() (httpRequest interface{}) {
	httpRequest = pj.httpRequest
	return
}

func (pj *PaginateJson) setHttpRequestHandler(httpRequestHandler *httprequest.HttpRequest) {
	pj.httpRequestHandler = httpRequestHandler
}

func (pj PaginateJson) getHttpRequestHandler() (httpRequestHandler *httprequest.HttpRequest) {
	httpRequestHandler = pj.httpRequestHandler
	return
}

func (pj PaginateJson) genHttpRequestHandler() (httpRequestHandler *httprequest.HttpRequest, err error) {
	configReq := pj.getConfigRequest()

	builderHrq := httprequest.BuilderHttpRequest{
		Config: configReq,
	}

	hrq := httprequest.NewHttpRequest(builderHrq)
	if hrq.Err() != nil {
		err = hrq.Err()
		return
	}

	httpRequestHandler = hrq
	return
}

func (pj *PaginateJson) setPageRequest(pageRequest PageRequest) {
	pj.pageRequest = pageRequest
}

func (pj PaginateJson) getPageRequest() (pageRequest PageRequest) {
	pageRequest = pj.pageRequest
	return
}

func (pj PaginateJson) genPageRequest() (pageRequest PageRequest, err error) {
	httpRawReq := pj.getHttpRequest()
	httpHandler := pj.getHttpRequestHandler()

	query := httpHandler.Query(httpRawReq)
	if httpHandler.Err() != nil {
		err = httpHandler.Err()
		return
	}

	// size
	sizeStr := query.Size()
	sizeInt, errSizeInt := strconv.Atoi(sizeStr)
	if errSizeInt != nil {
		sizeInt = 10
	}
	pageRequest.Size = sizeInt

	// page
	pageStr := query.Page()
	pageInt, errPageInt := strconv.Atoi(pageStr)
	if errPageInt != nil {
		pageInt = 0
	}
	pageRequest.Page = pageInt

	// sort
	sortStr := query.Sort()
	pageRequest.Sort = sortStr

	// order
	orderStr := query.Order()
	pageRequest.Order = orderStr

	// filters
	filtersStr := query.Filters()
	pageRequest.Filters = filtersStr

	// fields
	fieldsStr := query.Fields()
	pageRequest.Fields = fieldsStr

	return
}

func (pnj *PaginateJson) setResult(result Page) {
	pnj.result = result
}

func (pnj PaginateJson) getResult() Page {
	return pnj.result
}

func (pj *PaginateJson) setErr(err error) {
	pj.err = err
}

func (pj *PaginateJson) clearErr() {
	pj.err = nil
}

func (pj PaginateJson) getErr() (err error) {
	err = pj.err
	return
}
