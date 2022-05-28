package test

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"go-rest-api/pkg/log"
)

func MockRouter(logger log.Logger) *routing.Router {
	router := routing.New()
	router.Use(
	// todo accesslog
	// todo errors
	// todo content
	// todo cors
	)
	return router
}
