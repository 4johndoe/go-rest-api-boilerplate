package album

import (
	"context"
	"go-rest-api/internal/entity"
	"go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.Album, error)
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]entity.Album, error)
	Create(ctx context.Context, album entity.Album) error
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	//db *dbcontext.DB todo
	logger log.Logger
}
