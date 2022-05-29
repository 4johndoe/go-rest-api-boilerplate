package album

import (
	"context"
	"go-rest-api/internal/entity"
	"go-rest-api/pkg/dbcontext"
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
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Album, error) {
	panic("implement me")
}

func (r repository) Count(ctx context.Context) (int, error) {
	panic("implement me")
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Album, error) {
	panic("implement me")
}

func (r repository) Create(ctx context.Context, album entity.Album) error {
	panic("implement me")
}

func (r repository) Update(ctx context.Context, album entity.Album) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
