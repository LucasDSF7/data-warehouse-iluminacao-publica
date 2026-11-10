package viewsdw

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ViewComunicacoesIot(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW dw.view_comunicacoes_iot AS
	SELECT
		data, dia_no_mes, dia_no_ano, nome_dia_semana, semana_do_ano,
		numero_mes, nome_mes, ano, trimestre, trimestre_ano,
		trimestre_contrato_aracaju, id_lcu, codigo_barras, deveui,
		nome_lcu, id_exati_lcu, datahora_cadastro, latitude_lcu,
		longitude_lcu, id_grupo_ip, nome_grupo_ip, e_instalado,
		id_ponto_servico, plaqueta, endereco, latitude, longitude,
		dw.dim_lcu.cidade cidade,
		bairro, logradouro_eixo, logradouro_google, documento_evento,
		e_condominio, e_banco_credito, e_ponto_escuro, valor_creditos,
		e_exclusividade_ip, altura_poste, tipo_altura_poste, tipo_poste,
		e_telegestao, marco_telegestao, largura_via_veiculos,
		largura_via_pedestres, distancia_entre_postes,
		classificacao_via_veiculos, classificacao_via_pedestres,
		idcomando, datahora_comunicacao, tensao,
		luminancia, nivel_dimerizacao, corrente, potencia, fator_potencia,
		energia_consumida, rssi, snr
	FROM dw.fato_comunicacoes_iot s
	
	LEFT JOIN dw.dim_data ON dw.dim_data.data_key = s.data_key
	LEFT JOIN dw.dim_ponto_servico ON dw.dim_ponto_servico.ponto_servico_key = s.ponto_servico_key
	LEFT JOIN dw.dim_lcu ON dw.dim_lcu.lcu_key = s.lcu_key
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
