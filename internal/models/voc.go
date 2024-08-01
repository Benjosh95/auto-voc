package models

import "time"

type Voc struct {
	ID             string    `json:"id"`
	English        string    `json:"english"`
	German         string    `json:"german"`
	Status         int       `json:"status"`         // Leitner box number (1 to 5)
	ReviewCount    int       `json:"reviewCount"`    // Number of reviews
	NextReviewDate time.Time `json:"nextReviewDate"` // Next review date
}

type CreateVocRequest struct {
	English string `json:"english"`
	German  string `json:"german"`
}

type UpdateVocRequest struct {
	English        string    `json:"english"`
	German         string    `json:"german"`
	Status         int       `json:"status"`
	ReviewCount    int       `json:"reviewCount"`
	NextReviewDate time.Time `json:"nextReviewDate"`
}
