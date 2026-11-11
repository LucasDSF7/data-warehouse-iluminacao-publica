package schemas

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SchemaDW(dbpool *pgxpool.Pool) {
	query := `CREATE SCHEMA IF NOT EXISTS dw;
	COMMENT ON SCHEMA dw IS 'Schema modelado como Data Warehouse seguindo o livro
	 The Data Warehouse Toolkit Third Edition, com finalidade de
	 criar dados estruturados para decisões gerenciais.';`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func SchemaStaging(dbpool *pgxpool.Pool) {
	query := `CREATE SCHEMA IF NOT EXISTS staging;
	COMMENT ON SCHEMA staging IS 'Schema para armazenamento temporário de dados para etapa de tranformação e carregamento no DW.';
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func SchemaTicketsIot(dbpool *pgxpool.Pool) {
	query := `CREATE SCHEMA IF NOT EXISTS tickets;
	COMMENT ON SCHEMA staging IS 'Schema para depósito de views relacionadas a querys relacionadas a tickets da telegestão.';
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
