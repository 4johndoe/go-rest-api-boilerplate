package test

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	"go-rest-api/internal/config"
	"go-rest-api/pkg/dbcontext"
	"go-rest-api/pkg/log"
	"path"
	"runtime"
	"testing"
)

var db *dbcontext.DB

func DB(t *testing.T) *dbcontext.DB {
	if db != nil {
		return db
	}
	logger, _ := log.NewForTest()
	dir := getSourcePath()
	cfg, err := config.Load(dir+"/../../config/local.yml", logger)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	dbc, err := dbx.MustOpen("postgres", cfg.DSN)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	dbc.LogFunc = logger.Infof
	db = dbcontext.New(dbc)
	return db
}

func ResetTables(t *testing.T, db *dbcontext.DB, tables ...string) {
	for _, table := range tables {
		_, err := db.DB().TruncateTable(table).Execute()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
