package models

type Voc struct {
	ID      string `json:"id"`
	English string `json:"english"`
	German  string `json:"german"`
}

type CreateVocRequest struct {
	English string `json:"english"`
	German  string `json:"german"`
}
