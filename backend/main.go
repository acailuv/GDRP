package main

import (
	"log"
	"main/api/clock"
	"main/api/greeter"
	"net/http"

	"github.com/gorilla/mux"
)

type handlerClients struct {
	greeter greeter.Greeter
	clock   clock.Clock
}

func handleRequests(clients handlerClients) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/greet", clients.greeter.Greet)
	myRouter.HandleFunc("/clock", clients.clock.CurrentTime)

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {

	greeterData := greeter.GreeterData{
		Name:   "John Smith",
		Gender: "Male",
	}
	greeterClient := greeter.NewGreeter(greeterData)

	clock := clock.NewClock()

	clients := handlerClients{
		greeter: greeterClient,
		clock:   clock,
	}

	handleRequests(clients)

}
