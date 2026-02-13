package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tnh-jonghwan/omnichannel-gateway/config"
	"github.com/tnh-jonghwan/omnichannel-gateway/infrastructure"
	respondio "github.com/tnh-jonghwan/omnichannel-gateway/internal/respond-io"
)

func main() {
	// default value
	env := "dev"

	// go run ./cmd dev
	if len(os.Args) > 1 {
		fmt.Printf("os.Args[0]: %s\n", os.Args[0])
		fmt.Printf("os.Args[1]: %s\n\n", os.Args[1])
		env = os.Args[1]
	}

	// load env
	config.LoadEnv(env)

	// nats connection
	nc, err := infrastructure.NewNatsConnection()
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// gin router
	r := gin.Default()
	respondio.RegisterRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
