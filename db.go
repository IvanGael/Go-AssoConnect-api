// db.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Charge les variables d'environnement à partir du fichier .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Impossible de charger le fichier .env")
	}

	// Initialise la connexion à la base de données
	setupDB()
}

func setupDB() {
	var err error

	// Configuration de la connexion à PostgreSQL avec Gorm
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// // AutoMigrate pour créer automatiquement les tables
	// err = db.AutoMigrate(&User{}, &Article{}) // Ajoute tous les modèles ici
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// AutoMigrate pour créer automatiquement les tables
	err = db.AutoMigrate(&User{}, &Association{}, &Membership{}, &Event{}, &Ticket{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connexion à la base de données réussie!")
}

// CloseDB ferme la connexion à la base de données
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	log.Println("Fermeture de la connexion à la base de données.")
}
