package transformload

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MergeAccFactTicketIot(dbpool *pgxpool.Pool) {
	query := `MERGE INTO dw.accfact_tickets_iot t
	USING (
		SELECT s.id_ticket, s.status, s.tipo_ticket, s.datahora_criacao,
		s.datahora_iniciado, s.datahora_cancelado, s.datahora_conclusao,
		COALESCE(d1.lcu_key, -1) lcu_key,
		COALESCE(d2.data_criacao_key, -1) data_criacao_key,
		COALESCE(d3.data_iniciado_key, -1) data_iniciado_key,
		COALESCE(d4.data_cancelado_key, -1) data_cancelado_key,
		COALESCE(d5.data_conclusao_key, -1) data_conclusao_key,
		COALESCE(p.ponto_servico_key, -1) ponto_servico_key

		FROM staging.tickets_iot s

		LEFT JOIN dw.dim_lcu d1 ON s.deveui = d1.deveui
		AND d1.indicador_linha_valida = 'Corrente'

		LEFT JOIN dw.dim_data_criacao d2
		ON COALESCE(TO_CHAR(s.datahora_criacao, 'YYYYMMDD')::int, -1) = d2.data_criacao_key

		LEFT JOIN dw.dim_data_iniciado d3
		ON COALESCE(TO_CHAR(s.datahora_iniciado, 'YYYYMMDD')::int, -1) = d3.data_iniciado_key

		LEFT JOIN dw.dim_data_cancelado d4
		ON COALESCE(TO_CHAR(s.datahora_cancelado, 'YYYYMMDD')::int, -1) = d4.data_cancelado_key

		LEFT JOIN dw.dim_data_conclusao d5
		ON COALESCE(TO_CHAR(s.datahora_conclusao, 'YYYYMMDD')::int, -1) = d5.data_conclusao_key

		LEFT JOIN dw.dim_ponto_servico p ON p.id_ponto_servico = d1.id_exati_lcu
		
	) AS s
	ON s.id_ticket = t.id_ticket
	WHEN MATCHED AND s.status IS DISTINCT FROM t.status_ticket THEN
	UPDATE SET
	status_ticket = COALESCE(s.status, t.status_ticket),
	data_criacao_key = COALESCE(s.data_criacao_key, t.data_criacao_key),
	data_iniciado_key = COALESCE(s.data_iniciado_key, t.data_iniciado_key),
	data_cancelado_key = COALESCE(s.data_cancelado_key, t.data_cancelado_key),
	data_conclusao_key = COALESCE(s.data_conclusao_key, t.data_conclusao_key),
	lcu_key = COALESCE(s.lcu_key, t.lcu_key),
	datahora_criacao = COALESCE(s.datahora_criacao, t.datahora_criacao),
	datahora_iniciado = COALESCE(s.datahora_iniciado, t.datahora_iniciado),
	datahora_cancelado = COALESCE(s.datahora_cancelado, t.datahora_cancelado),
	datahora_conclusao = COALESCE(s.datahora_conclusao, t.datahora_conclusao),
	tipo_ticket = COALESCE(s.tipo_ticket, t.tipo_ticket),
	ponto_servico_key = COALESCE(s.ponto_servico_key, t.ponto_servico_key)

	WHEN NOT MATCHED THEN
	INSERT (
		id_ticket, data_criacao_key, data_iniciado_key, data_cancelado_key,
		data_conclusao_key, lcu_key, datahora_criacao, datahora_iniciado,
		datahora_cancelado, datahora_conclusao, status_ticket, tipo_ticket,
		ponto_servico_key
	) VALUES (
		s.id_ticket, s.data_criacao_key, s.data_iniciado_key, s.data_cancelado_key,
		s.data_conclusao_key, s.lcu_key, s.datahora_criacao, s.datahora_iniciado,
		s.datahora_cancelado, s.datahora_conclusao, s.status, s.tipo_ticket,
		s.ponto_servico_key
	);

	TRUNCATE TABLE staging.tickets_iot;
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
