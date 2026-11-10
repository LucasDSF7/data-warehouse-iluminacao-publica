package viewsdw

import "github.com/jackc/pgx/v5/pgxpool"

// Um ticket é um tipo de falha/ocorrência/alarme do
// sistema de telegestão. Falha no Driver indica que
// a luminária LED possui defeito no driver, sendo necessário
// intervenção em campo.
func TicketFalhaDriver(dbpool *pgxpool.Pool) {}
