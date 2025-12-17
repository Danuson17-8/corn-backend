package models

type Promotion struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	IsActive    bool   `json:"isActive"`
	Link        string `json:"link"`
}
