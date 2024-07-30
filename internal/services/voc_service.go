package services

import (
	"fmt"
)

type VocService struct {
}

func NewVocService() *VocService {
	return &VocService{}
}

func (s *VocService) GetVocs() {
	fmt.Printf("Get all Vocabularies...")
	// Return Vocs with Type and err
	return
}

func (s *VocService) CreateVoc() {
	fmt.Printf("Get all Vocabularies...")
	// Return Vocs with Type and err
	return
}
