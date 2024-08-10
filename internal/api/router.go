package api

import (
	"time"

	"github.com/Benjosh95/auto-voc/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewRouter(vocService *services.VocService, validate *validator.Validate) *gin.Engine {
	router := gin.Default()

	// set up middlewares (CORS, ...)
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init handler(s) with service(s), and more
	vocHandler := NewVocHandler(vocService, validate)
	appHandler := NewAppHandler()

	// Define routes and map them to handlers
	router.GET("/vocs", vocHandler.GetVocs) // filtering possible
	router.POST("/vocs", vocHandler.CreateVoc)
	router.PUT("/vocs/:id", vocHandler.UpdateVoc)
	router.DELETE("vocs/:id", vocHandler.DeleteVoc)

	router.GET("/health", appHandler.Health)

	return router
}
