// association.services.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetAssociation gère la requête GET pour récupérer une association par ID
func HandleGetAssociation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	associationID := vars["id"]

	var association Association
	if err := db.First(&association, associationID).Error; err != nil {
		http.Error(w, "Association non trouvé", http.StatusNotFound)
		return
	}

	var user User
	if err := db.First(&user, association.CreatedBy).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	association.TheUser = user

	response, err := json.Marshal(association)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation de l'association", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleGetAssociations gère la requête GET pour récupérer la liste des associations
func HandleGetAssociations(w http.ResponseWriter, r *http.Request) {
	var associations []Association
	if err := db.Find(&associations).Error; err != nil {
		http.Error(w, "Erreur lors de la récupération des associations", http.StatusInternalServerError)
		return
	}

	for i := 0; i <= len(associations)-1; i++ {
		var user User
		if err := db.First(&user, associations[i].CreatedBy).Error; err != nil {
			http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
			return
		}

		associations[i].TheUser = user
	}

	response, err := json.Marshal(associations)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation des associations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleCreateAssociation gère la requête POST pour créer une association
func HandleCreateAssociation(w http.ResponseWriter, r *http.Request) {
	var newAssociation Association
	err := json.NewDecoder(r.Body).Decode(&newAssociation)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de l'association", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newAssociation).Error; err != nil {
		log.Println("Erreur lors de la création de l'association:", err)
		http.Error(w, "Erreur lors de la création de l'association", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Association créé avec succès"))
}

// handleUpdateAssociation gère la requête PUT pour mettre à jour une association par ID
func HandleUpdateAssociation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	associationID := vars["id"]

	var existingAssociation Association

	// Vérifie si l'association avec l'ID spécifié existe
	if err := db.First(&existingAssociation, associationID).Error; err != nil {
		http.Error(w, "Association non trouvé", http.StatusNotFound)
		return
	}

	var updatedAssociation Association
	err := json.NewDecoder(r.Body).Decode(&updatedAssociation)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de mise à jour de l'association", http.StatusBadRequest)
		return
	}

	// Mise à jour de l'association dans la base de données
	if err := db.Model(&Association{}).Where("id = ?", associationID).Updates(updatedAssociation).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour de l'association", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Association mis à jour avec succès. ID : " + associationID))
}

// handleDeleteAssociation gère la requête DELETE pour supprimer une association par ID
func HandleDeleteAssociation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	associationID := vars["id"]

	var existingAssociation Association

	// Vérifie si l'association avec l'ID spécifié existe
	if err := db.First(&existingAssociation, associationID).Error; err != nil {
		http.Error(w, "Association non trouvé", http.StatusNotFound)
		return
	}

	// Suppression de l'association dans la base de données en utilisant le modèle existant
	if err := db.Delete(&existingAssociation).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression de l'association", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Association supprimé avec succès. ID : " + associationID))
}
