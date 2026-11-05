package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/LucasDSF7/data-warehouse-iluminacao-publica/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Lembrar de usar defer dbpool.Close()
func GetDriverConecta() *pgxpool.Pool {
	utils.LoadEnv()
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_CONN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}
