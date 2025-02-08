package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	code    int
	message any
}

type Map map[string]interface{}

func (e APIError) Error() string {
	return fmt.Sprintf("api error %d", e.code)
}

func NewAPIError(code int, message any) APIError {
	return APIError{code: code, message: message}
}

func InvalidJSON(errors map[string]interface{}) APIError {
	return APIError{code: http.StatusBadRequest, message: errors}
}

func InternalServerError(err error) APIError {
	return APIError{code: http.StatusInternalServerError, message: err}
}

func Unauthorized() APIError {
	return APIError{code: http.StatusUnauthorized, message: "User Unauthorized"}
}

func WriteJSON(w http.ResponseWriter, status int, message any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func DecodeJSON(body io.ReadCloser, payload any) error {

	return nil
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			if apiErr, ok := err.(APIError); ok {
				WriteJSON(w, apiErr.code, apiErr.message)
			}
			WriteJSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}
