package model

type Link struct {
	ID       int    `json:"id" db:"id"`
	baseUrl  string `json:"base_url" db:"base_url"`
	shortUrl string `json:"short_url" db:"short_url"`
}
