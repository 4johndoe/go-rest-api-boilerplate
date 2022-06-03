package errors

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/stretchr/testify/assert"
	"go-rest-api/pkg/log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("normal processing", func(t *testing.T) {
		logger, entries := log.NewForTest()
		handler := Handler(logger)
		ctx, res := buildContext(handler, handlerOk)
		assert.Nil(t, ctx.Next())
		assert.Zero(t, entries.Len())
		assert.Equal(t, http.StatusOK, res.Code)
	})

	// todo some more
}

// todo Test_buildErrorResponse

func buildContext(handlers ...routing.Handler) (*routing.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://127.0.0.1/users", nil)
	return routing.NewContext(res, req, handlers...), res
}

func handlerOk(c *routing.Context) error {
	return c.Write("test")
}

// todo handlerError

// todo handlerHTTPError

// todo handlerPanic
