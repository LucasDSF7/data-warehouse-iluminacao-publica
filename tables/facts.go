package tables

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FatoComunicacoesIot(dbpool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS dw.fato_comunicacoes_iot(
		data_key BIGINT REFERENCES dw.dim_data(data_key),
		lcu_key BIGINT REFERENCES dw.dim_lcu(lcu_key),
		ponto_servico_key BIGINT REFERENCES dw.dim_ponto_servico(ponto_servico_key),
		datahora_comunicacao TIMESTAMP,
		tensao DECIMAL(5,2),
		luminancia DECIMAL(8,2),
		nivel_dimerizacao SMALLINT,
		corrente DECIMAL(5,2),
		potencia DECIMAL(5,2),
		fator_potencia DECIMAL(3,2),
		energia_consumida DECIMAL(11,2),
		rssi SMALLINT,
		snr DECIMAL (4,1),
		CONSTRAINT uk_fato_lcu_tempo UNIQUE (lcu_key, datahora_comunicacao)
	);
	
	CREATE INDEX IF NOT EXISTS idx_fato_data ON dw.fato_comunicacoes_iot(data_key);
	CREATE INDEX IF NOT EXISTS idx_fato_lcu ON dw.fato_comunicacoes_iot(lcu_key);
	CREATE INDEX IF NOT EXISTS idx_fato_ponto_servico ON dw.fato_comunicacoes_iot(ponto_servico_key);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
