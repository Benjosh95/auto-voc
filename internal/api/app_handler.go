package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) Health(c *gin.Context) {
	c.Status(http.StatusOK)
}
