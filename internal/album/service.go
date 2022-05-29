package album

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-rest-api/internal/entity"
	"go-rest-api/pkg/log"
)

type Service interface {
	Get(ctx context.Context, id string) (Album, error)
	Query(ctx context.Context, offset, limit int) ([]Album, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateAlbumRequest) (Album, error)
	Update(ctx context.Context, id string, input UpdateAlbumRequest) (Album, error)
	Delete(ctx context.Context, id string) (Album, error)
}

type Album struct {
	entity.Album
}

type CreateAlbumRequest struct {
	Name string `json:"name"`
}

func (m CreateAlbumRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type UpdateAlbumRequest struct {
	Name string `json:"name"`
}

func (m UpdateAlbumRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func (s service) Get(ctx context.Context, id string) (Album, error) {
	panic("implement me")
}

func (s service) Query(ctx context.Context, offset, limit int) ([]Album, error) {
	panic("implement me")
}

func (s service) Count(ctx context.Context) (int, error) {
	panic("implement me")
}

func (s service) Create(ctx context.Context, input CreateAlbumRequest) (Album, error) {
	panic("implement me")
}

func (s service) Update(ctx context.Context, id string, input UpdateAlbumRequest) (Album, error) {
	panic("implement me")
}

func (s service) Delete(ctx context.Context, id string) (Album, error) {
	panic("implement me")
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}
