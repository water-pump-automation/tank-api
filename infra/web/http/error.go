package web

var (
	WaterTankNotFound            = "WATERTANK_404"
	WaterTankBadRequest          = "WATERTANK_400"
	WaterTankInvalidRequest      = "WATERTANK_422"
	WaterTankInternalServerError = "WATERTANK_500"
)

type APIError struct {
	Code    string                 `json:"code"`
	Content map[string]interface{} `json:"content"`
}

func NewAPIError(code string, message string) *APIError {
	return &APIError{
		Content: map[string]interface{}{"error": message},
		Code:    code,
	}
}
