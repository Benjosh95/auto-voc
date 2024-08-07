package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Benjosh95/auto-voc/internal/models"
	sq "github.com/Masterminds/squirrel"
)

type VocService struct {
	db *sql.DB
}

func NewVocService(db *sql.DB) *VocService {
	return &VocService{db: db}
}

// GetVocs retrieves vocabulary entries based on the provided filter.
func (s *VocService) GetVocs(filter models.VocFilter) ([]models.Voc, error) {
	query := sq.Select("*").From("vocabularies")

	if filter.NextReviewDate != "" {
		query = query.Where("DATE(nextReviewDate) = ?", filter.NextReviewDate)
	}
	if filter.Status != 0 {
		query = query.Where(sq.Eq{"status": filter.Status})
	}
	if filter.ReviewCount != 0 {
		query = query.Where(sq.Eq{"reviewCount": filter.ReviewCount})
	}

	// Convert the Squirrel query to SQL with PostgreSQL placeholders
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	// Debugging: Print the SQL query and arguments
	fmt.Printf("Generated SQL: %s\n", sql)
	fmt.Printf("Arguments: %v\n", args)

	rows, err := s.db.QueryContext(context.TODO(), sql, args...)
	if err != nil {
		fmt.Printf("Failed to get vocabularies: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var vocs []models.Voc
	for rows.Next() {
		var voc models.Voc
		if err := rows.Scan(&voc.ID, &voc.English, &voc.German, &voc.Status, &voc.ReviewCount, &voc.NextReviewDate); err != nil {
			fmt.Printf("Failed to scan row: %v\n", err)
			return nil, err
		}
		vocs = append(vocs, voc)
	}

	return vocs, nil
}

// CreateVoc creates a new vocabulary entry.
func (s *VocService) CreateVoc(voc models.CreateVocRequest) (string, error) {
	query := sq.Insert("vocabularies").
		Columns("english", "german").
		Values(voc.English, voc.German).
		Suffix("RETURNING id")

	sql, args, err := query.ToSql()
	if err != nil {
		return "", err
	}

	var id string
	err = s.db.QueryRowContext(context.TODO(), sql, args...).Scan(&id)
	if err != nil {
		log.Printf("Failed to create vocabulary: %v", err)
		return "", err
	}

	return id, nil
}

// UpdateVoc updates an existing vocabulary entry.
func (s *VocService) UpdateVoc(id string, voc models.UpdateVocRequest) (*models.Voc, error) {
	query := sq.Update("vocabularies").
		Set("english", voc.English).
		Set("german", voc.German).
		Set("status", voc.Status).
		Set("reviewCount", voc.ReviewCount).
		Set("nextReviewDate", voc.NextReviewDate).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, english, german, status, reviewCount, nextReviewDate")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var updatedVoc models.Voc
	err = s.db.QueryRowContext(context.TODO(), sql, args...).Scan(
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

// DeleteVoc deletes a vocabulary entry by ID.
func (s *VocService) DeleteVoc(id string) error {
	query := sq.Delete("vocabularies").Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(context.TODO(), sql, args...)
	if err != nil {
		log.Printf("Failed to delete vocabulary with id: %v", id)
		return err
	}

	return nil
}
