package types

type PaginationMetadata struct {
	// The current page of results
	Page int `json:"page" example:"0"`
	// The total amount of pages
	PageCount int `json:"pageCount" example:"3"`
	// The amount of results per page
	PageSize int `json:"pageSize" example:"1000"`
	// The total number of results
	Total int `json:"total" example:"1024"`
}
