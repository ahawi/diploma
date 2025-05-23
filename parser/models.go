package models

type Dog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Age         int    `json:"age"`
	Size        string `json:"size"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}
