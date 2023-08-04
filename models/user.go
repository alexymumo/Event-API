package models

import "time"

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	Email     string    `gorm:"size:255" json:"email"`
	Password  string    `gorm:"size:255" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
