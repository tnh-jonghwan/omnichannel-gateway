package respondio

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) SendMessage(c *gin.Context) {
	result := handler.service.SendMessage()

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
