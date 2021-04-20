package clock

import (
	"fmt"
	"net/http"
	"time"
)

func (c *clock) CurrentTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", time.Now())
}
