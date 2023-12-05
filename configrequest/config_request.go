package configrequest

import (
	"encoding/json"
	"paginatejson/lib"
	"strings"
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

type DefaultConjunction string

const (
	Or  DefaultConjunction = "or"
	And DefaultConjunction = "and"
)

func (dfo DefaultConjunction) String() string {
	return string(dfo)
}

const (
	DefaultSize int64 = 10
)

type DefaultParams string

const (
	SizeParamJson    DefaultParams = "size"
	PageParamJson    DefaultParams = "page"
	SortParamJson    DefaultParams = "sort"
	OrderParamJson   DefaultParams = "order"
	FiltersParamJson DefaultParams = "filters"
	FieldsParamJson  DefaultParams = "fields"
)

func (pjf DefaultParams) String() string {
	return string(pjf)
}

const (
	DefaultCachePrefix string = "paginate_json"
)

type BuilderConfigRequest struct {
	DefaultConjunction   string
	DefaultSize          int64
	CustomParamEnabled   bool
	SortParams           []string
	PageParams           []string
	OrderParams          []string
	SizeParams           []string
	FiltersParams        []string
	FieldSelectorEnabled bool
	FieldsParams         []string
	CacheAdapter         ICacheAdapter
	CacheCompression     bool
	CachePrefix          string
	JsonMarshal          func(v interface{}) ([]byte, error)
	JsonUnmarshal        func(data []byte, v interface{}) error
	ErrorEnabled         bool
}

// config for customize pagination result
type ConfigRequest struct {
	builder              BuilderConfigRequest
	defaultConjunction   string   // default: or
	defaultSize          int64    // default: 10
	customParamEnabled   bool     // default: false
	sortParams           []string // default: []string{"sort"}
	pageParams           []string // default: []string{"page"}
	orderParams          []string // default: []string{"order"}
	sizeParams           []string // default: []string{"size"}
	filtersParams        []string // default: []string{"filters"}
	fieldSelectorEnabled bool     // default: false
	fieldsParams         []string // default: []string{"fields"}
	// cacheAdapter         ICacheAdapter // required
	// cacheCompression     bool          // default: false
	// cachePrefix          string        // default: paginate_json
	jsonMarshal   func(v interface{}) ([]byte, error)
	jsonUnmarshal func(data []byte, v interface{}) error
	errorEnabled  bool // default: false
	err           error
}

// func (ciq ConfigRequest) CachePrefix() string {
// 	return ciq.getCachePrefix()
// }

// func (ciq *ConfigRequest) setCachePrefix(cachePrefix string) {
// 	ciq.cachePrefix = cachePrefix
// }

// func (ciq ConfigRequest) getCachePrefix() string {
// 	return ciq.cachePrefix
// }

// func (ciq ConfigRequest) genCachePrefix() (cachePrefix string) {
// 	builder := ciq.getBuilder()
// 	bCachePrefix := builder.CachePrefix

// 	if lib.IsEmptyStr(bCachePrefix) {
// 		cachePrefix = DefaultCachePrefix
// 		return
// 	}

// 	cachePrefix = bCachePrefix
// 	return
// }

func (ciq ConfigRequest) ErrorEnabled() bool {
	return ciq.getErrorEnabled()
}

func (ciq *ConfigRequest) setErrorEnabled(errorEnabled bool) {
	ciq.errorEnabled = errorEnabled
}

func (ciq ConfigRequest) getErrorEnabled() bool {
	return ciq.errorEnabled
}

func (ciq ConfigRequest) genErrorEnabled() (errorEnabled bool) {
	builder := ciq.getBuilder()
	bErrorEnabled := builder.ErrorEnabled
	errorEnabled = bErrorEnabled
	return
}

func (ciq ConfigRequest) Err() error {
	return ciq.getErr()
}

func (ciq *ConfigRequest) setErr(err error) {
	ciq.err = err
}

func (ciq ConfigRequest) getErr() error {
	return ciq.err
}

func NewConfigRequest(builder BuilderConfigRequest) (cfr *ConfigRequest) {
	cfr = new(ConfigRequest)

	// setBuilder
	cfr.setBuilder(builder)

	// setConjunction
	conjunction := cfr.genDefaultConjunction()
	cfr.setDefaultConjunction(conjunction)

	// setDefaultSize
	defaultSize := cfr.genDefaultSize()
	cfr.setDefaultSize(defaultSize)

	// setCustomParamEnabled
	customParamEnabled := cfr.genCustomParamEnabled()
	cfr.setCustomParamEnabled(customParamEnabled)

	// setSortParams
	sortParams := cfr.genSortParams()
	cfr.setSortParams(sortParams)

	// setPageParams
	pageParams := cfr.genPageParams()
	cfr.setPageParams(pageParams)

	// setOrderParams
	orderParams := cfr.genOrderParams()
	cfr.setOrderParams(orderParams)

	// setSizeParams
	sizeParams := cfr.genSizeParams()
	cfr.setSizeParams(sizeParams)

	// setFiltersParams
	filtersParams := cfr.genFiltersParams()
	cfr.setFiltersParams(filtersParams)

	// setFieldSelectorEnabled
	fieldSelectorEnabled := cfr.genFieldSelectorEnabled()
	cfr.setFieldSelectorEnabled(fieldSelectorEnabled)

	// setFieldsParams
	fieldsParams := cfr.genFieldsParams()
	cfr.setFieldsParams(fieldsParams)

	// // setCacheAdapter
	// cacheAdapter, errCacheAdapter := cfr.genCacheAdapter()
	// if errCacheAdapter != nil {
	// 	cfr.setErr(errCacheAdapter)
	// 	return
	// }
	// cfr.setCacheAdapter(cacheAdapter)

	// // setCacheCompression
	// cacheCompression := cfr.genCacheCompression()
	// cfr.setCacheCompression(cacheCompression)

	// // setCachePrefix
	// cachePrefix := cfr.genCachePrefix()
	// cfr.setCachePrefix(cachePrefix)

	// setJsonMarshal
	jsonMarshal := cfr.genJsonMarshal()
	cfr.setJsonMarshal(jsonMarshal)

	// setJsonUnmarshal
	jsonUnmarshal := cfr.genJsonUnmarshal()
	cfr.setJsonUnmarshal(jsonUnmarshal)

	// setErrorEnabled
	errorEnabled := cfr.genErrorEnabled()
	cfr.setErrorEnabled(errorEnabled)

	return
}

func (cfr ConfigRequest) DefaultConjunction() string {
	return cfr.getDefaultConjunction()
}

func (cfr ConfigRequest) DefaultSize() int64 {
	return cfr.getDefaultSize()
}

func (cfr ConfigRequest) CustomParamEnabled() bool {
	return cfr.getCustomParamEnabled()
}

func (cfr ConfigRequest) SortParams() []string {
	return cfr.getSortParams()
}

func (cfr ConfigRequest) PageParams() []string {
	return cfr.getPageParams()
}

func (cfr ConfigRequest) OrderParams() []string {
	return cfr.getOrderParams()
}

func (cfr ConfigRequest) SizeParams() []string {
	return cfr.getSizeParams()
}

func (cfr ConfigRequest) FiltersParams() []string {
	return cfr.getFiltersParams()
}

func (cfr ConfigRequest) FieldsParams() []string {
	return cfr.getFieldsParams()
}

func (cfr ConfigRequest) FieldSelectorEnabled() bool {
	return cfr.getFieldSelectorEnabled()
}

// func (cfr ConfigRequest) CacheAdapter() ICacheAdapter {
// 	return cfr.getCacheAdapter()
// }

// func (cfr ConfigRequest) CacheCompression() bool {
// 	return cfr.getCacheCompression()
// }

func (cfr ConfigRequest) JsonMarshal() func(v interface{}) ([]byte, error) {
	return cfr.getJsonMarshal()
}

func (cfr ConfigRequest) JsonUnmarshal() func(data []byte, v interface{}) error {
	return cfr.getJsonUnmarshal()
}

func (ciq *ConfigRequest) setBuilder(builder BuilderConfigRequest) {
	ciq.builder = builder
}

func (ciq ConfigRequest) getBuilder() BuilderConfigRequest {
	return ciq.builder
}

func (cfr *ConfigRequest) setDefaultConjunction(conjunction string) {
	cfr.defaultConjunction = conjunction
}

func (cfr *ConfigRequest) setDefaultSize(defaultSize int64) {
	cfr.defaultSize = defaultSize
}

func (cfr *ConfigRequest) setCustomParamEnabled(customParamEnabled bool) {
	cfr.customParamEnabled = customParamEnabled
}

func (cfr *ConfigRequest) setSortParams(sortParams []string) {
	cfr.sortParams = sortParams
}

func (cfr *ConfigRequest) setPageParams(pageParams []string) {
	cfr.pageParams = pageParams
}

func (cfr *ConfigRequest) setOrderParams(orderParams []string) {
	cfr.orderParams = orderParams
}

func (cfr *ConfigRequest) setSizeParams(sizeParams []string) {
	cfr.sizeParams = sizeParams
}

func (cfr *ConfigRequest) setFiltersParams(filtersParams []string) {
	cfr.filtersParams = filtersParams
}

func (cfr *ConfigRequest) setFieldsParams(fieldsParams []string) {
	cfr.fieldsParams = fieldsParams
}

func (cfr *ConfigRequest) setFieldSelectorEnabled(fieldSelectorEnabled bool) {
	cfr.fieldSelectorEnabled = fieldSelectorEnabled
}

// func (cfr *ConfigRequest) setCacheAdapter(cacheAdapter ICacheAdapter) {
// 	cfr.cacheAdapter = cacheAdapter
// }

// func (cfr *ConfigRequest) setCacheCompression(cacheCompression bool) {
// 	cfr.cacheCompression = cacheCompression
// }

func (cfr *ConfigRequest) setJsonMarshal(jsonMarshal func(v interface{}) ([]byte, error)) {
	cfr.jsonMarshal = jsonMarshal
}

func (cfr *ConfigRequest) setJsonUnmarshal(jsonUnmarshal func(data []byte, v interface{}) error) {
	cfr.jsonUnmarshal = jsonUnmarshal
}

func (cfr ConfigRequest) getDefaultConjunction() string {
	return cfr.defaultConjunction
}

func (cfr ConfigRequest) genDefaultConjunction() (conjunction string) {
	builder := cfr.getBuilder()
	bDefaultConjunction := builder.DefaultConjunction

	switch {
	case strings.EqualFold(bDefaultConjunction, "or"):
		{
			conjunction = "or"
			break
		}
	case strings.EqualFold(bDefaultConjunction, "and"):
		{
			conjunction = "and"
			break
		}
	default:
		{
			conjunction = "or"
			break
		}
	}

	return
}

func (cfr ConfigRequest) getDefaultSize() int64 {
	return cfr.defaultSize
}

func (cfr ConfigRequest) genDefaultSize() (defaultSize int64) {
	builder := cfr.getBuilder()
	bSize := builder.DefaultSize

	if lib.IsEmptyInt64(bSize) {
		defaultSize = DefaultSize
		return
	}

	defaultSize = bSize
	return
}

func (cfr ConfigRequest) getCustomParamEnabled() bool {
	return cfr.customParamEnabled
}

func (cfr ConfigRequest) genCustomParamEnabled() (customParamEnabled bool) {
	builder := cfr.getBuilder()
	bCustomParamEnabled := builder.CustomParamEnabled
	customParamEnabled = bCustomParamEnabled
	return
}

func (cfr ConfigRequest) getSortParams() []string {
	return cfr.sortParams
}

func (cfr ConfigRequest) genSortParams() (sortParams []string) {
	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		sortParams = append(sortParams, SortParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bSortParams := builder.SortParams

	if len(bSortParams) == 0 {
		sortParams = append(sortParams, SortParamJson.String())
		return
	}

	sortParams = append(sortParams, bSortParams...)
	return
}

func (cfr ConfigRequest) getPageParams() []string {
	return cfr.pageParams
}

func (cfr ConfigRequest) genPageParams() (pageParams []string) {
	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		pageParams = append(pageParams, PageParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bPageParams := builder.PageParams

	if len(bPageParams) == 0 {
		pageParams = append(pageParams, PageParamJson.String())
		return
	}

	pageParams = append(pageParams, bPageParams...)
	return
}

func (cfr ConfigRequest) getOrderParams() []string {
	return cfr.orderParams
}

func (cfr ConfigRequest) genOrderParams() (orderParams []string) {
	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		orderParams = append(orderParams, OrderParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bOrderParams := builder.OrderParams

	if len(bOrderParams) == 0 {
		orderParams = append(orderParams, OrderParamJson.String())
		return
	}

	orderParams = append(orderParams, bOrderParams...)
	return
}

func (cfr ConfigRequest) getSizeParams() []string {
	return cfr.sizeParams
}

func (cfr ConfigRequest) genSizeParams() (sizeParams []string) {
	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		sizeParams = append(sizeParams, SizeParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bSizeParams := builder.SizeParams

	if len(bSizeParams) == 0 {
		sizeParams = append(sizeParams, SizeParamJson.String())
		return
	}

	sizeParams = append(sizeParams, bSizeParams...)
	return
}

func (cfr ConfigRequest) getFiltersParams() []string {
	return cfr.filtersParams
}

func (cfr ConfigRequest) genFiltersParams() (filtersParams []string) {
	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		filtersParams = append(filtersParams, FiltersParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bFiltersParams := builder.FiltersParams
	if len(bFiltersParams) == 0 {
		filtersParams = append(filtersParams, FiltersParamJson.String())
		return
	}

	filtersParams = append(filtersParams, bFiltersParams...)
	return
}

func (cfr ConfigRequest) getFieldsParams() []string {
	return cfr.fieldsParams
}

func (cfr ConfigRequest) genFieldsParams() (fieldsParams []string) {
	fieldSelectorEnabled := cfr.getFieldSelectorEnabled()
	if !fieldSelectorEnabled {
		return
	}

	customParamEnabled := cfr.getCustomParamEnabled()
	if !customParamEnabled {
		fieldsParams = append(fieldsParams, FieldsParamJson.String())
		return
	}

	builder := cfr.getBuilder()
	bFieldsParams := builder.FieldsParams

	if len(bFieldsParams) == 0 {
		fieldsParams = append(fieldsParams, FieldsParamJson.String())
		return
	}

	fieldsParams = append(fieldsParams, bFieldsParams...)
	return
}

func (cfr ConfigRequest) getFieldSelectorEnabled() bool {
	return cfr.fieldSelectorEnabled
}

func (cfr ConfigRequest) genFieldSelectorEnabled() (fieldSelectorEnabled bool) {
	builder := cfr.getBuilder()
	bFieldSelectorEnabled := builder.FieldSelectorEnabled
	fieldSelectorEnabled = bFieldSelectorEnabled
	return
}

// func (cfr ConfigRequest) getCacheAdapter() ICacheAdapter {
// 	return cfr.cacheAdapter
// }

// func (cfr ConfigRequest) genCacheAdapter() (cacheAdapter ICacheAdapter, err error) {
// 	builder := cfr.getBuilder()
// 	bCacheAdapter := builder.CacheAdapter

// 	if bCacheAdapter == nil {
// 		err = errors.New("cache adapter is required")
// 		return
// 	}

// 	cacheAdapter = bCacheAdapter
// 	return
// }

// func (cfr ConfigRequest) getCacheCompression() bool {
// 	return cfr.cacheCompression
// }

// func (cfr ConfigRequest) genCacheCompression() (cacheCompression bool) {
// 	builder := cfr.getBuilder()
// 	bCacheCompression := builder.CacheCompression
// 	cacheCompression = bCacheCompression
// 	return
// }

func (cfr ConfigRequest) getJsonMarshal() func(v interface{}) ([]byte, error) {
	return cfr.jsonMarshal
}

func (cfr ConfigRequest) genJsonMarshal() (jsonMarshal func(v interface{}) ([]byte, error)) {
	builder := cfr.getBuilder()
	bJsonMarshal := builder.JsonMarshal

	if bJsonMarshal == nil {
		jsonMarshal = json.Marshal
		return
	}

	jsonMarshal = bJsonMarshal
	return
}

func (cfr ConfigRequest) getJsonUnmarshal() func(data []byte, v interface{}) error {
	return cfr.jsonUnmarshal
}

func (cfr ConfigRequest) genJsonUnmarshal() (jsonUnmarshal func(data []byte, v interface{}) error) {
	builder := cfr.getBuilder()
	bJsonUnmarshal := builder.JsonUnmarshal

	if bJsonUnmarshal == nil {
		jsonUnmarshal = json.Unmarshal
		return
	}

	jsonUnmarshal = bJsonUnmarshal
	return
}
