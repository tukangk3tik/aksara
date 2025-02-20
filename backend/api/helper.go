package api

type Pagination struct {
	Page  int `form:"page,default=1"`
	Limit int `form:"limit,default=1"`
}

// Empty used to return nothing
type Empty struct{}

// ErrorResponse is struct used to return error message to the client
type ErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

// SuccessResponse is struct used to return success message to the client
type SuccessResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}