package viewsdw

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AlertaComandoDesligado(dbpool *pgxpool.Pool) {
	query := `WITH horas_em_comunicar AS (
		SELECT ultima_comunicacao_iot.lcu_key,
		ultima_comunicacao_iot.ponto_servico_key,
		ultima_comunicacao_iot.datahora_comunicacao,
		CURRENT_TIMESTAMP::timestamp without time zone - ultima_comunicacao_iot.datahora_comunicacao AS interval_sem_comunicar
		FROM dw.ultima_comunicacao_iot
	), alarmes AS (
		SELECT ps.idcomando,
		date_trunc('second', min(u.interval_sem_comunicar)) AS interval_sem_comunicar,
		count(u.lcu_key) AS total_lcus,
		count(u.lcu_key) FILTER (WHERE u.interval_sem_comunicar > '04:00:00') AS quatro_horas_sem_comunicar,
		100 * count(u.lcu_key) FILTER (WHERE u.interval_sem_comunicar > '04:00:00') / NULLIF(count(u.lcu_key), 0) AS porcentagem_quatro_horas
		FROM horas_em_comunicar u
		JOIN dw.fato_comunicacoes_iot s ON u.lcu_key = s.lcu_key AND s.datahora_comunicacao = u.datahora_comunicacao
		LEFT JOIN dw.dim_lcu lcu ON lcu.lcu_key = u.lcu_key
		LEFT JOIN dw.dim_ponto_servico ps ON ps.ponto_servico_key = u.ponto_servico_key
		WHERE ps.idcomando <> 'Não preenchido' AND lcu.e_instalado = 'Instalado'
		GROUP BY ps.idcomando
		HAVING count(u.lcu_key) >= 12
		ORDER BY (date_trunc('second', min(u.interval_sem_comunicar))) DESC
	), alarme_descricao AS (
		SELECT
			CASE
				WHEN alarmes.interval_sem_comunicar >= '02:00:00'::interval THEN ((((('O comando de ID ' || alarmes.idcomando) || ' está há mais de ') || (EXTRACT(hour FROM alarmes.interval_sem_comunicar) + EXTRACT(day FROM alarmes.interval_sem_comunicar) * 24)) || ' horas e ') || EXTRACT(minute FROM alarmes.interval_sem_comunicar)) || ' minutos sem comunicar.'
				WHEN alarmes.porcentagem_quatro_horas >= 75 THEN ((('O comando de ID ' || alarmes.idcomando::text) || ' possui ') || alarmes.porcentagem_quatro_horas) || '% das LCUs sem comunicar há mais de 4 horas.'
			ELSE NULL
			END AS alarme
		FROM alarmes
	)
	SELECT alarme
	FROM alarme_descricao
	WHERE alarme IS NOT NULL;`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

// View para consultar Comandos elétricos do parque de Aracaju
// com falta de fase. É gerado um texto contendo o qual o comando
// e quantas luminárias estão com leitura de falta de fase.
// Foi adicionada uma restrição para evitar disparar alerta com
// poucos registros de falta de fase.
func AlertaFaltaFaseComando(dbpool *pgxpool.Pool) {
	query := `CREATE OR REPLACE VIEW tickets.alerta_comando_falta_fase AS
	WITH contagem_falta_fase AS (
		SELECT p.idcomando, count(l.lcu_key) FILTER (WHERE s.tensao <= 80) count_falta
		FROM dw.ultima_comunicacao_iot u
		JOIN dw.fato_comunicacoes_iot s ON s.lcu_key = u.lcu_key
		AND s.datahora_comunicacao = u.datahora_comunicacao
		LEFT JOIN dw.dim_lcu l ON l.lcu_key = u.lcu_key
		LEFT JOIN dw.dim_ponto_servico p ON p.ponto_servico_key = u.ponto_servico_key
		WHERE p.idcomando <> 'Não preenchido'
		AND l.e_instalado = 'Instalado'
		AND u.datahora_comunicacao::date >= CURRENT_DATE - INTERVAL '3 day'
		GROUP BY p.idcomando
		HAVING count(l.lcu_key) FILTER (WHERE s.tensao <= 80) >= 5
		ORDER BY count_falta DESC
	)
	SELECT 'Falta de fase no comando de ID ' || idcomando || ', com '
	|| count_falta || 'LCUs com tensão inferior a 80 V.'
	FROM contagem_falta_fase;`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}

// Um ticket é um tipo de falha/ocorrência/alarme do
// sistema de telegestão. Falha no Driver indica que
// a luminária LED possui defeito no driver, sendo necessário
// intervenção em campo.
func TicketFalhaDriver(dbpool *pgxpool.Pool) {}
