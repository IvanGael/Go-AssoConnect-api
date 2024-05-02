// ticket.services.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetTicket gère la requête GET pour récupérer un ticket par ID
func HandleGetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	var ticket Ticket
	if err := db.First(&ticket, ticketID).Error; err != nil {
		http.Error(w, "Ticket non trouvé", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(ticket)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation du ticket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleGetTickets gère la requête GET pour récupérer la liste des tickets
func HandleGetTickets(w http.ResponseWriter, r *http.Request) {
	var tickets []Ticket
	if err := db.Find(&tickets).Error; err != nil {
		http.Error(w, "Erreur lors de la récupération des tickets", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tickets)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation des tickets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleCreateTicket gère la requête POST pour créer un ticket
func HandleCreateTicket(w http.ResponseWriter, r *http.Request) {
	var newTicket Ticket
	err := json.NewDecoder(r.Body).Decode(&newTicket)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données du ticket", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newTicket).Error; err != nil {
		log.Println("Erreur lors de la création du ticket:", err)
		http.Error(w, "Erreur lors de la création du ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ticket créé avec succès"))
}

// handleUpdateTicket gère la requête PUT pour mettre à jour un ticket par ID
func HandleUpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	var existingTicket Ticket

	// Vérifie si le ticket avec l'ID spécifié existe
	if err := db.First(&existingTicket, ticketID).Error; err != nil {
		http.Error(w, "Ticket non trouvé", http.StatusNotFound)
		return
	}

	var updatedTicket Ticket
	err := json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de mise à jour du ticket", http.StatusBadRequest)
		return
	}

	// Mise à jour du ticket dans la base de données
	if err := db.Model(&Ticket{}).Where("id = ?", ticketID).Updates(updatedTicket).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour du ticket", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ticket mis à jour avec succès. ID : " + ticketID))
}

// handleDeleteTicket gère la requête DELETE pour supprimer un ticket par ID
func HandleDeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	var existingTicket Ticket

	// Vérifie si le ticket avec l'ID spécifié existe
	if err := db.First(&existingTicket, ticketID).Error; err != nil {
		http.Error(w, "Ticket non trouvé", http.StatusNotFound)
		return
	}

	// Suppression du ticket dans la base de données en utilisant le modèle existant
	if err := db.Delete(&existingTicket).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression du ticket", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ticket supprimé avec succès. ID : " + ticketID))
}
