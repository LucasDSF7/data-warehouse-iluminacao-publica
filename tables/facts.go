package tables

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Transaction Fact Table.
// Uma linha para cada comunicação de cada dispositivo Iot.
// Possui as métricas da luminária LED em que a LCU está instalada.
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

// Accumulating Snapshot Fact Table.
// Possui os principais milestones referentes a um Ticket de Telegestão.
// Um Ticket possui data de abertura, encaminhamento, conclusão e cancelamento.
// É um factless fact table, então será adicionada uma coluna para contagem.
// id_ticket é uma degenerate dimension.
func FatoTicketsTelegestao(dbpool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS dw.accfact_tickets_iot(
		id_ticket VARCHAR(30),
		data_criacao_key INT,
		data_iniciado_key INT,
		data_cancelado_key INT,
		data_conclusao_key INT,
		lcu_key BIGINT REFERENCES dw.dim_lcu(lcu_key),
		datahora_criacao TIMESTAMP,
		datahora_iniciado TIMESTAMP,
		datahora_cancelado TIMESTAMP,
		datahora_conclusao TIMESTAMP,
		status_ticket VARCHAR(20),
		tipo_ticket VARCHAR(50)
	);

	COMMENT ON TABLE dw.accfact_tickets_iot IS
		'Accumulating Snapshot Fact Table.
		Possui os principais milestones referentes a um Ticket de Telegestão.
		Um Ticket possui data de abertura, encaminhamento, conclusão e cancelamento.
		É um factless fact table, então será adicionada uma coluna para contagem.
		id_ticket é uma degenerate dimension.';
	COMMENT ON COLUMN dw.accfact_tickets_iot.id_ticket IS 'Degenerate dimension com id único de referência do ticket.';
	CREATE INDEX IF NOT EXISTS idx_accfact_id_ticket ON dw.accfact_tickets_iot (id_ticket);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
