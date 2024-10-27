package helpers

import (
	"encoding/json"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, statusCode int, status interface{}) {
	respb, _ := json.Marshal(status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respb)
}

func ReturnErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	http.Error(w, err.Error(), statusCode)
}
