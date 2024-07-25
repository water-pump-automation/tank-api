package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/get_group"
	"water-tank-api/app/usecases/get_tank"
)

type ExternalAPI struct {
	getTankUsecase  *get_tank.GetWaterTank
	getGroupUsecase *get_group.GetGroupWaterTank
}

func NewExternalAPI(
	getTankUsecase *get_tank.GetWaterTank,
	getGroupUsecase *get_group.GetGroupWaterTank,
) *ExternalAPI {
	return &ExternalAPI{
		getTankUsecase:  getTankUsecase,
		getGroupUsecase: getGroupUsecase,
	}
}

func (api *ExternalAPI) Route(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		response, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			writeInternalServerError(writer)
			return
		}
		writer.Write(response)
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/v1/water-tank/tank/{tank}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getTankExternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/group/{group}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupExternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/v1/water-tank/group", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupExternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

func getTankExternal(api *ExternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	tankName := request.PathValue("tank")
	groupName := request.Header.Get("group")

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank state, of group '%s'...", tankName, groupName))

	response, err := api.getTankUsecase.Get(ctx, nil, &water_tank.GetWaterTankState{
		TankName: tankName,
		Group:    groupName,
	})

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case get_tank.ErrWaterTankErrorNotFound.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankNotFound, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusNotFound)
		default:
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankInternalServerError, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.Write(errorBody)
		return
	}

	responseBytes, err := getResponseBody(writer, response)
	if err != nil {
		return
	}
	writer.Write(responseBytes)
	writer.WriteHeader(http.StatusOK)
}

func getGroupExternal(api *ExternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	groupName := request.PathValue("group")

	logs.Gateway().Info(fmt.Sprintf("Retrieving '%s' tank group...", groupName))

	response, err := api.getGroupUsecase.Get(ctx, nil, &water_tank.GetGroupTanks{
		Group: groupName,
	})

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case get_group.ErrWaterTankErrorGroupNotFound.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankNotFound, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusNotFound)
		default:
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankInternalServerError, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.Write(errorBody)
		return
	}

	responseBytes, err := getResponseBody(writer, response)
	if err != nil {
		return
	}
	writer.Write(responseBytes)
	writer.WriteHeader(http.StatusOK)
}
