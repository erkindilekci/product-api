package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var insertProductsStatement = `INSERT INTO products (name, price, discount, store) 
VALUES 
	('XBOX Series X', 1000.0, 10.0, 'Microsoft'),
	('Steelseries Rival 500', 100.0, 20.0, 'Amazon'),
	('Asus Vivobook', 600.0, 15.0, 'Asus Store'),
	('Macbook Pro M3 Pro', 3000.0, 0.0, 'Apple');`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductsResult, err := dbPool.Exec(ctx, insertProductsStatement)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected()))
	}
}
