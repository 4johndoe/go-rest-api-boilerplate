package album

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/internal/entity"
)

var errCRUD = errors.New("error crud")

type mockRepository struct {
	items []entity.Album
}

func (m mockRepository) Get(ctx context.Context, id string) (entity.Album, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Album{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(ctx context.Context, offset, limit int) ([]entity.Album, error) {
	return m.items, nil
}

func (m mockRepository) Create(ctx context.Context, album entity.Album) error {
	if album.Name == "error" {
		return errCRUD
	}
	m.items = append(m.items, album)
	return nil
}

func (m mockRepository) Update(ctx context.Context, album entity.Album) error {
	if album.Name == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.ID == album.ID {
			m.items[i] = album
			break
		}
	}
	return nil
}

func (m mockRepository) Delete(ctx context.Context, id string) error {
	for i, item := range m.items {
		if item.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}
