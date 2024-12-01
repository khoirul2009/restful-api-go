// internal/models/user.go
package models

import "time"

// User represents the structure of a user in the system
type User struct {
	ID              string    `json:"id" gorm:"primaryKey"`         // User ID, the primary key in the database
	Name            string    `json:"name" gorm:"size:100"`         // User's name, limited to 100 characters
	Email           string    `json:"email" gorm:"unique;not null"` // Email address, unique and not nullable
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"-" gorm:"not null"` // Password, not exposed in the JSON response
	CreatedAt       time.Time `json:"created_at"`        // Timestamp for when the user was created
	UpdatedAt       time.Time `json:"updated_at"`        // Timestamp for when the user was last updated

	Tokens []SessionToken `json:"tokens" gorm:"foreignKey:UserID;references:ID"`
}
