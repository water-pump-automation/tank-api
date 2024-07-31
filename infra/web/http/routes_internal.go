package web

import (
	"context"
	"encoding/json"
	"net/http"
	"tank-api/app/entity/validation"
	"tank-api/app/usecases/create_tank"
	"tank-api/app/usecases/update_tank_state"
)

type InternalAPI struct {
	createTankUsecase *create_tank.CreateTank
	updateTankUsecase *update_tank_state.UpdateTank
}

func NewInternalAPI(
	createTankUsecase *create_tank.CreateTank,
	updateTankUsecase *update_tank_state.UpdateTank,
) *InternalAPI {
	return &InternalAPI{
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

	mux.HandleFunc("/v1/tank", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			createTank(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/tank/{tank_name}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPatch:
			updateTank(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func createTank(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	input, err := getBody(writer, request)
	if err != nil {
		return
	}

	response, err := api.createTankUsecase.Create(ctx, input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case create_tank.ErrTankAlreadyExists.Error(), create_tank.ErrTankInvalidGroup.Error(), create_tank.ErrTankMaximumCapacityZero.Error(), create_tank.ErrTankInvalidName.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(TankInvalidRequest, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusUnprocessableEntity)
		default:
			errorBody, err = getResponseBody(writer, NewAPIError(TankInternalServerError, err.Error()))
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

func updateTank(api *InternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	input, err := getBody(writer, request)
	if err != nil {
		return
	}

	input["group"] = request.Header.Get("group")
	input["tank_name"] = request.PathValue("tank_name")

	err = api.updateTankUsecase.Update(ctx, input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case update_tank_state.ErrTankCurrentLevelBiggerThanMax.Error(), update_tank_state.ErrTankCurrentLevelSmallerThanZero.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(TankInvalidRequest, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusUnprocessableEntity)
		case update_tank_state.ErrTankErrorNotFound.Error():
			errorBody, err = getResponseBody(writer, NewAPIError(TankNotFound, err.Error()))
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusNotFound)
		default:
			errorBody, err = getResponseBody(writer, NewAPIError(TankInternalServerError, err.Error()))
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
