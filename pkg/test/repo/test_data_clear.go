package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, err := dbPool.Exec(ctx, "TRUNCATE products RESTART IDENTITY")
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Products table truncated")
	}
}
