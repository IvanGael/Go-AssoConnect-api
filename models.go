// models.go

package main

import (
	"time"
	// "gorm.io/gorm"
)

// // User représente le modèle d'utilisateur
// type User struct {
// 	// gorm.Model
// 	// ID        uint `gorm:"primarykey"`
// 	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	Username     string `gorm:"unique;not null"`
// 	Email        string `gorm:"unique;not null"`
// 	FirstName    string
// 	LastName     string
// 	IsAdmin      bool
// 	Associations []Association `gorm:"foreignKey:CreatedBy"`
// }

// type Association struct {
// 	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// 	Name          string `gorm:"unique;not null"`
// 	description   string `gorm:"unique;not null"`
// 	CreatedBy     string // Clé étrangère qui référence l'ID de l'utilisateur
// 	CreatedByUser User   `gorm:"foreignKey:CreatedBy"`
// }

// type Membership struct {
// 	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// 	UserId string
// 	TheUser User   `gorm:"foreignKey:UserId"`
// 	AssociationId string
// 	TheAsso Association `gorm:"foreignKey:AssociationId"`
// 	Status string
// 	ExpirationDate time.Time
// }

// type Event struct {
// 	ID            string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// 	Title string
// 	Description string
// 	Location string
// 	StartDate time.Time
// 	EndDate time.Time
// 	AssociationId string
// 	TheAsso Association `gorm:"foreignKey:AssociationId"`
// }

// User represents a user of the application.
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	FirstName string
	LastName  string
	IsAdmin   bool
}

// Association represents an association.
type Association struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"not null"`
	CreatedBy   uint      `gorm:"not null"` // Foreign key referencing User.ID
	TheUser     User      `gorm:"foreignKey:CreatedBy;not null"`
	Members     []Membership
	Events      []Event
}

// Membership represents a user's membership to an association.
type Membership struct {
	ID            uint        `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
	UserID        uint        `gorm:"not null"` // Foreign key referencing User.ID
	TheUser       User        `gorm:"foreignKey:UserID;not null"`
	AssociationID uint        `gorm:"not null"` // Foreign key referencing Association.ID
	TheAsso       Association `gorm:"foreignKey:AssociationID;not null"`
	Status        string      `gorm:"not null"` // Status can be "active" or "inactive"
	Expiration    time.Time
}

// Event represents an event organized by an association.
type Event struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	Title         string    `gorm:"not null"`
	Description   string
	Location      string
	StartDate     time.Time
	EndDate       time.Time
	AssociationID uint        `gorm:"not null"` // Foreign key referencing Association.ID
	TheAsso       Association `gorm:"foreignKey:AssociationID;not null"`
	Participants  []User      `gorm:"many2many:event_participants;"`
	Tickets       []Ticket
}

// Ticket represents a ticket for an event.
type Ticket struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	Title         string    `gorm:"not null"`
	Description   string
	Price         float64     `gorm:"not null"`
	AssociationID uint        `gorm:"not null"` // Foreign key referencing Association.ID
	TheAsso       Association `gorm:"foreignKey:AssociationID;not null"`
	EventID       uint        `gorm:"not null"` // Foreign key referencing Event.ID
	TheEvent      Event       `gorm:"foreignKey:EventID;not null"`
	BuyerID       uint        `gorm:"not null"` // Foreign key referencing User.ID
	TheBuyer      User        `gorm:"foreignKey:BuyerID;not null"`
}
