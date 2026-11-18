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

// Granularidade dos timestamps do sistema operacional é minutos.
func StagingOcorrencias(dbpool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS staging.ocorrencias (
		id_ponto_servico INT,
		data_hora_reclamacao TIMESTAMP,
		data_hora_limite_atendimento TIMESTAMP,
		data_hora_inicio_atendimento TIMESTAMP,
		data_hora_conclusao_atendimento TIMESTAMP,
		data_hora_autorizacao_ocorrencia TIMESTAMP,
		nome_reclamante VARCHAR(50),
		hash_reclamante TEXT,
		status_atendimento VARCHAR(20),
		motivo_atendimento VARCHAR(20),
		solucao_atendimento VARCHAR(20),
		nome_tipo_ocorrencia VARCHAR(50),
		nome_equipe VARCHAR(20),
		id_status_ocorrencia SMALLINT,
		nome_status_ocorrencia VARCHAR(20),
		nome_origem_ocorrencia VARCHAR(20),
		id_ordem_servico INT,
		index_ordem_servico SMALLINT,
		id_atendimento INT,
		latitude_ocorrencia FLOAT,
		longitude_ocorrencia FLOAT,
		id_ocorrencia INT,
		index_ocorrencia SMALLINT,
		protocolo_solicitacao text,
		contagem_solicitacao SMALLINT,
		contagem_atendimento SMALLINT,
		sigla_prioridade_ocorrencia VARCHAR(10),
		descricao_prioridade_ocorrencia VARCHAR(20),
		id_reprogramacao_atual INT
	);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func StagingTicketsIot(dbpool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS staging.tickets_iot (
		id_ticket VARCHAR(30),
		deveui VARCHAR(20),
		status VARCHAR(20),
		tipo_ticket VARCHAR(50),
		datahora_criacao TIMESTAMP,
		datahora_iniciado TIMESTAMP,
		datahora_cancelado TIMESTAMP,
		datahora_conclusao TIMESTAMP
	);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
