package entity

type SearchQuery struct {
	Term     string
	Fields   map[string]string
	Page     int
	PageSize int
}
