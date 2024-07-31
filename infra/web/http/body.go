package web

import (
	"encoding/json"
	"io"
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
	errorBody, err := getResponseBody(writer, NewAPIValidationError(TankBadRequest, err.Error()))
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write(errorBody)
}

func getBody(writer http.ResponseWriter, request *http.Request) (map[string]interface{}, error) {
	inputMap := make(map[string]interface{})

	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		writeInternalServerError(writer)
		return nil, err
	}
	defer request.Body.Close()

	err = json.Unmarshal(bodyBytes, &inputMap)
	if err != nil {
		writeInternalServerError(writer)
		return nil, err
	}
	return inputMap, nil
}
