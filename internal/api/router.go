package api

import (
	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(vocService *services.VocService) *gin.Engine {
	router := gin.Default()

	// Init handler(s) with service(s)
	vocHandler := NewVocHandler(vocService)

	// Define routes and map them to handlers
	router.GET("/vocs", vocHandler.GetVocs)
	router.POST("/vocs", vocHandler.CreateVoc)

	return router
}
