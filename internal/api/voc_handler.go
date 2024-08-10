package api

import (
	"fmt"
	"net/http"

	"github.com/Benjosh95/auto-voc/internal/models"
	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO: Add Context??
type VocHandler struct {
	vocService *services.VocService
	validate   *validator.Validate
}

func NewVocHandler(vocService *services.VocService, validate *validator.Validate) *VocHandler {
	return &VocHandler{
		vocService: vocService,
		validate:   validate,
	}
}

func (h *VocHandler) GetVocs(c *gin.Context) {
	var filter models.VocFilter

	// Bind query parameters to the filter struct
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters"})
		return
	}

	// Validate the filter struct
	if err := h.validate.Struct(filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed"})
		return
	}

	// Call the service layer to get vocs
	vocs, err := h.vocService.GetVocs(filter)
	if err != nil {
		fmt.Printf("Failed to get vocs: %v, err: %v\n", vocs, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error getting vocs",
		})
		return
	}

	fmt.Printf("Successfully received vocs: %v", vocs)
	c.JSON(http.StatusOK, vocs)
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error creating a voc",
		})
		return
	}

	fmt.Printf("Succesfully created voc with id: %v and content: %v = %v", id, req.English, req.German)
	c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error updating a voc",
		})
		return
	}

	fmt.Printf("Succesfully updated voc with id: %v and content: %v", id, req)
	c.JSON(http.StatusOK, updatedVoc)
}

func (h *VocHandler) DeleteVoc(c *gin.Context) {
	id := c.Param("id")

	err := h.vocService.DeleteVoc(id)
	if err != nil {
		// TODO: Error handling is wack
		fmt.Printf("Failed to delete a vocabulary with id: %v\n", id)
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": fmt.Sprintf("Failed to delete vocabulary with id: %v", id),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
