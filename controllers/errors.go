package controllers

import (
	"encoding/json"
	"errors"
	"water-tank-api/core/entity/access"
	"water-tank-api/core/usecases/get"
)

var (
	WaterTankOK                  = "WATERTANK_200"
	WaterTankNoContent           = "NETSTAT_204"
	WaterTankNotFound            = "WATERTANK_404"
	WaterTankBadRequest          = "WATERTANK_400"
	WaterTankInvalidRequest      = "WATERTANK_422"
	WaterTankInternalServerError = "WATERTANK_500"
)

var (
	WaterTankEmptyNameError = errors.New("Bad request. Name cannot be empty")
)

type ControllerResponse struct {
	Code    string                 `json:"code"`
	Content map[string]interface{} `json:"content"`
}

func NewControllerResponse(code string, content *get.WaterTankState) *ControllerResponse {
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

func NewControllerCreateResponse(code string, content access.AccessToken) *ControllerResponse {
	return &ControllerResponse{
		Content: map[string]interface{}{"access_token": content},
		Code:    code,
	}
}

func NewControllerGroupResponse(code string, content *get.WaterTankGroupState) *ControllerResponse {
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
