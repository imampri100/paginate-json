package httprequest

type IQuery interface {
	Size() (size string)
	Page() (page string)
	Sort() (sort string)
	Order() (order string)
	Filters() (filters string)
	Fields() (fields string)
}
