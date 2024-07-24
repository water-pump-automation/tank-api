package web

import (
	"encoding/json"
	"net/http"
	"water-tank-api/app/controllers"
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
		internalServerError(writer)
		return
	}

	writer.Write(responseBytes)
}

func getGroupExternal(controller *controllers.ExternalController, writer http.ResponseWriter, request *http.Request) {
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
		internalServerError(writer)
		return
	}

	writer.Write(responseBytes)
}
