package api

import (
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
	// request processing
	h.vocService.GetVocs()
	// ...
	// build response
	c.JSON(200, gin.H{
		"English": "endorse",
		"German":  "billigen",
	})
}

func (h *VocHandler) CreateVoc(c *gin.Context) {
	// request processing
	h.vocService.CreateVoc()
	// ...
	// build response
	c.JSON(200, gin.H{
		"Message": "Created Voc successfully",
	})
}
