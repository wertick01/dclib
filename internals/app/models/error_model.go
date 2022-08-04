package models

type ErrorModel struct {
	ErrorNumber string `json:"error_number"`
	ErrorTicker string `json:"error_ticker"`
	ErrorValue  string `json:"error_value"`
}
