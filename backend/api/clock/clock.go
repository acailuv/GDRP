package clock

import (
	"net/http"
)

type Clock interface {
	CurrentTime(w http.ResponseWriter, r *http.Request)
}

type clock struct {
}

func NewClock() Clock {
	return &clock{}
}
