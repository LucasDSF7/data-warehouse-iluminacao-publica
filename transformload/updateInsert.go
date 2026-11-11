package transformload

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateInsertDimLcu(dbpool *pgxpool.Pool, cidade string) {
	query := `UPDATE dw.dim_lcu t
		SET indicador_linha_valida = 'Expirado',
		data_expiracao_linha = GREATEST(data_validacao_linha, (CURRENT_DATE - INTERVAL '1 day')::date)
		FROM staging.comunicacoes_iot s
		WHERE t.deveui = s.deveui
		AND t.indicador_linha_valida = 'Corrente'
		AND (
			t.codigo_barras IS DISTINCT FROM s.codigo_barras OR
			t.nome_lcu IS DISTINCT FROM s.nome_lcu OR
			t.id_exati_lcu IS DISTINCT FROM s.id_exati_lcu OR
			t.latitude_lcu IS DISTINCT FROM s.latitude_lcu OR
			t.longitude_lcu IS DISTINCT FROM s.longitude_lcu OR
			t.id_grupo_ip IS DISTINCT FROM s.id_grupo_ip OR
			t.nome_grupo_ip IS DISTINCT FROM s.nome_grupo_ip
		);

		INSERT INTO dw.dim_lcu (
			id_lcu, codigo_barras, deveui, nome_lcu,
			id_exati_lcu, datahora_cadastro, cidade, latitude_lcu,
			longitude_lcu, id_grupo_ip, nome_grupo_ip, e_instalado,
			data_validacao_linha
		)
		SELECT
			s.id_lcu,
			s.codigo_barras,
			s.deveui,
			s.nome_lcu,
			s.id_exati_lcu,
			s.datahora_cadastro,
			'$1',
			s.latitude_lcu,
			s.longitude_lcu,
			s.id_grupo_ip,
			s.nome_grupo_ip,
			'Instalado',
			CURRENT_DATE
		FROM staging.comunicacoes_iot s
		WHERE NOT EXISTS (
			SELECT 1
			FROM dw.dim_lcu t
			WHERE t.deveui = s.deveui
			AND t.indicador_linha_valida = 'Corrente'
		);

		WITH removed_iot AS (
			UPDATE dw.dim_lcu t
			SET indicador_linha_valida = 'Expirado',
			data_expiracao_linha = GREATEST(data_validacao_linha, (CURRENT_DATE - INTERVAL '1 day')::date)
			WHERE NOT EXISTS (
				SELECT 1
				FROM staging.comunicacoes_iot s
				WHERE t.deveui = s.deveui
			)
			AND t.indicador_linha_valida = 'Corrente'
			AND t.lcu_key IS DISTINCT FROM -1
			AND t.e_instalado = 'Instalado'
			RETURNING id_lcu, codigo_barras, deveui, nome_lcu,
				id_exati_lcu, datahora_cadastro, cidade, latitude_lcu,
				longitude_lcu, id_grupo_ip, nome_grupo_ip, e_instalado,
				data_validacao_linha
		)
		INSERT INTO dw.dim_lcu (
			id_lcu, codigo_barras, deveui, nome_lcu,
			id_exati_lcu, datahora_cadastro, cidade, latitude_lcu,
			longitude_lcu, id_grupo_ip, nome_grupo_ip, e_instalado,
			data_validacao_linha
		)
		SELECT
			s.id_lcu,
			s.codigo_barras,
			s.deveui,
			s.nome_lcu,
			s.id_exati_lcu,
			s.datahora_cadastro,
			'$1',
			s.latitude_lcu,
			s.longitude_lcu,
			s.id_grupo_ip,
			s.nome_grupo_ip,
			'Não instalado',
			CURRENT_DATE
		FROM removed_iot s;
	`
	_, err := dbpool.Exec(context.Background(), query, cidade)
	if err != nil {
		fmt.Println(err)
	}
}
