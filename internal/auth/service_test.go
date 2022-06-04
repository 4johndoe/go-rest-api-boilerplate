package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/entity"
	"go-rest-api/internal/errors"
	"go-rest-api/pkg/log"
	"testing"
)

func Test_service_Authenticate(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService("test", 100, logger)
	_, err := s.Login(context.Background(), "unknown", "bad")
	assert.Equal(t, errors.Unathorized(""), err)
	token, err := s.Login(context.Background(), "demo", "pass")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func Test_service_authenticate(t *testing.T) {
	logger, _ := log.NewForTest()
	s := service{"test", 100, logger}
	assert.Nil(t, s.authenticate(context.Background(), "unknown", "bad"))
	assert.NotNil(t, s.authenticate(context.Background(), "demo", "pass"))
}

func Test_service_GenerateJWT(t *testing.T) {
	logger, _ := log.NewForTest()
	s := service{"test", 100, logger}
	token, err := s.generateJWT(entity.User{
		ID:   "100",
		Name: "demo",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, token)
	}
}
