// user.services.go

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetUser gère la requête GET pour récupérer un utilisateur par ID
func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation de l'utilisateur", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleGetUsers gère la requête GET pour récupérer la liste des utilisateurs
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, "Erreur lors de la récupération des utilisateurs", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Erreur lors de la sérialisation des utilisateurs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleCreateUser gère la requête POST pour créer un utilisateur
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de l'utilisateur", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newUser).Error; err != nil {
		log.Println("Erreur lors de la création de l'utilisateur:", err)
		http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Utilisateur créé avec succès"))
}

// handleUpdateUser gère la requête PUT pour mettre à jour un utilisateur par ID
func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var existingUser User

	// Vérifie si l'utilisateur avec l'ID spécifié existe
	if err := db.First(&existingUser, userID).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données de mise à jour de l'utilisateur", http.StatusBadRequest)
		return
	}

	// Mise à jour de l'utilisateur dans la base de données
	if err := db.Model(&User{}).Where("id = ?", userID).Updates(updatedUser).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour de l'utilisateur", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Utilisateur mis à jour avec succès. ID : " + userID))
}

// handleDeleteUser gère la requête DELETE pour supprimer un utilisateur par ID
func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var existingUser User

	// Vérifie si l'utilisateur avec l'ID spécifié existe
	if err := db.First(&existingUser, userID).Error; err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	// Suppression de l'utilisateur dans la base de données en utilisant le modèle existant
	if err := db.Delete(&existingUser).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression de l'utilisateur", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Utilisateur supprimé avec succès. ID : " + userID))
}
