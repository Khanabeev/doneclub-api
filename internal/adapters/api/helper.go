package api

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func GetStatus(r *http.Request) string {
	var status string
	keys, ok := r.URL.Query()["status"]

	if ok {
		status = keys[0]
	}
	return status
}
