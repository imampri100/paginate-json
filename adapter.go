package paginatejson

type IConfigRequest interface {
	DefaultConjunction() string
	DefaultSize() int64
	SortParams() []string
	PageParams() []string
	OrderParams() []string
	SizeParams() []string
	FiltersParams() []string
	FieldsParams() []string
	JsonMarshal() func(v interface{}) ([]byte, error)
	JsonUnmarshal() func(data []byte, v interface{}) error
	ErrorEnabled() bool
}
