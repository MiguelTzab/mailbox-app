package entity

type SearchResult struct {
	Items        []Email `json:"items"`
	TotalResults int     `json:"total_results"`
	Page         int     `json:"page"`
	PageSize     int     `json:"page_size"`
}
