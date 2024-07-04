package models

type Event struct {
	EventID     int       `json:"eventid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    *Location `json:"location"`
	Amount      string    `json:"amount"`
	Date        string    `json:"date"`
	Creator     string    `json:"creator"`
	Liked       bool      `json:"liked"`
	ImageUrl    string    `json:"imageUrl"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
