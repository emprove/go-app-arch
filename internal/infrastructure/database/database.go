package database

import "context"

type DB interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	Close() error
}

type Rows interface {
	Close()
}
