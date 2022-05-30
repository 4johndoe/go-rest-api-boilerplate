package album

import (
	"go-rest-api/internal/entity"
	"go-rest-api/internal/test"
	"go-rest-api/pkg/log"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Album{
		{"123", "album123", time.Now(), time.Now()},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), logger)

}
