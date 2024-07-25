package web

import "net/http"

func writeInternalServerError(writer http.ResponseWriter) {
	writer.Write([]byte("Internal server error"))
	writer.WriteHeader(http.StatusInternalServerError)
}
