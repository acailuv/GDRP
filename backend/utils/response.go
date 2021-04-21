package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func InternalServerError(w http.ResponseWriter, process string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ServerResponse{
		Status: http.StatusInternalServerError,
		Msg:    fmt.Sprintf("Internal Server Error: %v: %v", process, err),
	})
}

func OK(w http.ResponseWriter, operation string) {
	json.NewEncoder(w).Encode(ServerResponse{
		Status: http.StatusOK,
		Msg:    fmt.Sprintf("200 OK: %v Successful!", operation),
	})
}

func OKWithData(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
