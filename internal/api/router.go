package api

import (
	"time"

	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(vocService *services.VocService) *gin.Engine {
	router := gin.Default()

	// set up middlewares (CORS, ...)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init handler(s) with service(s)
	vocHandler := NewVocHandler(vocService)

	// Define routes and map them to handlers
	router.GET("/vocs", vocHandler.GetVocs)
	router.POST("/vocs", vocHandler.CreateVoc)
	router.DELETE("vocs/:id", vocHandler.DeleteVoc)

	return router
}
