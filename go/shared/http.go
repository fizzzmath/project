package shared

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}