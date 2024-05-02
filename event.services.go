// event.services.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetEvent gère la requête GET pour récupérer un event par ID
func HandleGetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["id"]

	var event Event
	if err := db.First(&event, eventID).Error; err != nil {
		http.Error(w, "Event non trouvé", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation de l'event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleGetEvents gère la requête GET pour récupérer la liste des events
func HandleGetEvents(w http.ResponseWriter, r *http.Request) {
	var events []Event
	if err := db.Find(&events).Error; err != nil {
		http.Error(w, "Erreur lors de la récupération des events", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(events)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation des events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleCreateEvent gère la requête POST pour créer un event
func HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de l'event", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newEvent).Error; err != nil {
		log.Println("Erreur lors de la création de l'event:", err)
		http.Error(w, "Erreur lors de la création de l'event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Event créé avec succès"))
}

// handleUpdateEvent gère la requête PUT pour mettre à jour un event par ID
func HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["id"]

	var existingEvent Event

	// Vérifie si l'event avec l'ID spécifié existe
	if err := db.First(&existingEvent, eventID).Error; err != nil {
		http.Error(w, "Event non trouvé", http.StatusNotFound)
		return
	}

	var updatedEvent Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de mise à jour de l'event", http.StatusBadRequest)
		return
	}

	// Mise à jour de l'event dans la base de données
	if err := db.Model(&Event{}).Where("id = ?", eventID).Updates(updatedEvent).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour de l'event", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Event mis à jour avec succès. ID : " + eventID))
}

// handleDeleteEvent gère la requête DELETE pour supprimer un event par ID
func HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["id"]

	var existingEvent Event

	// Vérifie si l'event avec l'ID spécifié existe
	if err := db.First(&existingEvent, eventID).Error; err != nil {
		http.Error(w, "Event non trouvé", http.StatusNotFound)
		return
	}

	// Suppression de l'event dans la base de données en utilisant le modèle existant
	if err := db.Delete(&existingEvent).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression de l'event", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Event supprimé avec succès. ID : " + eventID))
}
