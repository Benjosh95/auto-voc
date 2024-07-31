package api

import (
	"fmt"
	"net/http"

	"github.com/Benjosh95/auto-voc/internal/models"
	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/gin-gonic/gin"
)

type VocHandler struct {
	vocService *services.VocService
}

func NewVocHandler(vocService *services.VocService) *VocHandler {
	return &VocHandler{
		vocService: vocService,
	}
}

func (h *VocHandler) GetVocs(c *gin.Context) {
	// TODO: request processing?

	vocs, err := h.vocService.GetVocs()
	if err != nil {
		fmt.Printf("Failed to get vocs: %v, err: %v\n", vocs, err)
		c.JSON(500, gin.H{
			"Message": "Error getting vocs",
		})
		return
	}

	fmt.Printf("Succesfully recieved vocs: %v", vocs)
	c.JSON(200, gin.H{
		"vocs": vocs,
	})
}

func (h *VocHandler) CreateVoc(c *gin.Context) {
	// TODO request processing
	var req models.CreateVocRequest

	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Failed to parse body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	id, err := h.vocService.CreateVoc(req)
	if err != nil {
		fmt.Printf("Failed to create voc: %v\n", err)
		c.JSON(500, gin.H{
			"Message": "Error creating a voc",
		})
		return
	}

	fmt.Printf("Succesfully created voc with id: %v and content: %v = %v", id, req.English, req.German)
	c.JSON(200, gin.H{
		"Message": fmt.Sprintf("successfully created voc with id: %v", id),
	})
}
