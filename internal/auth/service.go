package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go-rest-api/internal/entity"
	"go-rest-api/internal/errors"
	"go-rest-api/pkg/log"
	"time"
)

type Service interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type Identity interface {
	GetID() string
	GetName() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

func NewService(signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{signingKey, tokenExpiration, logger}
}

func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unathorized("")
}

func (s service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)

	if username == "demo" && password == "pass" {
		logger.Infof("authentication successful")
		return entity.User{ID: "100", Name: "demo"}
	}

	logger.Infof("authentication failed")
	return nil
}

func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
