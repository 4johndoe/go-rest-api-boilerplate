package album

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/entity"
	"testing"
)

var errCRUD = errors.New("error crud")

func TestCreateAlbumRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     CreateAlbumRequest
		wantError bool
	}{
		{"success", CreateAlbumRequest{Name: "test"}, false},
		{"required", CreateAlbumRequest{Name: ""}, true},
		{"too long", CreateAlbumRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

// todo TestUpdateAlbumRequest_Validate

// todo Test_service_CRUD

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
