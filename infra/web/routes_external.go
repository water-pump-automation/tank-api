package web

import (
	"context"
	"encoding/json"
	"net/http"
	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/water_tank"
)

type ExternalRouter struct{}

func (r *ExternalRouter) Route(mux *http.ServeMux, controller *controllers.ExternalController) {
	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		response, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			internalServerError(writer)
			return
		}
		writer.Write(response)
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/v1/water-tank/tank/{tank}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getTankExternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/group/{group}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupExternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/v1/water-tank/group", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupExternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

func getTankExternal(controller *controllers.ExternalController, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	tankName := request.PathValue("tank")
	groupName := request.Header.Get("group")

	resp, _ := controller.Get(ctx, nil, &water_tank.GetWaterTankState{
		TankName: tankName,
		Group:    groupName,
	})

	writer.WriteHeader(mapError(resp.Code))

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		internalServerError(writer)
		return
	}

	writer.Write(responseBytes)
}

func getGroupExternal(controller *controllers.ExternalController, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	groupName := request.PathValue("group")

	resp, _ := controller.GetGroup(ctx, nil, &water_tank.GetGroupTanks{
		Group: groupName,
	})

	writer.WriteHeader(mapError(resp.Code))

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		internalServerError(writer)
		return
	}

	writer.Write(responseBytes)
}
