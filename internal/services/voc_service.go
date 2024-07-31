package services

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Printf("Trying to get all Vocabularies...")

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
		if err := rows.Scan(&voc.ID, &voc.English, &voc.German); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, err
		}
		vocs = append(vocs, voc)
	}

	return vocs, nil
}

// TODO: Add Ctx?
func (s *VocService) CreateVoc(voc models.CreateVocRequest) (string, error) {
	fmt.Printf("Trying to create a vocabulary...")

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
