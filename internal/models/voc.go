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

type VocFilter struct {
	NextReviewDate string `form:"nextReviewDate" validate:"omitempty,datetime=2006-01-02"`
	Status         int    `form:"status" validate:"omitempty,min=1,max=5"`
	ReviewCount    int    `form:"reviewCount" validate:"omitempty,min=0"`
	// Add more filters as needed
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
