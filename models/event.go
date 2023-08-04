package models

import "time"

type Event struct {
	EventID     int        `gorm:"primary_key;auto_increment" json:"id"`
	Title       string     `gorm:"size:255" json:"title"`
	Description string     `gorm:"size:255" json:"description"`
	Location    string     `gorm:"size:255" json:"location"`
	CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
