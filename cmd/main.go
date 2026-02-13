package main

import (
	"github.com/gin-gonic/gin"
	respondio "github.com/tnh-jonghwan/omnichannel-gateway/internal/respond-io"
)

func main() {
	r := gin.Default()

	respondio.RegisterRoutes(r)

	r.Run(":8080")
}
