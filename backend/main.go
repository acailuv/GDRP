package main

import (
	"fmt"
	"log"
	"net/http"

	api "main/api/user"
	"main/database/connection"

	"github.com/gorilla/mux"
)

type handlerClients struct {
	user api.User
}

func handleRequests(clients handlerClients) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user/{id}", clients.user.Select)
	router.HandleFunc("/users", clients.user.SelectAll)
	router.HandleFunc("/upsert", clients.user.Upsert)
	router.HandleFunc("/delete/{id}", clients.user.Delete)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {

	db := connection.NewConnection()

	userClient := api.NewUserClient(db)

	clients := handlerClients{
		user: userClient,
	}

	fmt.Println("Listening at *:5000")
	handleRequests(clients)

}
