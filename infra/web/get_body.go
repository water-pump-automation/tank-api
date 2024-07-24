package web

import (
	"encoding/json"
	"net/http"
	"water-tank-api/app/controllers"
)

func getBody(writer http.ResponseWriter, request *http.Request, body any) error {
	var bodyBytes []byte

	bodyReader, err := request.GetBody()
	if err != nil {
		internalServerError(writer)
		return err
	}
	_, err = bodyReader.Read(bodyBytes)
	if err != nil {
		internalServerError(writer)
		return err
	}
	defer bodyReader.Close()

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		responseError := controllers.NewControllerError(controllers.WaterTankBadRequest, "Bad request. Wrong type parameter")
		response, err := json.Marshal(responseError)
		if err != nil {
			internalServerError(writer)
			return err
		}
		writer.Write(response)
		writer.WriteHeader(http.StatusBadRequest)
		return err
	}
	return nil
}
