package main

import (
	"github.com/meysamhadeli/shop-golang-microservices/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/config"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/internal/products/server"
	"log"
	"os"
)

const dev = "development"
const production = "production"

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = dev
	}

	cfg, err := config.InitConfig(env)
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	appLogger.Fatal(server.NewServer(appLogger, cfg).Run())
}
