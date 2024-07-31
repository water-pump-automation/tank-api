package web

import (
	"context"
	"encoding/json"
	"net/http"
	"tank-api/app/entity/validation"
	"tank-api/app/usecases/get_group"
	"tank-api/app/usecases/get_tank"
)

type ExternalAPI struct {
	getTankUsecase  *get_tank.GetTank
	getGroupUsecase *get_group.GetGroupTank
}

func NewExternalAPI(
	getTankUsecase *get_tank.GetTank,
	getGroupUsecase *get_group.GetGroupTank,
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

	mux.HandleFunc("/v1/tank/{tank_name}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getTank(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/group/{group}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroup(api, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func getTank(api *ExternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	input := map[string]interface{}{
		"tank_name": request.PathValue("tank_name"),
		"group":     request.Header.Get("group"),
	}

	response, err := api.getTankUsecase.Get(ctx, input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case get_tank.ErrTankErrorNotFound.Error():
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

	responseBytes, err := getResponseBody(writer, response)
	if err != nil {
		return
	}
	writer.Write(responseBytes)
	writer.WriteHeader(http.StatusOK)
}

func getGroup(api *ExternalAPI, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	input := map[string]interface{}{
		"group": request.Header.Get("group"),
	}

	response, err := api.getGroupUsecase.Get(ctx, input)

	if _, ok := err.(validation.ValidationError); ok {
		writeBadRequestError(writer, err)
		return
	}

	if err != nil {
		var errorBody []byte

		switch err.Error() {
		case get_group.ErrTankErrorGroupNotFound.Error():
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

	responseBytes, err := getResponseBody(writer, response)
	if err != nil {
		return
	}
	writer.Write(responseBytes)
	writer.WriteHeader(http.StatusOK)
}
