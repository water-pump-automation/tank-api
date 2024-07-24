package web

import (
	"encoding/json"
	"net/http"
	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/access"
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
	type PostBody struct {
		Name            string              `json:"name"`
		Group           string              `json:"group"`
		MaximumCapacity water_tank.Capacity `json:"maximum_capacity"`
	}

	var body PostBody

	if err := getBody(writer, request, &body); err != nil {
		return
	}

	response, _ := controller.Create(body.Name, body.Group, body.MaximumCapacity)

	switch response.Code {
	case controllers.WaterTankBadRequest:
		writer.WriteHeader(http.StatusBadRequest)
	case controllers.WaterTankInvalidRequest:
		writer.WriteHeader(http.StatusUnprocessableEntity)
	case controllers.WaterTankInternalServerError:
		writer.WriteHeader(http.StatusInternalServerError)
	case controllers.WaterTankOK:
		writer.WriteHeader(http.StatusOK)
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		internalServerError(writer)
		return
	}
	writer.Write(responseBytes)
}

func updateTankInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	type PatchBody struct {
		CurrentWaterLevel water_tank.Capacity `json:"water_level"`
	}
	var body PatchBody

	tankName := request.PathValue("tank")
	accessToken := request.Header.Get("access_token")
	group := request.Header.Get("group")

	if err := getBody(writer, request, &body); err != nil {
		return
	}

	response, _ := controller.Update(tankName, group, access.AccessToken(accessToken), body.CurrentWaterLevel)

	switch response.Code {
	case controllers.WaterTankBadRequest:
		writer.WriteHeader(http.StatusBadRequest)
	case controllers.WaterTankInvalidRequest:
		writer.WriteHeader(http.StatusUnprocessableEntity)
	case controllers.WaterTankInternalServerError:
		writer.WriteHeader(http.StatusInternalServerError)
	case controllers.WaterTankOK:
		writer.WriteHeader(http.StatusNoContent)
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		internalServerError(writer)
		return
	}
	writer.Write(responseBytes)
}

func getTankInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	tankName := request.PathValue("tank")
	groupName := request.Header.Get("group")

	response, _ := controller.Get(tankName, groupName)

	switch response.Code {
	case controllers.WaterTankNotFound:
		writer.WriteHeader(http.StatusNotFound)
	case controllers.WaterTankInternalServerError:
		writer.WriteHeader(http.StatusInternalServerError)
	case controllers.WaterTankOK:
		writer.WriteHeader(http.StatusOK)
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte("Internal server error"))
	}

	writer.Write(responseBytes)
}

func getGroupInternal(controller *controllers.InternalController, writer http.ResponseWriter, request *http.Request) {
	groupName := request.PathValue("group")

	response, _ := controller.GetGroup(groupName)

	switch response.Code {
	case controllers.WaterTankNotFound:
		writer.WriteHeader(http.StatusNotFound)
	case controllers.WaterTankInternalServerError:
		writer.WriteHeader(http.StatusInternalServerError)
	case controllers.WaterTankOK:
		writer.WriteHeader(http.StatusOK)
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte("Internal server error"))
	}

	writer.Write(responseBytes)
}
