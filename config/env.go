package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	NodeEnv             string
	NatsUrl             string
	NatsUser            string
	NatsPassword        string
	MariadbHost         string
	MariadbRootPassword string
	MariadbDatabase     string
	MariadbUser         string
	MariadbPassword     string
	MariadbPort         string
}

var AppConfig Config

func LoadEnv(env string) {
	envFile := ".env." + env

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("env 파일 로드 실패")
	}

	AppConfig = Config{
		Port:                os.Getenv(EnvPort),
		NodeEnv:             os.Getenv(EnvNodeEnv),
		NatsUrl:             os.Getenv(EnvNatsUrl),
		NatsUser:            os.Getenv(EnvNatsUser),
		NatsPassword:        os.Getenv(EnvNatsPassword),
		MariadbHost:         os.Getenv(EnvMariadbHost),
		MariadbRootPassword: os.Getenv(EnvMariadbRootPassword),
		MariadbDatabase:     os.Getenv(EnvMariadbDatabase),
		MariadbUser:         os.Getenv(EnvMariadbUser),
		MariadbPassword:     os.Getenv(EnvMariadbPassword),
		MariadbPort:         os.Getenv(EnvMariadbPort),
	}
}
