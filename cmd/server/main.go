package main

import (
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"net/http"
	"os"
)

func main() {
	// todo config

	// todo db connect

	// build http server
	address := fmt.Sprintf(":%v", 8080)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(),
	}

	// start http server with graceful shutdown
	fmt.Println("server %v is running at %v", address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		os.Exit(-1)
	}
}

func buildHandler() http.Handler {
	router := routing.New()

	return router
}
