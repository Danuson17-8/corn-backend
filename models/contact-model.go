package models

type ContactUser struct {
	Topic   string `json:"topic"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
