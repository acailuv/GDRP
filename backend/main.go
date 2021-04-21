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
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/users", clients.user.SelectAll).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/user/{id}", clients.user.Select).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/upsert", clients.user.Upsert).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/user/{id}", clients.user.Delete).Methods(http.MethodDelete, http.MethodOptions)

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
