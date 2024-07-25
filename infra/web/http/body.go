package web

import (
	"encoding/json"
	"net/http"
)

func getResponseBody(writer http.ResponseWriter, response any) (responseBytes []byte, err error) {
	responseBytes, err = json.Marshal(response)
	if err != nil {
		writeInternalServerError(writer)
		return
	}
	return
}

func writeBadRequestError(writer http.ResponseWriter, err error) {
	errorBody, err := getResponseBody(writer, NewAPIError(WaterTankBadRequest, err.Error()))
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write(errorBody)
}

func getBody(writer http.ResponseWriter, request *http.Request, body any) error {
	var bodyBytes []byte

	bodyReader, err := request.GetBody()
	if err != nil {
		writeInternalServerError(writer)
		return err
	}
	_, err = bodyReader.Read(bodyBytes)
	if err != nil {
		writeInternalServerError(writer)
		return err
	}
	defer bodyReader.Close()

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		responseError := NewAPIError(WaterTankBadRequest, "Bad request. Wrong type parameter")
		response, err := json.Marshal(responseError)
		if err != nil {
			writeInternalServerError(writer)
			return err
		}
		writer.Write(response)
		writer.WriteHeader(http.StatusBadRequest)
		return err
	}
	return nil
}
