package respondio

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	service := NewService()
	handler := NewHandler(service)

	group := router.Group("/respondio")

	{
		group.GET("/send-message", handler.SendMessage)
	}
}
