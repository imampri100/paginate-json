package httprequest

type IConfigRequest interface {
	SortParams() []string
	PageParams() []string
	OrderParams() []string
	SizeParams() []string
	FiltersParams() []string
	FieldsParams() []string
}
