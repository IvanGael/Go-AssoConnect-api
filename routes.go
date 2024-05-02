// routes.go

package main

import (
	"github.com/gorilla/mux"
)

func setupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Routes de l'API

	// Routes publiques (accessible sans authentification)

	publicRouter := r.PathPrefix("/api").Subrouter()
	publicRouter.HandleFunc("/login", HandleLogin).Methods("POST")
	publicRouter.HandleFunc("/register", HandleRegister).Methods("POST")

	// Routes protégées (accessible uniquement avec un token JWT valide)
	protectedRouter := r.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(authenticate)
	//USERS
	protectedRouter.HandleFunc("/users", HandleGetUsers).Methods("GET")
	protectedRouter.HandleFunc("/users/{id}", HandleGetUser).Methods("GET")
	protectedRouter.HandleFunc("/users", HandleCreateUser).Methods("POST")
	protectedRouter.HandleFunc("/users/{id}", HandleUpdateUser).Methods("PUT")
	protectedRouter.HandleFunc("/users/{id}", HandleDeleteUser).Methods("DELETE")
	//ASSOCIATIONS
	protectedRouter.HandleFunc("/associations", HandleGetAssociations).Methods("GET")
	protectedRouter.HandleFunc("/associations/{id}", HandleGetAssociation).Methods("GET")
	protectedRouter.HandleFunc("/associations", HandleCreateAssociation).Methods("POST")
	protectedRouter.HandleFunc("/associations/{id}", HandleUpdateAssociation).Methods("PUT")
	protectedRouter.HandleFunc("/associations/{id}", HandleDeleteAssociation).Methods("DELETE")
	//EVENTS
	protectedRouter.HandleFunc("/events", HandleGetEvents).Methods("GET")
	protectedRouter.HandleFunc("/events/{id}", HandleGetEvent).Methods("GET")
	protectedRouter.HandleFunc("/events", HandleCreateEvent).Methods("POST")
	protectedRouter.HandleFunc("/events/{id}", HandleUpdateEvent).Methods("PUT")
	protectedRouter.HandleFunc("/events/{id}", HandleDeleteEvent).Methods("DELETE")
	//TICKETS
	protectedRouter.HandleFunc("/tickets", HandleGetTickets).Methods("GET")
	protectedRouter.HandleFunc("/tickets/{id}", HandleGetTicket).Methods("GET")
	protectedRouter.HandleFunc("/tickets", HandleCreateTicket).Methods("POST")
	protectedRouter.HandleFunc("/tickets/{id}", HandleUpdateTicket).Methods("PUT")
	protectedRouter.HandleFunc("/tickets/{id}", HandleDeleteTicket).Methods("DELETE")

	return r
}
