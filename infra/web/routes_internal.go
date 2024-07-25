package web

import (
	"context"
	"encoding/json"
	"net/http"
	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/water_tank"
)

type InternalRouter struct{}

func (r *InternalRouter) Route(mux *http.ServeMux, controller *controllers.InternalController) {
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
			createTankInternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/tank/{tank}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getTankInternal(controller, writer, request)
		case http.MethodPatch:
			updateTankInternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/water-tank/group/{group}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupInternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/v1/water-tank/group", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getGroupInternal(controller, writer, request)
		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}

func createTankInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	var input water_tank.CreateInput

	ctx := context.Background()

	if err := getBody(writer, request, &input); err != nil {
		return
	}

	resp, _ := controller.Create(ctx, nil, &input)

	writer.WriteHeader(mapError(resp.Code))

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		internalServerError(writer)
		return
	}
	writer.Write(responseBytes)
}

func updateTankInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	var input water_tank.UpdateWaterLevelInput

	ctx := context.Background()

	tankName := request.PathValue("tank")
	group := request.Header.Get("group")

	if err := getBody(writer, request, &input); err != nil {
		return
	}

	input.Group = group
	input.TankName = tankName

	resp, _ := controller.Update(ctx, nil, &input)

	writer.WriteHeader(mapError(resp.Code))

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		internalServerError(writer)
		return
	}
	writer.Write(responseBytes)
}

func getTankInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
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
		writer.Write([]byte("Internal server error"))
	}

	writer.Write(responseBytes)
}

func getGroupInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	groupName := request.PathValue("group")

	resp, _ := controller.GetGroup(ctx, nil, &water_tank.GetGroupTanks{
		Group: groupName,
	})

	writer.WriteHeader(mapError(resp.Code))

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte("Internal server error"))
	}

	writer.Write(responseBytes)
}
