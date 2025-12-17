package models

type CornMenu struct {
	ID          string    `json:"id"`
	Name        Translate `json:"name"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
}
