package response

// Empty used to return nothing
type Empty struct{}

// ErrorResponse is struct used to return error message to the client
type ErrorResponse struct {
	ErrorCode string   `json:"error_code"`
	Message   string   `json:"message"`
	Fields    []string `json:"fields"`
}

func BuildErrorResponse(errorCode, message string, fields []string) ErrorResponse {
	fieldsValue := make([]string, len(fields))
	copy(fieldsValue, fields)

	return ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
		Fields:    fieldsValue,
	}
}

// SuccessResponse is struct used to return success message to the client
type SuccessResponse struct {
	Data any `json:"data"`
}

// DataTableResponse is struct used to return success message to the client
type DataTableResponse struct {
	Message  string            `json:"message"`
	Data     any               `json:"data"`
	MetaData DataTableMetaData `json:"meta_data"`
}

type DataTableMetaData struct {
	CurrentPage int32 `json:"current_page"`
	PerPage     int32 `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
}
