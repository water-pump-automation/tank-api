package web

import (
	"net/http"
	"water-tank-api/app/controllers/response"
)

func mapError(code string) int {
	switch code {
	case response.WaterTankBadRequest:
		return http.StatusBadRequest
	case response.WaterTankInvalidRequest:
		return http.StatusUnprocessableEntity
	case response.WaterTankNotFound:
		return http.StatusNotFound
	case response.WaterTankInternalServerError:
		return http.StatusInternalServerError
	case response.WaterTankOK:
		return http.StatusNoContent
	case response.WaterTankNoContent:
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}
