package tables

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StagingComunicacoesIot(dbpool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS staging.comunicacoes_iot (
		id_lcu VARCHAR(50),
		codigo_barras VARCHAR(30),
		deveui VARCHAR(50),
		nome_lcu VARCHAR(50),
		id_exati_lcu INT,
		datahora_cadastro timestamp,
		latitude_lcu FLOAT,
		longitude_lcu FLOAT,
		id_grupo_ip VARCHAR(40),
		nome_grupo_ip VARCHAR(40),
		datahora_comunicacao timestamp,
		datahora_linha timestamp,
		tensao DECIMAL(5,2),
		luminancia DECIMAL(8,2),
		nivel_dimerizacao SMALLINT,
		corrente DECIMAL(5,2),
		potencia DECIMAL(5,2),
		fator_potencia DECIMAL(3,2),
		energia_consumida DECIMAL(11,2),
		rssi SMALLINT,
		snr DECIMAL (4,1)
	)
	
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
