package infrastructure

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tnh-jonghwan/omnichannel-gateway/config"
)

func NewNatsConnection() (*nats.Conn, error) {
	nc, err := nats.Connect(config.AppConfig.NatsUrl, nats.UserInfo(config.AppConfig.NatsUser, config.AppConfig.NatsPassword))
	if err != nil {
		return nil, fmt.Errorf("nats connect error: %w", err)
	}

	return nc, nil
}
