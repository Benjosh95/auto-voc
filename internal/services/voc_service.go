package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type VocService struct {
	db *sql.DB
}

func NewVocService(db *sql.DB) *VocService {
	return &VocService{db: db}
}

// TODO: Add Ctx
func (s *VocService) GetVocs() {
	fmt.Printf("Get all Vocabularies...")
	// Return Vocs with Type and err
	return
}

// TODO: Add Ctx
func (s *VocService) CreateVoc() {
	fmt.Printf("Create a Vocabulary")
	query := `INSERT INTO vocabularies (english, german) VALUES ($1, $2)`
	// TODO: Add Ctx
	_, err := s.db.ExecContext(context.TODO(), query, "success", "Erfolg")
	if err != nil {
		log.Printf("Failed to create vocabulary: %v", err)
	}
	// Return Vocs with Type and err
	return
}
