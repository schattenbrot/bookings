package dbrepo

import (
	"database/sql"

	"github.com/Schattenbrot/bookings/internal/config"
	"github.com/Schattenbrot/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: app,
		DB:  conn,
	}
}

func NewTestingRepo(app *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: app,
	}
}
