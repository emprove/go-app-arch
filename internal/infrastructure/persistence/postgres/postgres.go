package postgres

import (
	"context"

	"go-app-arch/internal/infrastructure/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Query(ctx context.Context, sql string, args ...any) (database.Rows, error) {
	return p.Pool.Query(ctx, sql, args...)
}

func (p *Postgres) Close() error {
	p.Pool.Close()
	return nil
}
