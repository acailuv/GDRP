package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, process string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Internal Server Error: %v: %v", process, err)
}

func OK(w http.ResponseWriter, operation string) {
	fmt.Fprintf(w, "200 OK: %v Successful!", operation)
}

func OKWithData(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
