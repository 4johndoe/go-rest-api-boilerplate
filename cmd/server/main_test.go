package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-rest-api/pkg/log"
	"testing"
	"time"
)

func Test_logDBQuery(t *testing.T) {
	logger, entries := log.NewForTest()
	f := logDBQuery(logger)
	f(context.Background(), time.Millisecond*3, "sql", nil, nil)
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB query successful", entries.All()[0].Message)
	}
	entries.TakeAll()

	f(context.Background(), time.Millisecond*3, "sql", nil, fmt.Errorf("test"))
	if assert.Equal(t, 1, entries.Len()) {
		assert.Equal(t, "DB query error: test", entries.All()[0].Message)
	}
}
