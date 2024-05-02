// membership.services.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetMembership gère la requête GET pour récupérer un membership par ID
func HandleGetMembership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	membershipID := vars["id"]

	var membership Membership
	if err := db.First(&membership, membershipID).Error; err != nil {
		http.Error(w, "Membership non trouvé", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(membership)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation du membership", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleGetMemberships gère la requête GET pour récupérer la liste des memberships
func HandleGetMemberships(w http.ResponseWriter, r *http.Request) {
	var memberships []Membership
	if err := db.Find(&memberships).Error; err != nil {
		http.Error(w, "Erreur lors de la récupération des memberships", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(memberships)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation des memberships", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleCreateMembership gère la requête POST pour créer un membership
func HandleCreateMembership(w http.ResponseWriter, r *http.Request) {
	var newMembership Membership
	err := json.NewDecoder(r.Body).Decode(&newMembership)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données du membership", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newMembership).Error; err != nil {
		log.Println("Erreur lors de la création du membership:", err)
		http.Error(w, "Erreur lors de la création du membership", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Membership créé avec succès"))
}

// handleUpdateMembership gère la requête PUT pour mettre à jour un membership par ID
func HandleUpdateMembership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	membershipID := vars["id"]

	var existingMembership Membership

	// Vérifie si le membership avec l'ID spécifié existe
	if err := db.First(&existingMembership, membershipID).Error; err != nil {
		http.Error(w, "Membership non trouvé", http.StatusNotFound)
		return
	}

	var updatedMembership Membership
	err := json.NewDecoder(r.Body).Decode(&updatedMembership)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de mise à jour du membership", http.StatusBadRequest)
		return
	}

	// Mise à jour du membership dans la base de données
	if err := db.Model(&Membership{}).Where("id = ?", membershipID).Updates(updatedMembership).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour du membership", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Membership mis à jour avec succès. ID : " + membershipID))
}

// handleDeleteMembership gère la requête DELETE pour supprimer un membership par ID
func HandleDeleteMembership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	membershipID := vars["id"]

	var existingMembership Membership

	// Vérifie si le membership avec l'ID spécifié existe
	if err := db.First(&existingMembership, membershipID).Error; err != nil {
		http.Error(w, "Membership non trouvé", http.StatusNotFound)
		return
	}

	// Suppression du membership dans la base de données en utilisant le modèle existant
	if err := db.Delete(&existingMembership).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression du membership", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Membership supprimé avec succès. ID : " + membershipID))
}
