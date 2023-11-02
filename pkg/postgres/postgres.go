package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Builder *squirrel.StatementBuilderType
}

func Build(ctx context.Context, url string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	var pool *pgxpool.Pool

	for i := 0; i < 5; i++ {
		pool, err = pgxpool.NewWithConfig(ctx, config)
		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &Postgres{pool, &builder}, nil
}
