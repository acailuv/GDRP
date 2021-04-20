package greeter

import (
	"fmt"
	"net/http"
)

func (g *greeter) Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v. You are a %v", g.data.Name, g.data.Gender)
}
