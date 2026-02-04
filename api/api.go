package api

import (
	"net/http"
	"encoding/json"
	"go-learn/handlers"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance float64 
}

type ErrorResponse struct {
	Code int
	Message string
}

func writeError (w http.ResponseWriter, code int, message string) {
	response := ErrorResponse{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}

	InternalServerErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
)