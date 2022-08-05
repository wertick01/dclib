package models

type ErrorModel struct {
	ErrorCode    string `json:"error_number"`
	ErrorDetails string `json:"error_ticker"`
	ErrorValue   string `json:"error_value"`
}
