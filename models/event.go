package models

import "time"

type Event struct {
	EventID     int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    *Location `json:"location"`
	Amount      string    `json:"amount"`
	Date        string    `json:"date"`
	Creator     string    `json:"creator"`
	Liked       bool      `json:"liked"`
	ImageUrl    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Location struct {
	Longitude int64 `json:"longitude"`
	Latitude  int64 `json:"latitude"`
}
