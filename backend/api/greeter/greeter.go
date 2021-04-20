package greeter

import (
	"net/http"
)

type Greeter interface {
	Greet(w http.ResponseWriter, r *http.Request)
}

type greeter struct {
	data GreeterData
}

type GreeterData struct {
	Name   string
	Gender string
}

func NewGreeter(data GreeterData) Greeter {
	return &greeter{
		data: data,
	}
}
