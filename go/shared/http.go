package shared

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID *int
	Username string
	FullName *string
	Initials *string
	Email string
	Bio *string
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func SuccessResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
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

func DeleteCookie(w http.ResponseWriter, key string) {
	cookie := http.Cookie{
		Name:   key,
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)
}