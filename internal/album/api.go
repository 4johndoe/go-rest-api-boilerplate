package album

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"go-rest-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers. todo add authHandler routing.Handler,
func RegisterHandler(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/albums/<id>", res.get)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	album, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(album)
}
