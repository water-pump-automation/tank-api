package web

import "net/http"

func writeInternalServerError(writer http.ResponseWriter) {
	http.Error(writer, "Internal server error", http.StatusInternalServerError)
}
