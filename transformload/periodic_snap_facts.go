package transformload

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MergePeriodicSnapLcus(dbpool *pgxpool.Pool) {
	query := `MERGE INTO dw.periodicfact_lcu t
	USING(
		SELECT
		COALESCE(d.data_key, -1) data_key,
		COALESCE(l.lcu_key, -1) lcu_key,
		COALESCE(p.ponto_servico_key, -1) ponto_servico_key,
		CASE
			WHEN l.e_instalado = 'Não instalado' THEN 'Removida'
			WHEN s.datahora_comunicacao IS NULL THEN 'Pending'
			WHEN CURRENT_TIMESTAMP::timestamp - s.datahora_comunicacao
			>= INTERVAL '24 hour' THEN 'Off'
			ELSE 'On'
		END AS status
		FROM staging.comunicacoes_iot s
		LEFT JOIN dw.dim_data d
		ON COALESCE(TO_CHAR(CURRENT_DATE, 'YYYYMMDD')::int, -1) = d.data_key
		LEFT JOIN dw.dim_lcu l
		ON l.deveui = s.deveui AND l.indicador_linha_valida = 'Corrente'
		LEFT JOIN dw.dim_ponto_servico p
		ON p.id_ponto_servico = s.id_exati_lcu
	) AS s
	ON s.data_key = t.data_key AND s.lcu_key = t.lcu_key
	WHEN MATCHED AND s.status IS DISTINCT FROM t.status THEN
	UPDATE SET
	lcu_key = COALESCE(s.lcu_key, t.lcu_key),
	status = COALESCE(s.status, t.status)
	WHEN NOT MATCHED THEN
	INSERT (
		data_key, lcu_key, ponto_servico_key, status
	)
	VALUES (
		s.data_key, s.lcu_key, s.ponto_servico_key, s.status
	);

	UPDATE dw.periodicfact_lcu t
	SET status = 'Removida'
	FROM dw.dim_lcu l
	WHERE l.indicador_linha_valida = 'Expirado'
	AND t.data_key = TO_CHAR(CURRENT_DATE, 'YYYYMMDD')::int
	AND l.lcu_key = t.lcu_key
	AND t.status IS DISTINCT FROM 'Removida'
	AND EXISTS (
		SELECT 1
		FROM dw.dim_lcu l2
		WHERE l.deveui = l2.deveui
		AND l2.indicador_linha_valida = 'Corrente'
		AND l2.e_instalado = 'Não instalado'
	);
	`
	// stating.comunicacoes_iot será TRUNCATE ao final do ETL do fato_comunicacoes_iot.
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
