package response

import (
	"encoding/json"
	"errors"
	"water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/usecases/ports"
)

var (
	WaterTankOK                  = "WATERTANK_200"
	WaterTankNoContent           = "WATERTANK_204"
	WaterTankNotFound            = "WATERTANK_404"
	WaterTankBadRequest          = "WATERTANK_400"
	WaterTankInvalidRequest      = "WATERTANK_422"
	WaterTankInternalServerError = "WATERTANK_500"
)

var (
	ErrWaterTankEmptyNameError = errors.New("bad request. name cannot be empty")
)

type ControllerResponse struct {
	Code    string                 `json:"code"`
	Content map[string]interface{} `json:"content"`
}

func NewControllerResponse(code string, content *ports.WaterTankState) *ControllerResponse {
	bytes, _ := json.Marshal(content)

	m := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &m)

	return &ControllerResponse{
		Content: m,
		Code:    code,
	}
}

func NewControllerEmptyResponse(code string) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{},
		Code:    code,
	}
}

func NewControllerGroupResponse(code string, content *ports.WaterTankGroupState) *ControllerResponse {
	bytes, _ := json.Marshal(content)

	m := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &m)

	return &ControllerResponse{
		Content: m,
		Code:    code,
	}
}

func NewControllerError(code string, message string) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{"error": message},
		Code:    code,
	}
}

func SwitchError(usecaseErr error_stack.Error) (response *ControllerResponse) {
	switch usecaseErr.EntityError() {
	case nil:
		response = NewControllerError(WaterTankNotFound, usecaseErr.LastUsecaseError().Error())
	default:
		response = NewControllerError(WaterTankInternalServerError, usecaseErr.LastUsecaseError().Error())
	}

	err := usecaseErr.LastUsecaseError()
	logs.Gateway().Error(err.Error())
	return
}

func NewValidationError() *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{},
		Code:    WaterTankBadRequest,
	}
}

func (response *ControllerResponse) AddDetails(field, value string) {
	response.Content[field] = value
}
