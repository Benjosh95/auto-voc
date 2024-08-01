package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/Benjosh95/auto-voc/internal/models"
)

type VocService struct {
	db *sql.DB
}

func NewVocService(db *sql.DB) *VocService {
	return &VocService{db: db}
}

// TODO: Add Ctx?
func (s *VocService) GetVocs() ([]models.Voc, error) {
	query := `SELECT * FROM vocabularies`

	rows, err := s.db.QueryContext(context.TODO(), query)
	if err != nil {
		log.Printf("Failed to get Vocabularies: %v", err)
		return nil, err
	}
	defer rows.Close()

	var vocs []models.Voc
	for rows.Next() {
		var voc models.Voc
		if err := rows.Scan(&voc.ID, &voc.English, &voc.German, &voc.Status, &voc.ReviewCount, &voc.NextReviewDate); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		vocs = append(vocs, voc)
	}

	return vocs, nil
}

// TODO: Add Ctx?
func (s *VocService) CreateVoc(voc models.CreateVocRequest) (string, error) {
	query := `INSERT INTO vocabularies (english, german) VALUES ($1, $2) RETURNING id`

	var id string
	// TODO: Add Ctx
	err := s.db.QueryRowContext(context.TODO(), query, voc.English, voc.German).Scan(&id)
	if err != nil {
		log.Printf("Failed to create vocabulary: %v", err)
		return "", err
	}

	return id, nil
}

// TODO: better respond the updated voc
func (s *VocService) UpdateVoc(id string, voc models.UpdateVocRequest) (*models.Voc, error) {
	query := `
		UPDATE vocabularies
		SET english = $1,
		    german = $2,
		    status = $3,
		    reviewCount = $4,
		    nextReviewDate = $5
		WHERE id = $6
		RETURNING id, english, german, status, reviewCount, nextReviewDate;
	`

	var updatedVoc models.Voc
	err := s.db.QueryRowContext(context.TODO(), query, voc.English, voc.German, voc.Status, voc.ReviewCount, voc.NextReviewDate, id).Scan(
		&updatedVoc.ID,
		&updatedVoc.English,
		&updatedVoc.German,
		&updatedVoc.Status,
		&updatedVoc.ReviewCount,
		&updatedVoc.NextReviewDate,
	)
	if err != nil {
		return nil, err
	}

	return &updatedVoc, nil
}

func (s *VocService) DeleteVoc(id string) error {
	query := `DELETE FROM vocabularies WHERE id = $1`

	_, err := s.db.ExecContext(context.TODO(), query, id)
	if err != nil {
		log.Printf("Failed to delete vocabulary with id: %v", id)
		return err
	}

	return nil
}
