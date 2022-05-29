package album

import (
	"context"
	"go-rest-api/internal/entity"
	"go-rest-api/pkg/dbcontext"
	"go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	Get(ctx context.Context, id string) (entity.Album, error)
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]entity.Album, error)
	Create(ctx context.Context, album entity.Album) error
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Album, error) {
	var album entity.Album
	err := r.db.With(ctx).Select().Model(id, &album)
	return album, err
}

// Count returns the number of the album records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("album").Row(&count)
	return count, err
}

// Create saves a new album record in the database.
// It returns the ID of the newly inserted album record.
func (r repository) Create(ctx context.Context, album entity.Album) error {
	return r.db.With(ctx).Model(&album).Insert()
}

// Update saves the changes to an album in the database.
func (r repository) Update(ctx context.Context, album entity.Album) error {
	return r.db.With(ctx).Model(&album).Update()
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	album, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&album).Delete()
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Album, error) {
	var albums []entity.Album
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&albums)
	return albums, err
}
