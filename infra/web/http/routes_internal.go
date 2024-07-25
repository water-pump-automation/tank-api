package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"water-tank-api/app/entity/logs"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/entity/water_tank"
	"water-tank-api/app/usecases/create_tank"
	"water-tank-api/app/usecases/get_group"
	"water-tank-api/app/usecases/get_tank"
	"water-tank-api/app/usecases/ports"
	"water-tank-api/app/usecases/update_tank_state"
)

type InternalAPI struct {
	getTankUsecase    *get_tank.GetWaterTank
	getGroupUsecase   *get_group.GetGroupWaterTank
	createTankUsecase *create_tank.CreateWaterTank
	updateTankUsecase *update_tank_state.UpdateWaterTank
}

func NewInternalAPI(
	getTankUsecase *get_tank.GetWaterTank,
	getGroupUsecase *get_group.GetGroupWaterTank,
	createTankUsecase *create_tank.CreateWaterTank,
	updateTankUsecase *update_tank_state.UpdateWaterTank,
) *InternalAPI {
	return &InternalAPI{
		getTankUsecase:    getTankUsecase,
		getGroupUsecase:   getGroupUsecase,
		createTankUsecase: createTankUsecase,
		updateTankUsecase: updateTankUsecase,
	}
}

func (api *InternalAPI) Route(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		response, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			writer.Write([]byte("Internal server error"))
		}
		writer.Write(response)
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/v1/water-tank", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			createTankInternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/tank/{tank}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getTankInternal(api, writer, request)
		case http.MethodPatch:
			updateTankInternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/group/{group}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupInternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/v1/water-tank/group", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupInternal(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

func createTankInternal(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
	var input water_tank.CreateInput

	ctx := context.Background()

	if err := getBody(writer, request, &input); err != nil {
		return
	}

	logs.Gateway().Info(
		fmt.Sprintf("Creating '%s' tank for group '%s' with %s capacity...",
			input.TankName, input.Group, ports.ConvertCapacityToLiters(input.MaximumCapacity)),
	)

	response, err := api.createTankUsecase.Create(ctx, nil, &input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case create_tank.ErrWaterTankAlreadyExists.Error(), create_tank.ErrWaterTankInvalidGroup.Error(), create_tank.ErrWaterTankMaximumCapacityZero.Error(), create_tank.ErrWaterTankInvalidName.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankInvalidRequest, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusUnprocessableEntity)
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
	writer.WriteHeader(http.StatusCreated)
}

func updateTankInternal(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
	var input water_tank.UpdateWaterLevelInput

	ctx := context.Background()

	tankName := request.PathValue("tank")
	group := request.Header.Get("group")

	if err := getBody(writer, request, &input); err != nil {
		return
	}

	input.Group = group
	input.TankName = tankName

	logs.Gateway().Info(
		fmt.Sprintf("Updating '%s' tank's, of group '%s', water level to %s",
			input.TankName, input.Group, ports.ConvertCapacityToLiters(input.NewWaterLevel)),
	)

	err := api.updateTankUsecase.Update(ctx, nil, &input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case update_tank_state.ErrWaterTankCurrentWaterLevelBiggerThanMax.Error(), update_tank_state.ErrWaterTankCurrentWaterLevelSmallerThanZero.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(WaterTankInvalidRequest, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusUnprocessableEntity)
		case update_tank_state.ErrWaterTankErrorNotFound.Error():
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

	writer.WriteHeader(http.StatusNoContent)
}

func getTankInternal(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
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

func getGroupInternal(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
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
