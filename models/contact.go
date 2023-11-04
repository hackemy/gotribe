package models

type Contact struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
}
