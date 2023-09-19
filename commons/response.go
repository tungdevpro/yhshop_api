package commons

type successResponse struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data"`
	Paging     interface{} `json:"paging,omitempty"`
	Filter     interface{} `json:"filter,omitempty"`
}

func CreateNewSuccessResp(data interface{}, msg string) *successResponse {
	return &successResponse{
		Paging:     nil,
		Filter:     nil,
		StatusCode: Ok,
		Message:    msg,
		Data:       data,
	}
}

func SimpleSuccessResp(data interface{}) *successResponse {
	return NewSuccessResp(data, nil, nil)
}

func NewSuccessResp(data interface{}, paging interface{}, filter interface{}) *successResponse {
	return &successResponse{
		StatusCode: Ok,
		Message:    "Success",
		Data:       data,
		Paging:     paging,
		Filter:     filter,
	}
}

type AppError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
}

func NewAppError(code int, msg string) *AppError {
	return &AppError{
		StatusCode: code,
		Message:    msg,
	}
}
