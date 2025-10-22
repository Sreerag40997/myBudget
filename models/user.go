package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"not null" json:"-"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone,omitempty"`
	ProfileImage string    `json:"profile_image,omitempty"`
	Role         int       `json:"role" gorm:"default:0"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
