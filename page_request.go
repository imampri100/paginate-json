package paginatejson

// PageRequest struct
type PageRequest struct {
	Size    int
	Page    int
	Sort    string
	Order   string
	Filters string
	Fields  string
}
