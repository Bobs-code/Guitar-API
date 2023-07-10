package models

type Guitar struct {
	Id          int    `json:"id"`
	Brand_id    int    `json:"brand_id"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
