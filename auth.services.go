// auth.services.go

package main

import (
	"encoding/json"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Implémente la logique de login et génère le token JWT en cas de succès
	// ...

	var user User

	// Exemple de génération du token et envoi de la réponse
	token, err := generateToken(user)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Implémente la logique d'inscription et génère le token JWT en cas de succès
	// ...

	var user User

	// Exemple de génération du token et envoi de la réponse
	token, err := generateToken(user)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}
