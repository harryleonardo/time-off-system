package vo

// SuccessResponse ...
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    int         `json:"status"`
	ErrorCode string      `json:"error_code"`
}
