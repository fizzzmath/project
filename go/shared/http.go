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

func SetCookie(w http.ResponseWriter, key string, value string) {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}