package healthcheck

import (
	"go-rest-api/internal/test"
	"go-rest-api/pkg/log"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	RegisterHandler(router, "0.9.0")
	test.Endpoint(t, router, test.APITestCase{
		"ok", "GET", "/healthcheck", "", nil, http.StatusOK, `"OK 0.9.0"`,
	})
}
