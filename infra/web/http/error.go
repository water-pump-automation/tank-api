package web

import "encoding/json"

var (
	WaterTankNotFound            = "WATERTANK_404"
	WaterTankBadRequest          = "WATERTANK_400"
	WaterTankInvalidRequest      = "WATERTANK_422"
	WaterTankInternalServerError = "WATERTANK_500"
)

type APIError struct {
	Code  string                 `json:"code"`
	Error map[string]interface{} `json:"error"`
}

func NewAPIError(code string, message string) *APIError {
	return &APIError{
		Error: map[string]interface{}{"error": message},
		Code:  code,
	}
}

func NewAPIValidationError(code string, message string) *APIError {
	validationErr := make(map[string]interface{})

	json.Unmarshal([]byte(message), &validationErr)
	return &APIError{
		Error: validationErr,
		Code:  code,
	}
}
