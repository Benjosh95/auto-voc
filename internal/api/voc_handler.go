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
	vocs, err := h.vocService.GetVocs()
	if err != nil {
		fmt.Printf("Failed to get vocs: %v, err: %v\n", vocs, err)
		c.JSON(500, gin.H{
			"Message": "Error getting vocs",
		})
		return
	}

	fmt.Printf("Succesfully recieved vocs: %v", vocs)
	c.JSON(200, vocs)
}

func (h *VocHandler) CreateVoc(c *gin.Context) {
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

func (h *VocHandler) UpdateVoc(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateVocRequest

	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Failed to parse body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedVoc, err := h.vocService.UpdateVoc(id, req)
	if err != nil {
		fmt.Printf("Failed to update voc: %v\n", err)
		c.JSON(500, gin.H{
			"Message": "Error updating a voc",
		})
		return
	}

	fmt.Printf("Succesfully updated voc with id: %v and content: %v", id, req)
	c.JSON(200, updatedVoc)
}

func (h *VocHandler) DeleteVoc(c *gin.Context) {
	id := c.Param("id")

	err := h.vocService.DeleteVoc(id)
	if err != nil {
		fmt.Printf("Failed to delete a vocabulary with id: %v", id)
		c.JSON(500, gin.H{
			"Message": fmt.Sprintf("Failed to delete vocabulary with id: %v", id),
		})
		return
	}

	c.JSON(204, nil)
}
