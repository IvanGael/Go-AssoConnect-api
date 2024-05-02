//authUtils.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func generateToken(user User) (string, error) {
	// Charge les variables d'environnement à partir du fichier .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		// Ajoute d'autres revendications selon les besoins
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func extractTokenFromHeader(r *http.Request) string {
	// Récupère le token à partir de l'en-tête Authorization
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}

	// Le token est généralement envoyé dans le format "Bearer <token>"
	// Nous devons extraire le token lui-même
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}

// Middleware pour vérifier le token JWT
func authenticate(next http.Handler) http.Handler {
	// Charge les variables d'environnement à partir du fichier .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Le token est valide, passe à la prochaine étape
		next.ServeHTTP(w, r)
	})
}

func getUserIDFromToken(r *http.Request) (uint, error) {
	// Charge les variables d'environnement à partir du fichier .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}

	tokenString := extractTokenFromHeader(r)
	if tokenString == "" {
		return 0, errors.New("Token not provided")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, errors.New("Invalid token")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("Invalid user ID in token")
	}

	return uint(userID), nil
}

func HandleSecureResource(w http.ResponseWriter, r *http.Request) {
	// Obtient l'ID de l'utilisateur à partir du token
	userID, err := getUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'ID de l'utilisateur", http.StatusUnauthorized)
		return
	}

	// Utilise l'ID de l'utilisateur pour effectuer des opérations sécurisées
	// ...

	// Répond avec des données sécurisées
	response := map[string]string{"message": "Ressource sécurisée", "userID": fmt.Sprint(userID)}
	json.NewEncoder(w).Encode(response)
}
