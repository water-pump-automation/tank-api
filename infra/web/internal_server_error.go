package web

import "net/http"

func internalServerError(writer http.ResponseWriter) {
	writer.Write([]byte("Internal server error"))
	writer.WriteHeader(http.StatusInternalServerError)
}
