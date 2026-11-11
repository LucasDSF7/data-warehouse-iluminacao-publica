package tables

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RolePlayDataCriacao(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW dw.dim_data_criacao AS
	SELECT
		data_key AS data_criacao_key,
		data AS data_criacao,
		dia_no_mes AS dia_no_mes_criacao,
		dia_no_ano AS dia_no_ano_criacao,
		nome_dia_semana AS nome_dia_semana_criacao,
		semana_do_ano AS semana_do_ano_criacao,
		numero_mes AS numero_mes_criacao,
		nome_mes AS nome_mes_criacao,
		ano AS ano_criacao,
		trimestre AS trimestre_criacao,
		trimestre_ano AS trimestre_ano_criacao,
		trimestre_contrato_aracaju AS trimestre_contrato_aracaju_criacao
	FROM dw.dim_data
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func RolePlayDataIniciado(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW dw.dim_data_iniciado AS
	SELECT
		data_key AS data_iniciado_key,
		data AS data_iniciado,
		dia_no_mes AS dia_no_mes_iniciado,
		dia_no_ano AS dia_no_ano_iniciado,
		nome_dia_semana AS nome_dia_semana_iniciado,
		semana_do_ano AS semana_do_ano_iniciado,
		numero_mes AS numero_mes_iniciado,
		nome_mes AS nome_mes_iniciado,
		ano AS ano_iniciado,
		trimestre AS trimestre_iniciado,
		trimestre_ano AS trimestre_ano_iniciado,
		trimestre_contrato_aracaju AS trimestre_contrato_aracaju_iniciado
	FROM dw.dim_data
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func RolePlayDataCancelado(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW dw.dim_data_cancelado AS
	SELECT
		data_key AS data_cancelado_key,
		data AS data_cancelado,
		dia_no_mes AS dia_no_mes_cancelado,
		dia_no_ano AS dia_no_ano_cancelado,
		nome_dia_semana AS nome_dia_semana_cancelado,
		semana_do_ano AS semana_do_ano_cancelado,
		numero_mes AS numero_mes_cancelado,
		nome_mes AS nome_mes_cancelado,
		ano AS ano_cancelado,
		trimestre AS trimestre_cancelado,
		trimestre_ano AS trimestre_ano_cancelado,
		trimestre_contrato_aracaju AS trimestre_contrato_aracaju_cancelado
	FROM dw.dim_data
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

func RolePlayDataConclusao(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW dw.dim_data_conclusao AS
	SELECT
		data_key AS data_conclusao_key,
		data AS data_conclusao,
		dia_no_mes AS dia_no_mes_conclusao,
		dia_no_ano AS dia_no_ano_conclusao,
		nome_dia_semana AS nome_dia_semana_conclusao,
		semana_do_ano AS semana_do_ano_conclusao,
		numero_mes AS numero_mes_conclusao,
		nome_mes AS nome_mes_conclusao,
		ano AS ano_conclusao,
		trimestre AS trimestre_conclusao,
		trimestre_ano AS trimestre_ano_conclusao,
		trimestre_contrato_aracaju AS trimestre_contrato_aracaju_conclusao
	FROM dw.dim_data
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
