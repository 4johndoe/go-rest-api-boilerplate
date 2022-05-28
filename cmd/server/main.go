package main

import (
	"flag"
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"go-rest-api/internal/config"
	"go-rest-api/pkg/log"
	"net/http"
	"os"
)

var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// todo db connect

	// build http server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, cfg),
	}

	// start http server with graceful shutdown
	fmt.Println("server %v is running at %v", address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(-1)
	}
}

func buildHandler(logger log.Logger, cfg *config.Config) http.Handler {
	router := routing.New()

	return router
}
