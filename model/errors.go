package model

type ErrorDetail struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type Errors struct {
	ErrorCode    string                 `json:"code"`
	ErrorMessage string                 `json:"message"`
	Items        map[string]ErrorDetail `json:"items"`
}
