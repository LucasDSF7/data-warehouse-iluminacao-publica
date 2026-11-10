package transformload

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertFatoComunicacoesIot(dbpool *pgxpool.Pool) {
	query := `INSERT INTO dw.fato_comunicacoes_iot (
		data_key, lcu_key, ponto_servico_key, datahora_comunicacao,
		tensao, luminancia, nivel_dimerizacao, corrente, potencia,
		fator_potencia, energia_consumida, rssi, snr
	)
	SELECT
		COALESCE(dimdata.data_key, -1),
		COALESCE(dimlcu.lcu_key, -1),
		COALESCE(dimponto_servico.ponto_servico_key, -1),
		s.datahora_comunicacao,
		s.tensao,
		s.luminancia,
		s.nivel_dimerizacao,
		s.corrente,
		s.potencia,
		s.fator_potencia,
		s.energia_consumida,
		s.rssi,
		s.snr
	FROM staging.comunicacoes_iot s

	LEFT JOIN dw.dim_data dimdata
	ON TO_CHAR(s.datahora_comunicacao::date, 'YYYYMMDD')::integer = dimdata.data_key

	LEFT JOIN dw.dim_lcu dimlcu
	ON s.deveui = dimlcu.deveui AND dimlcu.indicador_linha_valida = 'Corrente'

	LEFT JOIN dw.dim_ponto_servico dimponto_servico
	ON s.id_exati_lcu = dimponto_servico.id_ponto_servico

	ON CONFLICT(lcu_key, datahora_comunicacao) DO NOTHING;
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
