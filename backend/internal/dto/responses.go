package dto

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Pagination PaginationMeta `json:"pagination"`
}

type PaginationMeta struct {
	Page        int   `json:"page"`
	PageSize    int   `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int64 `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}
