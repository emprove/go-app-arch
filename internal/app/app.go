package app

import (
	"context"
	"sync"

	"go-app-arch/internal/config"
	"go-app-arch/internal/database/postgres"
	"go-app-arch/internal/env"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Cfg *config.Cfg
	DS  *config.DynamicState
	DB  *pgxpool.Pool
	WG  sync.WaitGroup
}

func NewApp() (*Application, error) {
	ds := config.NewDynamicState(env.GetString("APP_LOCALE"))
	locales := []config.Locale{
		{
			Title:    "English",
			Iso:      "en",
			Position: 1,
		},
		{
			Title:    "Русский",
			Iso:      "ru",
			Position: 2,
		},
	}
	allowedOrigins := []string{env.GetString("URL_SHOP"), env.GetString("URL_ADMIN")}
	dbCfg := &config.DBCfg{Dsn: env.GetString("DB_DSN")}
	cfg := config.NewConfig(
		dbCfg,
		env.GetString("APP_URL"),
		env.GetString("APP_LUM_URL"),
		env.GetString("URL_SHOP"),
		env.GetString("URL_ADMIN"),
		env.GetInt("HTTP_PORT"),
		locales,
		allowedOrigins,
	)

	db, err := postgres.New(context.Background(), cfg.GetDBConfig().Dsn)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Cfg: cfg,
		DS:  ds,
		DB:  db,
	}

	return app, nil
}
