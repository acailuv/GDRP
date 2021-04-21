package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadBody(body io.ReadCloser, model interface{}) error {
	reqBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	json.Unmarshal(reqBody, &model)
	return nil
}
