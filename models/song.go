package models

type Song struct {
	ID     int    `json:"id"`
	Judul  string `json:"judul"`
	Artis  string `json:"artis"`
	Durasi int    `json:"durasi"`
}
