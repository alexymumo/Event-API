package models

type Event struct {
	EventID     int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	//CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	//UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
