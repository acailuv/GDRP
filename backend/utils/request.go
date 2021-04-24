package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadBody(body io.ReadCloser, model interface{}) error {
	reqBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	json.Unmarshal(reqBody, &model)
	return nil
}

func SetupCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS")
	(*w).Header().Add("Access-Control-Allow-Headers", "*")
}
